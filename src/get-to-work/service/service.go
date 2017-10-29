package service

import (
	"fmt"

	"github.com/tmc/keyring"
)

// Service defines a service used in go-to-work
type Service interface {
	GetName() string
}

const namePrefix = "GetToWork"

// FullName returns the application namespaced name for the service
func FullName(s Service) string {
	return fmt.Sprintf("%s::%s", namePrefix, s.GetName())
}

// SaveCredentials persists credentials to the OSX keychain
func SaveCredentials(s Service, token string) (err error) {
	err = keyring.Set(namePrefix, s.GetName(), token)
	return
}

// LoadCredentials returns the token for the harvest service
func LoadCredentials(s Service) (token string, e error) {
	token, _ = keyring.Get(namePrefix, s.GetName())
	// What to do if the keychain doesn't exist?
	return
}

// HasCredentials returns true if a user's credentials have been set
func HasCredentials(s Service) (foundCredentials bool) {
	foundCredentials = false
	token, _ := LoadCredentials(s)

	if len(token) > 0 {
		foundCredentials = true
	}

	return
}
