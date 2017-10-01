package commands

import (
	"get-to-work/config"
	"get-to-work/prompts"
	"get-to-work/service"

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
		cfg.Harvest.Username = email
		cfg.Save(cfgFile)

		harvest := service.NewHarvestService()
		err := harvest.SignIn(subdomain, email, password)

		if err != nil {
			println("Error: Harvest Authentication failed.")
		}

		harvest.SaveCredentials(email, password)

		email, password = prompts.PivotalTracker()
		pt := service.NewPivotalTrackerService()
		pt.SignIn(email, password)
		return nil
	},
}
