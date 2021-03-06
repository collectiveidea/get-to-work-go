package service

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/salsita/go-pivotaltracker.v1/v5/pivotal"
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
	err := pt.SignIn(token)
	if err != nil {
		t.Error(err)
	}

	projects, err := pt.GetProjects()
	if err != nil {
		t.Error(err)
	}
	assert.Nil(t, err)
	assert.NotEmpty(t, projects)
	assert.IsType(t, pivotal.Project{}, *projects[0])
}
