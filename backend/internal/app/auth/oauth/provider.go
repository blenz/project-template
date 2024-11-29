package oauth

import (
	"context"
	"fmt"
	"log"

	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

type Provider interface {
	GetAuthURL() string
	GetIdentity(code string) (string, error)
}

type provider struct {
	oauth2.Config
}

func NewCognitoProvider(clientID, clientSecret, redirectURL, issuerURL string) provider {
	p, err := oidc.NewProvider(context.Background(), issuerURL)
	if err != nil {
		log.Fatalf("Failed to create OIDC provider: %v", err)
	}

	return provider{
		oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RedirectURL:  redirectURL,
			Endpoint:     p.Endpoint(),
			Scopes:       []string{oidc.ScopeOpenID, "email", "openid", "profile"},
		},
	}
}

func (p provider) GetAuthURL() string {
	state := "state"
	return p.AuthCodeURL(state, oauth2.AccessTypeOffline)
}

func (p provider) GetIdentity(code string) (string, error) {
	token, err := p.Exchange(context.Background(), code)
	if err != nil {
		return "", err
	}

	fmt.Println(token.Extra("id_token"))
	return "", nil
}
