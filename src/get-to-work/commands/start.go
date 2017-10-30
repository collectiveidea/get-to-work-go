package commands

import (
	"fmt"
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
		var ptStoryID int
		cfg, _ := config.DefaultConfig()

		harvest := service.NewHarvestService()
		token, _ = service.LoadCredentials(harvest)
		err = harvest.SignIn(cfg.Harvest.AccountID, token)

		pt := service.NewPivotalTrackerService()
		token, _ = service.LoadCredentials(pt)
		pt.SignIn(token)

		// Get pivotal tracker story id
		firstArg := c.Args().Get(0)
		if firstArg == "" {
			ptStoryID = cfg.PivotalTracker.LastStoryID
		} else {
			ptStoryID = GetPTStoryID(firstArg)
		}

		if ptStoryID == 0 {
			fmt.Println("\n\nCould not find a previously started story.\n Please pass a Pivotal Tracker Stroy URL as an argument")
			return
		}

		// The user passed in the argument
		projID, _ := strconv.Atoi(cfg.PivotalTracker.ProjectID)
		story := pt.GetStory(projID, ptStoryID)

		// Create a harvest time entry w/ the new content
		cfg.Harvest.LatTimeEntryID, err = harvest.StartTimer(cfg.Harvest.ProjectID, cfg.Harvest.TaskID, story.Name+"\n\n"+story.URL)
		if err != nil {
			fmt.Println(err)
			return
		}

		cfg.PivotalTracker.LastStoryID = ptStoryID
		cfg.SaveDefaultConfig()

		return nil

	},
}
