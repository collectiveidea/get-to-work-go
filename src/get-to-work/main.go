package main

import (
	"fmt"
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
		commands.Open,
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
