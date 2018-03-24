package commands

import (
	"fmt"
	"github.com/collectiveidea/get-to-work-go/config"
	"github.com/fatih/color"
	"github.com/urfave/cli"
	"os/exec"
)

// Open prepares the project directory for use
var Open = cli.Command{
	Name:  "open",
	Usage: "Open a web browser showing the current story you're working on",
	Action: func(c *cli.Context) (err error) {
		cfg, err := config.DefaultConfig()
		if err != nil {
			color.Red("Could not open configuration file")
			return
		}

		// How do we get a subdomain for Harvest?
		ptURL := fmt.Sprintf("https://www.pivotaltracker.com/n/projects/%s", cfg.PivotalTracker.ProjectID)
		cmd := exec.Command("open", ptURL)
		cmd.Run()
		return
	},
}
