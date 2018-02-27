package commands

import (
	"fmt"
	"github.com/collectiveidea/get-to-work-go/config"
	"github.com/collectiveidea/get-to-work-go/service"

	"github.com/fatih/color"
	"github.com/urfave/cli"
)

// Stop stops a running timer
var Stop = cli.Command{
	Name:  "stop",
	Usage: "Stop your last running timer",
	Action: func(c *cli.Context) (err error) {
		cfg, _ := config.DefaultConfig()
		harvest := service.NewHarvestService()
		token, _ := service.LoadCredentials(harvest)

		err = harvest.SignIn(cfg.Harvest.AccountID, token)
		if err != nil {
			color.Red("Could not sign in to Harvest.")
			return
		}

		entryID := cfg.Harvest.LatTimeEntryID
		if entryID == 0 {
			fmt.Println("\n\nNo existing time entry found. No timer has been stopped.")
			return
		}

		err = harvest.Stoptimer(entryID)
		if err != nil {
			color.Red("Could not stop timer %s", entryID)
		}

		return
	},
}
