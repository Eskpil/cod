package internal

import (
	"context"
	"fmt"
	"strings"

	"github.com/antchfx/xmlquery"
	"github.com/digitalocean/go-libvirt"
	"github.com/eskpil/salmon/pkg/models"
	"github.com/google/uuid"

	log "github.com/sirupsen/logrus"

	machineService "github.com/eskpil/salmon/services/api/internal/machines"
	storagePoolService "github.com/eskpil/salmon/services/api/internal/storagePools"
)

func (m *Context) PerformRoutine(ctx context.Context) error {
	for host, l := range m.Hosts {
		// Collect domains and their respective interfaces.

		hostConfig := HostConfig{}

		for _, c := range m.config.Hosts {
			if c.Name == host {
				hostConfig = c
			}
		}

		// Collect all the storage pools and respective volumes
		{
			pools, _, err := l.ConnectListAllStoragePools(1, libvirt.ConnectListStoragePoolsInactive|libvirt.ConnectListStoragePoolsActive)

			if err != nil {
				log.Errorf("Could not get storage pools of host: %s: %v\n", host, err)
				continue
			}

			for _, pool := range pools {
				active, err := l.StoragePoolIsActive(pool)
				if err != nil {
					log.Errorf("Could not get state of storage pool: %s, %v\n", pool.Name, err)
					continue
				}

				// If the storage pool is inactive it cannot be used by any domains
				// and therefor is irrelevant for us.
				if 1 > active {
					continue
				}

				poolId, err := uuid.FromBytes(pool.UUID[:])
				if err != nil {
					log.Errorf("Could not construct a poolId for: %s: %v\n", pool.Name, err)
					continue
				}

				xml, err := l.StoragePoolGetXMLDesc(pool, 0)
				if err != nil {
					log.Errorf("Could not get xml description of pool: %s: %v\n", pool.Name, err)
					continue
				}

				doc, err := xmlquery.Parse(strings.NewReader(xml))
				if err != nil {
					log.Errorf("Could not construct a new xmlquery from storage pool xml desc: %s: %v\n", pool.Name, err)
					continue
				}

				targetPath := xmlquery.FindOne(doc, "//pool/target/path")

				p := models.StoragePool{
					Id:         poolId.String(),
					Name:       pool.Name,
					TargetPath: targetPath.InnerText(),
					Host:       host,
				}

				if _, err := storagePoolService.GetById(ctx, p.Id); err != nil {
					storagePoolService.Create(ctx, p)
				}
			}
		}

		{
			domains, err := l.Domains()
			if err != nil {
				log.Errorf("Could not get domains: %v", err)
				continue
			}

			for _, domain := range domains {
				active, err := l.DomainIsActive(domain)
				if err != nil {
					log.Errorf("Could not get state of domain: %v\n", err)
					continue
				}

				// If the domain is active (running) active will be set to one. And since
				// right now we only care about active machines we skip everything else.
				if 1 > active {
					continue
				}

				domainId, err := uuid.FromBytes(domain.UUID[:])

				if err != nil {
					log.Errorf("Could not construct a domainId: %v\n", err)
					continue
				}

				// 2 means we use the guest agent living on the domain.
				hostname, err := l.DomainGetHostname(domain, 2)

				if err != nil {
					log.Errorf("Could not get the hostname of the domain: %v\n", err)
					hostname = "<unknown>"
				}

				fqdn := ""

				if hostConfig.DNS.Zone != "" {
					fqdn = fmt.Sprintf("%s.%s", hostname, hostConfig.DNS.Zone)
				} else {
					fqdn = "<unknown>"
				}

				machine := models.Machine{
					Id:       domainId.String(),
					Name:     domain.Name,
					Host:     host,
					Hostname: hostname,
					Fqdn:     fqdn,
				}

				interfaces, err := l.DomainInterfaceAddresses(domain, 1, 0)
				if err != nil {
					log.Errorf("Could not get interfaces for domain: %d\n", domain.Name)
				}

				for _, i := range interfaces {
					ipAddrs := []models.IpAddr{}

					for _, a := range i.Addrs {
						ipAddrs = append(ipAddrs, models.IpAddr{
							Type:   a.Type,
							Addr:   a.Addr,
							Prefix: a.Prefix,
						})
					}

					c := models.Interface{
						Name:    i.Name,
						Mac:     i.Hwaddr[0],
						IpAddrs: ipAddrs,
					}

					machine.Interfaces = append(machine.Interfaces, c)
				}

				if _, err = machineService.GetById(ctx, machine.Id); err != nil {
					if err = machineService.Create(ctx, machine); err != nil {
						log.Errorf("Failed to create machine: %s in the database: %v\n", machine.Fqdn, err)
					}
				}

				// m.Machines[machine.Id] = machine
			}

		}
	}

	return nil
}
