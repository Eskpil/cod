package models

type Machine struct {
	Id         string      `json:"id" yaml:"id"`
	Fqdn       string      `json:"fqdn" yaml:"fqdn"`
	Name       string      `json:"name" yaml:"name"`
	Groups     []string    `json:"groups" yaml:"groups"`
	Host       string      `json:"host" yaml:"host"`
	Hostname   string      `json:"hostname" yaml:"hostname"`
	Interfaces []Interface `json:"interfaces" yaml:"interfaces"`
}
