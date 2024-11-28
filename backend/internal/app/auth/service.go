package auth

import (
	"errors"
)

type Service interface {
	Authenticate(username, password string) error
}

type service struct{}

func NewService() service {
	return service{}
}

func (s service) Authenticate(username, password string) error {
	// simple auth for now
	if !(username == "test" && password == "test") {
		return errors.New("invalid creds")
	}

	return nil
}
