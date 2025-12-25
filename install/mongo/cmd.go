package mongo

import (
	"github.com/dawnsgo/dawn_cli/internal/exec"
	"github.com/dawnsgo/dawn_cli/internal/flag"
	"github.com/urfave/cli/v2"
)

var Command = &cli.Command{
	Name:  "mongo",
	Usage: "install the mongo toolchain",
	Flags: []cli.Flag{
		flag.Version,
	},
	Action: func(ctx *cli.Context) error {
		exec.Install(exec.Package{
			Name:    "mongo_dao_generator",
			Module:  "github.com/dawnsgo/mongo_dao_generatorr",
			Version: ctx.String(flag.Version.Name),
		})

		return nil
	},
}
