package service

import (
	"github.com/pivotal/gumshoe/trackerapi"
)

// PivotalTrackerService defines a harvest service
type PivotalTrackerService struct {
	Service
	Client *trackerapi.Client
}

func NewPivotalTrackerService() (pivotalTrackerService *PivotalTrackerService) {
	return &PivotalTrackerService{}
}

// SignIn signs a user into
func (pts *PivotalTrackerService) SignIn(email string, password string) {
	config := trackerapi.NewConfiguration()
	client, err := trackerapi.NewClient(config)

	if err == nil {
		pts.Client = client
		client.Authenticate(email, password)

		if !client.IsAuthenticated() {
			println("could not authenticate user with Pivotal Tracker")
		}
	}
}
