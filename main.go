package main

import (
	"fmt"
	"os"

	"github.com/dawnsgo/dawn_cli/create"
	"github.com/dawnsgo/dawn_cli/install"
	"github.com/dawnsgo/dawn_cli/internal/version"
	"github.com/dawnsgo/dawn_cli/upgrade"
	"github.com/urfave/cli/v2"
)

func main() {
	cli.VersionFlag = &cli.BoolFlag{
		Name:               "version",
		Aliases:            []string{"v"},
		Usage:              "show the duc-cli version information",
		DisableDefaultText: true,
	}

	app := &cli.App{
		Name:        "dawn_cli",
		Version:     version.ToolVersion,
		Description: "dawn project building tool",
		Authors: []*cli.Author{{
			Name:  "jstin",
			Email: "jstinapoll@gmail.com",
		}},
		Commands: []*cli.Command{
			create.Command,
			install.Command,
			upgrade.Command,
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
