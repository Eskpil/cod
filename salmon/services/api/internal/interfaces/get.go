package interfaces

import (
	"context"

	"github.com/eskpil/salmon/pkg/models"
	"github.com/eskpil/salmon/services/api/database"
)

func getAllAddrs(ctx context.Context, ifaceid string) ([]models.IpAddr, error) {
	addrs := []models.IpAddr{}

	rows, err := database.Conn.ExecRows(ctx, `SELECT type, addr, prefix FROM addrs WHERE ifaceid = ?`, ifaceid)
	if err != nil {
		return addrs, err
	}

	for rows.Next() {
		addr := models.IpAddr{}
		if err := rows.Scan(&addr.Type, &addr.Addr, &addr.Prefix); err != nil {
			return addrs, nil
		}

		addrs = append(addrs, addr)
	}

	return addrs, nil
}

func GetAll(ctx context.Context, machineId string) ([]models.Interface, error) {
	interfaces := []models.Interface{}

	rows, err := database.Conn.ExecRows(ctx, `SELECT mac, id, name FROM interfaces WHERE machineid = ?`, machineId)
	if err != nil {
		return interfaces, err
	}

	for rows.Next() {
		iface := models.Interface{}
		id := ""
		if err := rows.Scan(&iface.Mac, &id, &iface.Name); err != nil {
			return interfaces, err
		}

		if addrs, err := getAllAddrs(ctx, id); err != nil {
			return interfaces, err
		} else {
			iface.IpAddrs = addrs
		}

		interfaces = append(interfaces, iface)
	}

	if err := rows.Err(); err != nil {
		return interfaces, err
	}

	return interfaces, nil
}
