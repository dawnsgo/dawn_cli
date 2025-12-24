package registry

import (
	_ "embed"
)

const (
	Consul = "consul"
	Etcd   = "etcd"
	Nacos  = "nacos"
)

const (
	EtcdPackage   = "github.com/dawnsgo/dawn/registry/etcd"
	NacosPackage  = "github.com/dawnsgo/dawn/registry/nacos"
	ConsulPackage = "github.com/dawnsgo/dawn/registry/consul"
)

var (
	//go:embed consul.toml
	ConsulTemplate string
	//go:embed etcd.toml
	EtcdTemplate string
	//go:embed nacos.toml
	NacosTemplate string
)
