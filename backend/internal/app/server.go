package app

import (
	"database/sql"
	"test-app/internal/app/auth"
	"test-app/internal/app/auth/oauth"
	"test-app/internal/app/users"

	"github.com/labstack/echo/v4"
)

type Server interface {
	Run() error
}

type server struct {
	rtr *echo.Echo
	cfg *Config
	db  *sql.DB
}

func NewServer(cfg *Config, db *sql.DB) (server, func()) {
	srv := server{
		rtr: echo.New(),
		db:  db,
		cfg: cfg,
	}

	api := srv.rtr.Group("/api")
	api.Use(initMiddleware(cfg)...)

	for _, handler := range []Handler{
		NewHandler(srv.db),

		auth.NewHandler(
			auth.NewService(),
			oauth.NewCognitoProvider(cfg.CognitoClientId, cfg.CognitoClientSecret, cfg.CognitoRedirectUrl, cfg.CognitoIssuerUrl),
			cfg.SessionTTL,
		),

		users.NewHandler(users.NewRepository(srv.db)),
	} {
		handler.RegisterRoutes(api)
	}

	return srv, func() { srv.rtr.Close() }
}

func (s server) Run() error {
	return s.rtr.Start(":" + s.cfg.Port)
}
