package service

import (
	"fmt"
	"os"
	"testing"

	"github.com/pivotal/gumshoe/trackerapi/domain"
	"github.com/stretchr/testify/assert"
)

func TestInitializePivotalTrackerService(t *testing.T) {
	pt := NewPivotalTrackerService()

	assert.NotNil(t, pt)
	assert.Equal(t, "PivotalTracker", pt.Name)
}

func TestPivotalTrackerGetName(t *testing.T) {
	pt := NewPivotalTrackerService()
	assert.Equal(t, "PivotalTracker", pt.GetName())
}

func TestPivotalTrackerSignIn(t *testing.T) {
	email := os.Getenv("pivotal_tracker_email")
	password := os.Getenv("pivotal_tracker_password")
	pt := NewPivotalTrackerService()
	err := pt.SignIn(email, password)

	assert.Nil(t, err)
	assert.Equal(t, true, pt.Client.IsAuthenticated())
}

func TestGetProjects(t *testing.T) {
	email := os.Getenv("pivotal_tracker_email")
	password := os.Getenv("pivotal_tracker_password")
	pt := NewPivotalTrackerService()
	pt.SignIn(email, password)

	fmt.Println(pt.GetProjects())
	projects := pt.GetProjects()
	assert.NotEmpty(t, projects)
	assert.IsType(t, domain.Project{}, projects[0])
}
