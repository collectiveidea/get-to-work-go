package service

import (
	"get-to-work/config"

	"github.com/adlio/harvest"
)

// HarvestService defines a harvest service
type HarvestService struct {
	Name string
	Service
	User *harvest.User
	API  *harvest.API
}

// GetUsername returns the configured username
func (hs *HarvestService) GetUsername() (username string) {
	cfg, _ := config.DefaultConfig()
	username = cfg.Harvest.Username
	return
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
func (hs *HarvestService) SignIn(subdomain string, email string, password string) error {
	hs.API = harvest.NewBasicAuthAPI(subdomain, email, password)
	res := WhoAmIResponse{}

	// Get the user
	err := hs.API.Get(
		"/account/who_am_i",
		harvest.Defaults(),
		&res,
	)

	if err == nil {
		hs.User = res.User
	}

	return err
}

// GetProjects returns projects
func (hs *HarvestService) GetProjects() (projects []*harvest.Project) {
	pr := ProjectsResponse{}
	hs.API.Get("/daily", harvest.Defaults(), &pr)
	projects = pr.Projects

	return
}
