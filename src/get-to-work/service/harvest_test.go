package service

import (
	"os"
	"testing"

	"github.com/adlio/harvest"
	"github.com/stretchr/testify/assert"
)

func TestSignIn(t *testing.T) {
	accountID := os.Getenv("harvest_account_id")
	token := os.Getenv("harvest_token")

	if accountID == "" || token == "" {
		t.Skip("Configure `harvest_account_id` and `harvest_token` as environment variables to run this test")
		return
	}

	h := NewHarvestService()
	err := h.SignIn(accountID, token)

	assert.Nil(t, err)
}

func TestProjects(t *testing.T) {
	accountID := os.Getenv("harvest_account_id")
	token := os.Getenv("harvest_token")

	h := NewHarvestService()
	h.SignIn(accountID, token)

	projects := h.GetProjects()
	assert.NotEmpty(t, projects)
	assert.IsType(t, &harvest.Project{}, projects[0])
}
