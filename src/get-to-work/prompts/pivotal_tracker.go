package prompts

import (
	prompt "github.com/segmentio/go-prompt"
	"gopkg.in/salsita/go-pivotaltracker.v1/v5/pivotal"
)

// PivotalTracker prompts the user for pivotal tracker credentials
func PivotalTracker() (token string) {
	println("Step #2: Pivotal Tracker Setup")
	println("------------------------------")
	token = prompt.String("token")
	return token
}

func pivotalTrackerProjectNames(projects []*pivotal.Project) (names []string) {
	names = make([]string, len(projects))

	for i, v := range projects {
		names[i] = v.Name
	}

	return
}

// PivotalTrackerChooseProject prompts the user to choose a project
func PivotalTrackerChooseProject(projects []*pivotal.Project) (proj *pivotal.Project) {
	projectMenu := pivotalTrackerProjectNames(projects)
	selection := prompt.Choose("Choose a project", projectMenu)

	proj = projects[selection]
	return
}
