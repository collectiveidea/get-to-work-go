package service

import (
	"get-to-work/config"

	"gopkg.in/salsita/go-pivotaltracker.v1/v5/pivotal"
)

// PivotalTrackerService defines a harvest service
type PivotalTrackerService struct {
	Service
	Name   string
	Client *pivotal.Client
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
func (pt *PivotalTrackerService) SignIn(token string) (err error) {
	pt.Client = pivotal.NewClient(token)
	_, _, err = pt.Client.Me.Get()

	if err == nil {
		println("could not authenticate user with Pivotal Tracker")
	}

	return
}

// GetProjects returns projects
func (pt *PivotalTrackerService) GetProjects() (projects []*pivotal.Project, err error) {
	projects, _, err = pt.Client.Projects.List()

	return
}
