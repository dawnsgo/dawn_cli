package network

import (
	_ "embed"
)

const (
	WS  = "ws"
	TCP = "tcp"
)

const (
	WSPackage  = "github.com/dawnsgo/dawn/network/ws"
	TCPPackage = "github.com/dawnsgo/dawn/network/tcp"
)

var (
	//go:embed ws_server.toml
	WSServerTemplate string
	//go:embed ws_client.toml
	WSClientTemplate string
	//go:embed tcp_server.toml
	TCPServerTemplate string
	//go:embed tcp_client.toml
	TCPClientTemplate string
)
