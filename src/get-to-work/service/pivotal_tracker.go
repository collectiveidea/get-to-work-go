package service

import (
	"get-to-work/config"

	"github.com/pivotal/gumshoe/trackerapi"
	"github.com/pivotal/gumshoe/trackerapi/domain"
	"github.com/pivotal/gumshoe/trackerapi/presenters"
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
func (pt *PivotalTrackerService) SignIn(email string, password string) (err error) {
	config := trackerapi.NewConfiguration()
	pt.Client, err = trackerapi.NewClient(config)

	if err == nil {
		pt.Client.Authenticate(email, password)

		if !pt.Client.IsAuthenticated() {
			println("could not authenticate user with Pivotal Tracker")
		}
	}

	return
}

// GetProjects returns projects
func (pt *PivotalTrackerService) GetProjects() (projects []domain.Project) {
	stringer := pt.Client.Projects()
	pres, ok := stringer.(presenters.Projects)

	if ok {
		projects = pres.Projects
	}

	return
}
