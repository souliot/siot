package urfave

import (
	"github.com/souliot/siot/internal/basic"

	"github.com/urfave/cli/v2"
)

var (
	CliCommand = &cli.Command{
		Name:  "cli",
		Usage: "generate a golang cli project",
		Action: func(c *cli.Context) error {
			return DefaultCliGenerate.Run(c)
		},
	}
)

var DefaultCliGenerate = new(Cli)

type Cli struct{}

func (c *Cli) Run(ctx *cli.Context) (err error) {
	basic.GenerateBasic("llz")
	return
}
