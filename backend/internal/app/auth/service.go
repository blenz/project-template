package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type Service interface {
	Authenticate(username, password string) (string, error)
}

type service struct {
	jwtSecret string
}

func NewService(jwtSecret string) service {
	return service{
		jwtSecret: jwtSecret,
	}
}

func (s service) Authenticate(username, password string) (string, error) {
	// simple auth for now
	if !(username == "test" && password == "test") {
		return "", errors.New("invalid creds")
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Minute).Unix(),
		},
	)

	return token.SignedString([]byte(s.jwtSecret))
}
