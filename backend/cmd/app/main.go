package main

import (
	"os"
	"test-app/internal/app"
)

func main() {
	cfg := &app.Config{
		Env:  os.Getenv("ENV"),
		Port: os.Getenv("PORT"),

		DBHost: os.Getenv("DB_HOST"),
		DBUser: os.Getenv("DB_USER"),
		DBPass: os.Getenv("DB_PASS"),
		DBName: os.Getenv("DB_NAME"),

		JWTSecret: os.Getenv("JWT_SECRET"),
	}

	db, close := app.NewDatabase(cfg)
	defer close()

	srv, close := app.NewServer(cfg, db)
	defer close()

	if err := srv.Run(); err != nil {
		panic(err.Error())
	}
}
