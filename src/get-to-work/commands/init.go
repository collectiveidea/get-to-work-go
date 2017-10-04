package commands

import (
	"fmt"
	"get-to-work/config"
	"get-to-work/prompts"
	"get-to-work/service"

	"strconv"

	"github.com/urfave/cli"
)

// Init prepares the project directory for use
var Init = cli.Command{
	Name:  "init",
	Usage: "Prepare the current project directory for go-to-work",
	Action: func(c *cli.Context) error {
		// Create a configuration file
		cfg, _ := config.DefaultConfig()

		fmt.Print("\n\n")

		// Prompt for Harvest credentials
		subdomain, email, password := prompts.Harvest()
		cfg.Harvest.Subdomain = subdomain
		cfg.Harvest.Username = email
		cfg.SaveDefaultConfig()

		harvest := service.NewHarvestService()
		err := harvest.SignIn(subdomain, email, password)

		prj := prompts.HarvestChooseProject(harvest.GetProjects())
		cfg.Harvest.ProjectID = strconv.FormatInt(prj.ID, 10)
		cfg.SaveDefaultConfig()

		if err != nil {
			println("Error: Harvest Authentication failed.")
		}

		service.SaveCredentials(harvest, email, password)

		fmt.Print("\n\n")

		email, password = prompts.PivotalTracker()
		cfg.PivotalTracker.Username = email
		cfg.SaveDefaultConfig()

		pt := service.NewPivotalTrackerService()
		pt.SignIn(email, password)
		service.SaveCredentials(pt, email, password)

		return nil
	},
}
