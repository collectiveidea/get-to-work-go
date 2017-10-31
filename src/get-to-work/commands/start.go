package commands

import (
	"fmt"
	"get-to-work/config"
	"get-to-work/service"
	"regexp"
	"strconv"

	"github.com/fatih/color"
	"github.com/urfave/cli"
)

// GetPTStoryID returns a Pivotal Tracker story ID given a Pivotal Tracker story URL
func GetPTStoryID(URL string) int {
	r, _ := regexp.Compile("[0-9]+$")
	storyID, _ := strconv.Atoi(r.FindString(URL))

	return storyID
}

// Start prepares the project directory for use
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
		if err != nil {
			color.Red("ERROR: Could not sign into Harvest service")
			fmt.Println(err)
			return
		}

		pt := service.NewPivotalTrackerService()
		token, _ = service.LoadCredentials(pt)

		err = pt.SignIn(token)
		if err != nil {
			color.Red("ERROR: Could not sign into Pivotal Tracker service")
			fmt.Println(err)
			return
		}

		// Get pivotal tracker story id
		firstArg := c.Args().Get(0)
		if firstArg == "" {
			ptStoryID = cfg.PivotalTracker.LastStoryID
		} else {
			ptStoryID = GetPTStoryID(firstArg)
		}

		if ptStoryID == 0 {
			fmt.Println("")
			fmt.Println("")
			color.Red("Could not find a previously started story.\nPlease pass a Pivotal Tracker Stroy URL as an argument")
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
		err = cfg.SaveDefaultConfig()
		if err != nil {
			color.Red("Could not save .get-to-work")
			return
		}

		return nil

	},
}
