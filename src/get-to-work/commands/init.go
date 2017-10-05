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
	Action: func(c *cli.Context) (err error) {
		// Create a configuration file
		cfg, _ := config.DefaultConfig()

		fmt.Print("\n\n")

		// Prompt for Harvest credentials
		harvest := service.NewHarvestService()
		var subdomain string
		var email string
		var password string

		if !service.HasCredentials(harvest) {
			subdomain, email, password = prompts.Harvest()
			cfg.Harvest.Subdomain = subdomain
			cfg.Harvest.Username = email
			cfg.SaveDefaultConfig()
			service.SaveCredentials(harvest, email, password)
		} else {
			subdomain = cfg.Harvest.Subdomain
			email, password, err = service.LoadCredentials(harvest)
			if err != nil {
				println("Could not load Harvest credentials")
			}
		}

		err = harvest.SignIn(subdomain, email, password)
		if err != nil {
			println("Error: Harvest Authentication failed.")
		}

		prj := prompts.HarvestChooseProject(harvest.GetProjects())
		cfg.Harvest.ProjectID = strconv.FormatInt(prj.ID, 10)
		cfg.SaveDefaultConfig()

		fmt.Print("\n\n")

		pt := service.NewPivotalTrackerService()

		if !service.HasCredentials(pt) {
			email, password = prompts.PivotalTracker()
			cfg.PivotalTracker.Username = email
			cfg.SaveDefaultConfig()

			service.SaveCredentials(pt, email, password)
		} else {
			email, password, err = service.LoadCredentials(pt)
			if err != nil {
				println("Could not load PivotalTracker credentials")
			}
		}

		pt.SignIn(email, password)

		ptproj := prompts.PivotalTrackerChooseProject(pt.GetProjects())
		cfg.PivotalTracker.ProjectID = strconv.FormatInt(int64(ptproj.ID), 10)
		cfg.SaveDefaultConfig()

		return nil
	},
}
