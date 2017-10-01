package service

import (
	"get-to-work/config"

	"github.com/adlio/harvest"
	"github.com/tmc/keyring"
)

const keyringService = "GetToWork::Harvest"

// HarvestService defines a harvest service
type HarvestService struct {
	Service
	User *harvest.User
}

// WhoAmIResponse defines the response from the /account/who_am_i endpoint
type WhoAmIResponse struct {
	User *harvest.User `json:"user"`
}

// NewHarvestService creates a HarvestService instance
func NewHarvestService() (harvestService *HarvestService) {
	return &HarvestService{}
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

// SaveCredentials persists credentials to the OSX keychain
func (hs *HarvestService) SaveCredentials(username string, password string) (err error) {
	err = keyring.Set(keyringService, username, password)
	return
}

// LoadCredentials returns the username and password for the harvest service
func (hs *HarvestService) LoadCredentials() (username string, password string, e error) {
	cfg, _ := config.DefaultConfig()
	username = cfg.Harvest.Username
	password, _ = keyring.Get(keyringService, username)

	return
}
