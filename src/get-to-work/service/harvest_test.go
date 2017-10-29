package service

import (
	"os"
	"testing"

	"github.com/adlio/harvest"
	"github.com/stretchr/testify/assert"
)

func TestSignIn(t *testing.T) {
	subdomain := os.Getenv("harvest_subdomain")
	token := os.Getenv("harvest_token")

	if subdomain == "" || token == "" {
		t.Skip("Configure `harvest_email` and `harvest_password` as environment variables to run this test")
		return
	}

	h := NewHarvestService()
	h.SignIn(subdomain, token)

	assert.Equal(t, h.User.Email, "chris.rittersdorf@collectiveidea.com")
}

func TestProjects(t *testing.T) {
	subdomain := os.Getenv("harvest_subdomain")
	token := os.Getenv("harvest_token")

	h := NewHarvestService()
	h.SignIn(subdomain, token)

	projects := h.GetProjects()
	assert.NotEmpty(t, projects)
	assert.IsType(t, &harvest.Project{}, projects[0])
}
