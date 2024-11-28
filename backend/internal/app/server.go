package app

import (
	"database/sql"
	"test-app/internal/app/auth"
	"test-app/internal/app/users"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	api.Use(
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     []string{"http://localhost:4000"},
			AllowCredentials: true,
		}),
		session.Middleware(sessions.NewCookieStore([]byte(cfg.SessionSecret))),
	)

	for _, handler := range []Handler{
		NewHandler(srv.db),

		auth.NewHandler(auth.NewService(), cfg.SessionTTL),

		users.NewHandler(users.NewRepository(srv.db)),
	} {
		handler.RegisterRoutes(api)
	}

	return srv, func() { srv.rtr.Close() }
}

func (s server) Run() error {
	return s.rtr.Start(":" + s.cfg.Port)
}
