package main

import (
	"os"

	"github.com/collectiveidea/get-to-work-go/commands"
	"github.com/urfave/cli"
)

const version = "0.0.2"

func main() {
	app := cli.NewApp()
	app.Name = "get-to-work"
	app.Version = version
	app.Description = "Keep track of what you're doing in Harvest"

	app.Commands = []cli.Command{
		commands.Init,
	}

	app.Run(os.Args)
}
