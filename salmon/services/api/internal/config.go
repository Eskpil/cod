package internal

type ConfigFile struct {
	Ssh struct {
		KeyPath string `yaml:"keyPath,omitempty"`
	} `yaml:"ssh"`
	DNS struct {
		Zone string `json:"zone"`
	} `yaml:"dns"`
	Hosts []HostConfig
}

type HostConfig struct {
	Name string `yaml:"name"`
	DNS  struct {
		Zone string `json:"zone"`
	} `yaml:"dns"`
	Hostname       string `yaml:"hostname"`
	Username       string `yaml:"username,omitempty"`
	LibvirtAddress string `yaml:"libvirtAddress,omitempty"`
}
