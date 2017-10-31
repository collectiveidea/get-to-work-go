package commands

import (
	"fmt"
	"get-to-work/config"
	"get-to-work/prompts"
	"get-to-work/service"

	"strconv"

	"github.com/fatih/color"
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

			err = cfg.SaveDefaultConfig()
			if err != nil {
				color.Red("Could not save .get-to-work")
				return
			}

			err = service.SaveCredentials(harvest, token)
			if err != nil {
				color.Red("Could not Save Harvest credentials to your keychain.")
			}

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
			return
		}

		projAssignment := prompts.HarvestChooseProject(harvest.GetProjects())
		cfg.Harvest.ProjectID = strconv.FormatInt(projAssignment.Project.ID, 10)
		err = cfg.SaveDefaultConfig()

		if err != nil {
			color.Red("Could not save .get-to-work")
			return
		}

		taskAssignments := harvest.GetTasks(projAssignment)
		taskAssignment := prompts.HarvestChooseTask(taskAssignments)

		cfg.Harvest.TaskID = strconv.FormatInt(taskAssignment.Task.ID, 10)

		err = cfg.SaveDefaultConfig()
		if err != nil {
			color.Red("Could not save .get-to-work")
			return
		}

		fmt.Print("\n\n")

		pt := service.NewPivotalTrackerService()

		prompts.PivotalTrackerBanner()
		if !service.HasCredentials(pt) {
			token = prompts.PivotalTracker()

			err = cfg.SaveDefaultConfig()
			if err != nil {
				color.Red("Could not save .get-to-work")
				return
			}

			err = service.SaveCredentials(pt, token)
			if err != nil {
				color.Red("Could not save Pivotal Tracker credentials to your keychain.")
			}

		} else {
			token, err = service.LoadCredentials(pt)
			if err != nil {
				println("Could not load PivotalTracker credentials")
			}
		}

		err = pt.SignIn(token)
		if err != nil {
			color.Red("ERROR: Could not sign into Pivotal Tracker service")
			fmt.Println(err)
			return
		}

		projects, _ := pt.GetProjects()
		ptproj := prompts.PivotalTrackerChooseProject(projects)
		cfg.PivotalTracker.ProjectID = strconv.FormatInt(int64(ptproj.Id), 10)

		err = cfg.SaveDefaultConfig()
		if err != nil {
			color.Red("Could not save .get-to-work")
			return
		}

		return nil
	},
}
