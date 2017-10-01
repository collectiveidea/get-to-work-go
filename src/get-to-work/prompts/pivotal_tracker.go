package prompts

import "github.com/segmentio/go-prompt"

// PivotalTracker prompts the user for pivotal tracker credentials
func PivotalTracker() (string, string) {
	println("Step #2: Pivotal Tracker Setup")
	println("------------------------------")
	email := prompt.String("email")
	password := prompt.PasswordMasked("password")

	return email, password
}
