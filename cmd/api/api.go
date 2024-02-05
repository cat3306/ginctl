package api

import (
	"github.com/urfave/cli/v2"
)

type Cmd struct {
}

func (a *Cmd) CliCmd() *cli.Command {
	return &cli.Command{
		Name:        "api",
		Description: "",
		Usage:       "generate api related files",
		Subcommands: a.Subcommands(),
	}
}
func (a *Cmd) Subcommands() []*cli.Command {
	return []*cli.Command{
		a.goCmd(),
	}
}
func (a *Cmd) goCmd() *cli.Command {
	return &cli.Command{
		Name:        "go",
		Description: "",
		Usage:       "Generate go files for provided epi in api file",
		Action:      a.goAction,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "api",
				Required: true,
				Usage:    "The api file",
			},
			&cli.StringFlag{
				Name:     "dir",
				Required: true,
				Usage:    "The target dir",
			},
		},
	}
}
func (a *Cmd) goAction(ctx *cli.Context) error {
	return nil
}
