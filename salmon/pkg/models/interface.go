package models

type IpAddr struct {
	Type   int32  `json:"type" yaml:"type"`
	Addr   string `json:"addr" yaml:"addr"`
	Prefix uint32 `json:"prefix" yaml:"prefix"`
}

type Interface struct {
	Name    string   `json:"name" yaml:"name"`
	Mac     string   `json:"mac" yaml:"mac"`
	IpAddrs []IpAddr `json:"addrs" yaml:"addrs"`
}
