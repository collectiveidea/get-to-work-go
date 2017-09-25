package prompts

import "github.com/segmentio/go-prompt"

// Harvest prompts a user for harvest credentials
func Harvest() (string, string, string) {
	println("Step #1: Harvest Setup")
	println("----------------------")
	subdomain := prompt.String("subdomain")
	email := prompt.String("email")
	password := prompt.PasswordMasked("password")

	return subdomain, email, password
}
