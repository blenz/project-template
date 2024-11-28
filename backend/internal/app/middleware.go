package app

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func initMiddleware(cfg *Config) []echo.MiddlewareFunc {
	return []echo.MiddlewareFunc{
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     []string{"http://localhost:4000"},
			AllowCredentials: true,
		}),

		session.Middleware(sessions.NewCookieStore([]byte(cfg.SessionSecret))),

		middleware.CSRFWithConfig(middleware.CSRFConfig{
			CookiePath: "/",
		}),

		middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
			KeyLookup: "cookie:session",
			Validator: func(key string, c echo.Context) (bool, error) {
				sess, err := session.Get("session", c)
				return !sess.IsNew, err
			},
			ErrorHandler: func(err error, c echo.Context) error {
				return echo.ErrUnauthorized
			},
			Skipper: func(c echo.Context) bool {
				return c.Path() == "/api/auth/login"
			},
		}),
	}
}
