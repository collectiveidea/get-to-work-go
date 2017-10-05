package prompts

import (
	"github.com/pivotal/gumshoe/trackerapi/domain"
	"github.com/segmentio/go-prompt"
)

// PivotalTracker prompts the user for pivotal tracker credentials
func PivotalTracker() (string, string) {
	println("Step #2: Pivotal Tracker Setup")
	println("------------------------------")
	email := prompt.String("email")
	password := prompt.PasswordMasked("password")

	return email, password
}

func pivotalTrackerProjectNames(projects []domain.Project) (names []string) {
	names = make([]string, len(projects))

	for i, v := range projects {
		names[i] = v.Name
	}

	return
}

// PivotalTrackerChooseProject prompts the user to choose a project
func PivotalTrackerChooseProject(projects []domain.Project) (proj domain.Project) {
	projectMenu := pivotalTrackerProjectNames(projects)
	selection := prompt.Choose("Choose a project", projectMenu)

	proj = projects[selection]
	return
}
