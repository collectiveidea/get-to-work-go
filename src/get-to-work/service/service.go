package service

import (
	"fmt"

	"github.com/tmc/keyring"
)

// Service defines a service used in go-to-work
type Service interface {
	GetName() string
	GetUsername() string
}

const namePrefix = "GetToWork"

// FullName returns the application namespaced name for the service
func FullName(s Service) string {
	return fmt.Sprintf("%s::%s", namePrefix, s.GetName())
}

// SaveCredentials persists credentials to the OSX keychain
func SaveCredentials(s Service, username string, password string) (err error) {
	err = keyring.Set(FullName(s), username, password)
	return
}

// LoadCredentials returns the username and password for the harvest service
func LoadCredentials(s Service) (username string, password string, e error) {
	username = s.GetUsername()
	password, _ = keyring.Get(FullName(s), username)
	// What to do if the keychain doesn't exist?
	return
}
