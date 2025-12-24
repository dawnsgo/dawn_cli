package template

const AppOutput = "app/app.go"

const AppTemplate = `
package app

import "github.com/dawnsgo/dawn/cluster/mesh"

func Init(proxy *mesh.Proxy) {
	// TODO: init service
}
`
