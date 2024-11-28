package app

import "time"

type Config struct {
	Env  string
	Port string

	DBHost string
	DBUser string
	DBPass string
	DBName string

	JWTSecret  string
	SessionTTL time.Duration
}
