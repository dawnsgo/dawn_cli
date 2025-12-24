package install

import (
	"github.com/dawnsgo/dawn_cli/install/gorm"
	"github.com/dawnsgo/dawn_cli/install/grpc"
	"github.com/dawnsgo/dawn_cli/install/mongo"
	"github.com/dawnsgo/dawn_cli/install/proto"
	"github.com/dawnsgo/dawn_cli/install/rpcx"
	"github.com/urfave/cli/v2"
)

var Command = &cli.Command{
	Name:        "install",
	Usage:       "install and update the toolchain",
	Description: "install and update the toolchain",
	Subcommands: []*cli.Command{
		grpc.Command,
		rpcx.Command,
		gorm.Command,
		mongo.Command,
		proto.Command,
	},
}
