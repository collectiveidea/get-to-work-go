package commands

import (
	"get-to-work/prompts"
	"get-to-work/service"
	"get-to-work/config"

	"github.com/urfave/cli"
)

// Init prepares the project directory for use
var Init = cli.Command{
	Name:  "init",
	Usage: "Prepare the current project directory for go-to-work",
	Action: func(c *cli.Context) error {
		// Create a configuration file
		cfgFile := ".get-to-work"
		cfg, _ := config.FromFile(cfgFile)

		// Prompt for Harvest credentials
		subdomain, email, password := prompts.Harvest()
		cfg.Harvest.Subdomain = subdomain
		cfgErr := cfg.Save(cfgFile)

		if cfgErr != nil {

		}
		harvest := service.NewHarvestService()
		err := harvest.SignIn(subdomain, email, password)



		if err != nil {
			println("Error: Harvest Authentication failed.")
		}

		email, password = prompts.PivotalTracker()
		pt := service.NewPivotalTrackerService()
		pt.SignIn(email, password)
		return nil
	},
}
