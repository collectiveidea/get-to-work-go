package main

import (
	"os"

	"get-to-work/commands"

	"github.com/urfave/cli"
)

const version = "0.0.3"

func main() {
	app := cli.NewApp()
	app.Name = "get-to-work"
	app.Version = version
	app.Description = "Keep track of what you're doing in Harvest"

	app.Commands = []cli.Command{
		commands.Init,
		commands.Start,
		commands.Stop,
	}

	app.Run(os.Args)
}
