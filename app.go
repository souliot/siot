package main

import (
	"os"
	"siot/internal/urfave"

	"github.com/urfave/cli/v2"
)

type App struct {
	Name    string
	Usage   string
	Version string
}

func (c *App) run() (err error) {
	app := cli.NewApp()
	app.Name = c.Name
	app.Usage = c.Usage
	app.Version = c.Version
	app.Commands = []*cli.Command{
		urfave.CliCommand,
	}
	app.Flags = []cli.Flag{}
	err = app.Run(os.Args)
	if err != nil {
		return
	}
	return
}
