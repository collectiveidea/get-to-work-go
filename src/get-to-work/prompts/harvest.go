package prompts

import (
	"github.com/adlio/harvest"
	"github.com/segmentio/go-prompt"
)

// Harvest prompts a user for harvest credentials
func Harvest() (string, string, string) {
	println("Step #1: Harvest Setup")
	println("----------------------")
	subdomain := prompt.String("subdomain")
	email := prompt.String("email")
	password := prompt.PasswordMasked("password")

	return subdomain, email, password
}

func harvestProjectNames(projects []*harvest.Project) (names []string) {
	names = make([]string, len(projects))

	for i, v := range projects {
		names[i] = v.Name
	}

	return
}

// HarvestChooseProject prompts the user to choose a project
func HarvestChooseProject(projects []*harvest.Project) (proj *harvest.Project) {
	projectMenu := harvestProjectNames(projects)
	selection := prompt.Choose("Choose a project", projectMenu)

	proj = projects[selection]
	return
}
