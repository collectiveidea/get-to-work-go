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
		var account_id string
		var token string

		if !service.HasCredentials(harvest) {
			account_id, token = prompts.Harvest()
			cfg.Harvest.AccountID = account_id
			cfg.SaveDefaultConfig()
			service.SaveCredentials(harvest, token)
		} else {
			account_id = cfg.Harvest.AccountID
			token, err = service.LoadCredentials(harvest)
			if err != nil {
				println("Could not load Harvest credentials")
			}
		}

		err = harvest.SignIn(account_id, token)
		if err != nil {
			fmt.Println(err)
			println("Error: Harvest Authentication failed.")
		}

		prj := prompts.HarvestChooseProject(harvest.GetProjects())
		cfg.Harvest.ProjectID = strconv.FormatInt(prj.ID, 10)
		cfg.SaveDefaultConfig()

		fmt.Print("\n\n")

		pt := service.NewPivotalTrackerService()

		if !service.HasCredentials(pt) {
			token = prompts.PivotalTracker()

			cfg.SaveDefaultConfig()
			service.SaveCredentials(pt, token)
		} else {
			token, err = service.LoadCredentials(pt)
			if err != nil {
				println("Could not load PivotalTracker credentials")
			}
		}

		pt.SignIn(token)
		projects, err := pt.GetProjects()
		ptproj := prompts.PivotalTrackerChooseProject(projects)
		cfg.PivotalTracker.ProjectID = strconv.FormatInt(int64(ptproj.Id), 10)
		cfg.SaveDefaultConfig()

		return nil
	},
}
