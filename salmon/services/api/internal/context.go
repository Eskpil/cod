package internal

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/digitalocean/go-libvirt"
	"github.com/eskpil/salmon/pkg/models"
	_ "github.com/eskpil/salmon/services/api/database"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
	"gopkg.in/yaml.v2"

	log "github.com/sirupsen/logrus"
)

type Context struct {
	Hosts    map[string]*libvirt.Libvirt
	Machines map[string]models.Machine

	config ConfigFile
}

func connectWithHost(config ConfigFile, host HostConfig) (*ssh.Client, error) {
	sshKey, err := ioutil.ReadFile(config.Ssh.KeyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read ssh key: %w", err)
	}

	signer, err := ssh.ParsePrivateKeyWithPassphrase(sshKey, []byte("test"))
	if err != nil {
		return nil, err
	}

	username := host.Username
	if username == "" {
		username = "salmon"
	}

	hostKeyCallback, err := knownhosts.New(os.ExpandEnv("$HOME/.ssh/known_hosts"))
	if err != nil {
		return nil, fmt.Errorf("failed to read ssh known hosts: %w", err)
	}

	cfg := ssh.ClientConfig{
		User:            username,
		HostKeyCallback: hostKeyCallback,
		Auth:            []ssh.AuthMethod{ssh.PublicKeys(signer)},
		Timeout:         2 * time.Second,
	}

	sshClient, err := ssh.Dial("tcp", fmt.Sprintf("%s:22", host.Hostname), &cfg)

	if err != nil {
		return nil, err
	}

	return sshClient, err
}

func NewContext() (*Context, error) {
	if len(os.Args) < 2 {
		return nil, fmt.Errorf("Missing config file argument")
	}

	data, err := os.ReadFile(os.Args[1])

	if err != nil {
		return nil, err
	}

	var config ConfigFile
	err = yaml.Unmarshal(data, &config)

	if err != nil {
		return nil, err
	}

	context := &Context{
		Hosts:    make(map[string]*libvirt.Libvirt),
		Machines: make(map[string]models.Machine),
		config:   config,
	}

	for _, host := range config.Hosts {
		sshClient, err := connectWithHost(config, host)
		if err != nil {
			return nil, err
		}

		c, err := sshClient.Dial("unix", host.LibvirtAddress)
		if err != nil {
			return nil, err
		}

		log.Infof("Connected with: %s over ssh transport", host.Hostname)

		l := libvirt.New(c)

		if err := l.Connect(); err != nil {
			return nil, fmt.Errorf("failed to connect: %v", err)
		}

		context.Hosts[host.Name] = l
	}

	return context, nil
}
