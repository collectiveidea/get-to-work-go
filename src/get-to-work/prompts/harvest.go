package prompts

import (
	"github.com/adlio/harvest"
	"github.com/segmentio/go-prompt"
)

// Harvest prompts a user for harvest credentials
func Harvest() (account_id string, token string) {
	println("Step #1: Harvest Setup")
	println("----------------------")
	println("Sign into Harvest and create a new Personal Access Token")
	println("by visiting https://id.getharvest.com/oauth2/access_tokens/new")
	println("")
	println("Then copy and paste the following information:")

	account_id = prompt.String("Account ID")
	token = prompt.String("Your Token")

	return
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
