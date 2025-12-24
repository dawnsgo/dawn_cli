package config

import (
	_ "embed"
)

const (
	File   = "file"
	Consul = "consul"
	Etcd   = "etcd"
	Nacos  = "nacos"
)

const (
	EtcdPackage   = "github.com/dawnsgo/dawn/config/etcd"
	NacosPackage  = "github.com/dawnsgo/dawn/config/nacos"
	ConsulPackage = "github.com/dawnsgo/dawn/config/consul"
)

var (
	//go:embed file.toml
	FileTemplate string
	//go:embed consul.toml
	ConsulTemplate string
	//go:embed etcd.toml
	EtcdTemplate string
	//go:embed nacos.toml
	NacosTemplate string
)
