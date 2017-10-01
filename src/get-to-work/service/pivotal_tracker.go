package service

import (
	"get-to-work/config"

	"github.com/pivotal/gumshoe/trackerapi"
)

// PivotalTrackerService defines a harvest service
type PivotalTrackerService struct {
	Service
	Name   string
	Client *trackerapi.Client
}

// NewPivotalTrackerService returns a new instance of the pivotal tracker service
func NewPivotalTrackerService() (pivotalTrackerService *PivotalTrackerService) {
	pivotalTrackerService = &PivotalTrackerService{}
	pivotalTrackerService.Name = "PivotalTracker"

	return
}

// GetName returns the name value
func (pt *PivotalTrackerService) GetName() (name string) {
	name = pt.Name
	return
}

// GetUsername returns the cached username for the pivotaltracker service
func (pt PivotalTrackerService) GetUsername() (username string) {
	cfg, _ := config.DefaultConfig()
	username = cfg.PivotalTracker.Username

	return
}

// SignIn signs a user into
func (pt PivotalTrackerService) SignIn(email string, password string) {
	config := trackerapi.NewConfiguration()
	client, err := trackerapi.NewClient(config)

	if err == nil {
		pt.Client = client
		client.Authenticate(email, password)

		if !client.IsAuthenticated() {
			println("could not authenticate user with Pivotal Tracker")
		}
	}
}
