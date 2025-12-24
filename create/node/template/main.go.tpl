package main

import (
	{{- if .Alone }}
    "{{.Module}}/app"
    {{- else}}
	"{{.Module}}/{{.Name}}/app"
    {{- end}}

	{{- if ne .Config "none"}}
	cc "github.com/dawnsgo/dawn/config/{{.Config}}/{{.DawnMajorVersion}}"
	{{- end}}
	"github.com/dawnsgo/dawn/locate/{{.Locate}}/{{.DawnMajorVersion}}"
	"github.com/dawnsgo/dawn/registry/{{.Registry}}/{{.DawnMajorVersion}}"
	{{- if ne .Transport "none"}}
	"github.com/dawnsgo/dawn/transport/{{.Transport}}/{{.DawnMajorVersion}}"
	{{- end}}
	"github.com/dawnsgo/dawn/{{.DawnMajorVersion}}"
	"github.com/dawnsgo/dawn/{{.DawnMajorVersion}}/cluster/node"
	{{- if ne .Config "none"}}
	"github.com/dawnsgo/dawn/config"
	{{- end}}
)

func main() {
	{{- if ne .Config "none"}}
	// 设置配置中心
	config.SetConfigurator(config.NewConfigurator(config.WithSources(cc.NewSource())))
	{{- end}}
	// 创建容器
	container := dawn.NewContainer()
	// 创建用户定位器
	locator := {{.Locate}}.NewLocator()
	// 创建服务注册发现
	registry := {{.Registry}}.NewRegistry()
	{{- if ne .Transport "none"}}
	// 创建RPC传输器
	transporter := {{.Transport}}.NewTransporter()
	{{- end}}
	// 创建节点组件
	component := node.NewNode(
		node.WithLocator(locator),
		node.WithRegistry(registry),
		{{- if ne .Transport "none"}}
		node.WithTransporter(transporter),
		{{- end}}
	)
	// 初始化应用
	app.Init(component.Proxy())
	// 添加节点组件
	container.Add(component)
	// 启动容器
	container.Serve()
}
