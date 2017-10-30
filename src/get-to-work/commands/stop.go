package commands

import (
	"fmt"
	"get-to-work/config"
	"get-to-work/service"

	"github.com/urfave/cli"
)

// Init prepares the project directory for use
var Stop = cli.Command{
	Name:  "stop",
	Usage: "Stop your last running timer",
	Action: func(c *cli.Context) (err error) {
		cfg, _ := config.DefaultConfig()
		harvest := service.NewHarvestService()
		token, _ := service.LoadCredentials(harvest)
		err = harvest.SignIn(cfg.Harvest.AccountID, token)

		if err != nil {
			return
		}

		entryID := cfg.Harvest.LatTimeEntryID

		if entryID == 0 {
			fmt.Println("\n\nNo existing time entry found. No timer has been stopped.")
			return
		}

		harvest.Stoptimer(entryID)
		return
	},
}
