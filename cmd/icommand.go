package cmd

import "github.com/urfave/cli/v2"

type Command interface {
	CliCmd() *cli.Command
}
