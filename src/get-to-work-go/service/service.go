package service

import (
	"fmt"
)

// Service defines a service used in go-to-work
type Service interface {
	Name() string
}

const namePrefix = "GoToWork"

// FullName returns the application namespaced name for the service
func FullName(s Service) string {
	return fmt.Sprintf("%s::%s", namePrefix, s.Name())
}
