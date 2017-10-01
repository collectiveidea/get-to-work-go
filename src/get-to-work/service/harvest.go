package service

import (
	"github.com/adlio/harvest"
)

// HarvestService defines a harvest service
type HarvestService struct {
	Name string
	Service
	User *harvest.User
}

// WhoAmIResponse defines the response from the /account/who_am_i endpoint
type WhoAmIResponse struct {
	User *harvest.User `json:"user"`
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
	api := harvest.NewBasicAuthAPI(subdomain, email, password)
	res := WhoAmIResponse{}

	// Get the user
	err := api.Get(
		"/account/who_am_i",
		harvest.Defaults(),
		&res,
	)

	if err == nil {
		hs.User = res.User
	}

	return err
}
