package service

import (
	"os"
	"testing"

	"github.com/adlio/harvest"
	"github.com/stretchr/testify/assert"
)

func TestSignIn(t *testing.T) {
	subdomain := os.Getenv("harvest_subdomain")
	email := os.Getenv("harvest_email")
	password := os.Getenv("harvest_password")

	if subdomain == "" || email == "" || password == "" {
		t.Skip("Configure `harvest_email` and `harvest_password` as environment variables to run this test")
		return
	}

	h := NewHarvestService()
	h.SignIn(subdomain, email, password)

	assert.Equal(t, h.User.Email, "chris.rittersdorf@collectiveidea.com")
}

func TestProjects(t *testing.T) {
	subdomain := os.Getenv("harvest_subdomain")
	email := os.Getenv("harvest_email")
	password := os.Getenv("harvest_password")

	h := NewHarvestService()
	h.SignIn(subdomain, email, password)

	projects := h.GetProjects()
	assert.NotEmpty(t, projects)
	assert.IsType(t, &harvest.Project{}, projects[0])
}
