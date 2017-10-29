package service

import (
	"fmt"

	"github.com/adlio/harvest"
)

// HarvestService defines a harvest service
type HarvestService struct {
	Name string
	Service
	User *harvest.User
	API  *harvest.API
}

type ProjectAssignment struct {
	ID               int64            `json:"id,omitempty"`
	IsProjectManager bool             `json:"is_project_manager"`
	IsActive         bool             `json:"is_active"`
	Project          *harvest.Project `json:"project"`
	Client           *harvest.Client  `json:client`
}

type UserAssignmentsResponse struct {
	ProjectAssignments []*ProjectAssignment `json:"project_assignments"`
	PerPage            int64                `json:"per_page"`
	TotalPages         int64                `json:"total_pages"`
	TotalEntries       int64                `json:"total_entries"`
	NextPage           *int64               `json:"next_page"`
	PreviousPage       *int64               `json:"previous_page"`
	Page               int64                `json:"page"`
}

// WhoAmIResponse defines the response from the /account/who_am_i endpoint
type WhoAmIResponse struct {
	User *harvest.User `json:"user"`
}

// ProjectsResponse is a collection of projects returned from /daily
type ProjectsResponse struct {
	Projects []*harvest.Project `json:"projects"`
}

// NewHarvestService creates a HarvestService instance
func NewHarvestService() (harvestService *HarvestService) {
	harvestService = &HarvestService{Name: "Harvest"}
	return
}

// GetName returns the name value
func (hs *HarvestService) GetName() (name string) {
	name = hs.Name
	return
}

// SignIn signs a harvest user in
func (hs *HarvestService) SignIn(account_id string, token string) error {
	hs.API = harvest.NewTokenAPI(account_id, token)
	// Get the user
	user := harvest.User{}
	err := hs.API.Get("/users/me", harvest.Defaults(), &user)

	return err
}

// GetProjects returns projects
func (hs *HarvestService) GetProjects() (projects []*harvest.Project) {
	res := UserAssignmentsResponse{}
	err := hs.API.Get("/users/me/project_assignments", harvest.Defaults(), &res)

	for _, asg := range res.ProjectAssignments {
		projects = append(projects, asg.Project)
	}

	if err != nil {
		fmt.Println(err)
	}

	return
}
