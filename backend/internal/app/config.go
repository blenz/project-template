package app

import "time"

type Config struct {
	Env  string
	Port string

	DBHost string
	DBUser string
	DBPass string
	DBName string

	SessionTTL    time.Duration
	SessionSecret string

	CognitoClientId     string
	CognitoClientSecret string
	CognitoRedirectUrl  string
	CognitoIssuerUrl    string
}
