package service

import (
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
	token := os.Getenv("pivotal_tracker_token")
	pt := NewPivotalTrackerService()
	err := pt.SignIn(token)

	assert.Nil(t, err)
}

func TestGetProjects(t *testing.T) {
	token := os.Getenv("pivotal_tracker_token")
	pt := NewPivotalTrackerService()
	pt.SignIn(token)

	projects, err := pt.GetProjects()
	assert.Nil(t, err)
	assert.NotEmpty(t, projects)
	assert.IsType(t, domain.Project{}, projects[0])
}
