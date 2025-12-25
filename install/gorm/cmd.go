package gorm

import (
	"github.com/dawnsgo/dawn_cli/internal/exec"
	"github.com/dawnsgo/dawn_cli/internal/flag"
	"github.com/urfave/cli/v2"
)

var Command = &cli.Command{
	Name:  "gorm",
	Usage: "install the gorm toolchain",
	Flags: []cli.Flag{
		flag.Version,
	},
	Action: func(ctx *cli.Context) error {
		exec.Install(exec.Package{
			Name:    "gorm_dao_generator",
			Module:  "github.com/dawnsgo/gorm_dao_generator",
			Version: ctx.String(flag.Version.Name),
		})

		return nil
	},
}
