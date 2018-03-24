package prompts

import (
	"fmt"
	hs "github.com/collectiveidea/get-to-work-go/service"

	"github.com/fatih/color"
	"github.com/manlycode/harvest"
	"github.com/segmentio/go-prompt"
)

// HarvestBanner shows the banner for the harvest init Step
func HarvestBanner() {
	color.Cyan("Step #1: Harvest Setup")
	color.Cyan("----------------------")
}

// Harvest prompts a user for harvest credentials
func Harvest() (accountID string, token string) {
	color.Green("Sign into Harvest and create a new Personal Access Token")
	color.Green("by visiting https://id.getharvest.com/oauth2/access_tokens/new")
	color.Green("")
	color.Green("Then copy and paste the following information:")

	accountID = prompt.String("Account ID")
	token = prompt.String("Your Token")

	return
}

func harvestProjectNames(projects []*hs.ProjectAssignment) (names []string) {
	names = make([]string, len(projects))

	for i, v := range projects {
		names[i] = fmt.Sprintf("%s - %s", v.Client.Name, v.Project.Name)
	}

	return
}

func harvestTaskNames(tasks []*harvest.TaskAssignment) (names []string) {
	names = make([]string, len(tasks))

	for i, v := range tasks {
		names[i] = v.Task.Name
	}

	return
}

// HarvestChooseProject prompts the user to choose a project
func HarvestChooseProject(projects []*hs.ProjectAssignment) (proj *hs.ProjectAssignment) {
	projectMenu := harvestProjectNames(projects)
	selection := prompt.Choose("Choose a project", projectMenu)

	proj = projects[selection]
	return
}

// HarvestChooseTask prompts the user to choose a task
func HarvestChooseTask(tasks []*harvest.TaskAssignment) (task *harvest.TaskAssignment) {
	taskMenu := harvestTaskNames(tasks)
	selection := prompt.Choose("Choose a task", taskMenu)

	task = tasks[selection]
	return
}
