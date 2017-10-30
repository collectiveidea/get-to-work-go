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

		fmt.Print("\n")

		// Prompt for Harvest credentials
		harvest := service.NewHarvestService()
		accountID := cfg.Harvest.AccountID

		var token string

		prompts.HarvestBanner()
		if accountID == "" || !service.HasCredentials(harvest) {
			accountID, token = prompts.Harvest()
			cfg.Harvest.AccountID = accountID
			cfg.SaveDefaultConfig()
			service.SaveCredentials(harvest, token)
		} else {
			accountID = cfg.Harvest.AccountID
			token, err = service.LoadCredentials(harvest)
			if err != nil {
				println("Could not load Harvest credentials")
			}
		}

		err = harvest.SignIn(accountID, token)
		if err != nil {
			fmt.Println(err)
			println("Error: Harvest Authentication failed.")
		}

		projAssignment := prompts.HarvestChooseProject(harvest.GetProjects())
		cfg.Harvest.ProjectID = strconv.FormatInt(projAssignment.Project.ID, 10)
		cfg.SaveDefaultConfig()

		taskAssignments := harvest.GetTasks(projAssignment)
		taskAssignment := prompts.HarvestChooseTask(taskAssignments)

		cfg.Harvest.TaskID = strconv.FormatInt(taskAssignment.Task.ID, 10)
		cfg.SaveDefaultConfig()

		fmt.Print("\n\n")

		pt := service.NewPivotalTrackerService()

		prompts.PivotalTrackerBanner()
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
		projects, _ := pt.GetProjects()
		ptproj := prompts.PivotalTrackerChooseProject(projects)
		cfg.PivotalTracker.ProjectID = strconv.FormatInt(int64(ptproj.Id), 10)
		cfg.SaveDefaultConfig()

		return nil
	},
}
