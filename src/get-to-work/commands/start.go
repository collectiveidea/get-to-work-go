package commands

import (
	"get-to-work/config"
	"get-to-work/service"
	"regexp"
	"strconv"

	"github.com/urfave/cli"
)

func GetPTStoryID(URL string) int {
	r, _ := regexp.Compile("[0-9]+$")
	storyID, _ := strconv.Atoi(r.FindString(URL))

	return storyID
}

// Init prepares the project directory for use
var Start = cli.Command{
	Name:  "start",
	Usage: "Start a timer",
	Action: func(c *cli.Context) (err error) {
		// Create a configuration file
		var token string
		cfg, _ := config.DefaultConfig()

		harvest := service.NewHarvestService()
		token, _ = service.LoadCredentials(harvest)
		err = harvest.SignIn(cfg.Harvest.AccountID, token)

		pt := service.NewPivotalTrackerService()
		token, _ = service.LoadCredentials(pt)
		pt.SignIn(token)

		// Get pivotal tracker story id
		ptStoryID := GetPTStoryID(c.Args().Get(0))
		projID, _ := strconv.Atoi(cfg.PivotalTracker.ProjectID)
		story := pt.GetStory(projID, ptStoryID)

		// Create a harvest time entry w/ the new content
		harvest.StartTimer(cfg.Harvest.ProjectID, cfg.Harvest.TaskID, story.Name+"\n\n"+story.URL)

		return nil
	},
}
