package app

import (
    {{- if .Alone }}
    "{{.Module}}/app/logic"
    {{- else}}
	"{{.Module}}/{{.Name}}/app/logic"
    {{- end}}

	"github.com/dawnsgo/dawn/cluster/node"
)

func Init(proxy *node.Proxy) {
	logic.NewCore(proxy).Init()
}
