package main

import (
	"fmt"
	"os"

	"github.com/cat3306/ginctl/cmd"
	"github.com/cat3306/ginctl/cmd/api"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:        "ginctl",
		Usage:       "ginctl [command]",
		Description: "A cli tool to generate gin api, grpc, model code",
		Commands:    make([]*cli.Command, 0),
	}
	registerCommand(app, new(api.Cmd))
	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}

func registerCommand(app *cli.App, c cmd.Command) {
	app.Commands = append(app.Commands, c.CliCmd())
}
