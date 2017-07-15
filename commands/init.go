package commands

import (
	"os"

	"github.com/collectiveidea/get-to-work-go/prompts"
	"github.com/collectiveidea/get-to-work-go/service"
	"github.com/urfave/cli"
)

// Init prepares the project directory for use
var Init = cli.Command{
	Name:  "init",
	Usage: "Prepare the current project directory for go-to-work",
	Action: func(c *cli.Context) error {
		// Create a configuration file
		os.Create(".get-to-work")

		// Prompt for Harvest credentials
		subdomain, email, password := prompts.Harvest()
		harvest := service.NewHarvestService()
		err := harvest.SignIn(subdomain, email, password)

		if err != nil {
			println("Error: Harvest Authentication failed.")
		}

		println("")
		email, password = prompts.PivotalTracker()
		pt := service.NewPivotalTrackerService()
		pt.SignIn(email, password)

		return nil
	},
}
