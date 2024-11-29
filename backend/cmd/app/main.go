package main

import (
	"os"
	"test-app/internal/app"
	"time"
)

func main() {
	cfg := &app.Config{
		Env:  os.Getenv("ENV"),
		Port: os.Getenv("PORT"),

		DBHost: os.Getenv("DB_HOST"),
		DBUser: os.Getenv("DB_USER"),
		DBPass: os.Getenv("DB_PASS"),
		DBName: os.Getenv("DB_NAME"),

		SessionSecret: os.Getenv("SESSION_SECRET"),
		SessionTTL: func() time.Duration {
			if ttl, err := time.ParseDuration(os.Getenv("SESSION_TTL")); err == nil {
				return ttl
			}
			return 5 * time.Minute
		}(),

		CognitoClientId:     os.Getenv("COGNITO_CLIENT_ID"),
		CognitoClientSecret: os.Getenv("COGNITO_CLIENT_SECRET"),
		CognitoRedirectUrl:  os.Getenv("COGNITO_REDIRECT_URL"),
		CognitoIssuerUrl:    os.Getenv("COGNITO_ISSUER_URL"),
	}

	db, close := app.NewDatabase(cfg)
	defer close()

	srv, close := app.NewServer(cfg, db)
	defer close()

	if err := srv.Run(); err != nil {
		panic(err.Error())
	}
}
