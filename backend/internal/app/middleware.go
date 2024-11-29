package app

import (
	"slices"
	"test-app/internal/app/auth"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func initMiddleware(cfg *Config) []echo.MiddlewareFunc {
	return []echo.MiddlewareFunc{
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     []string{"http://localhost:3000"},
			AllowCredentials: true,
		}),

		session.Middleware(sessions.NewCookieStore([]byte(cfg.SessionSecret))),

		middleware.CSRFWithConfig(middleware.CSRFConfig{
			CookiePath: "/",
		}),

		middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
			KeyLookup: "cookie:" + auth.SessionCookie,
			Validator: func(key string, c echo.Context) (bool, error) {
				sess, err := session.Get(auth.SessionCookie, c)
				return !sess.IsNew, err
			},
			ErrorHandler: func(err error, c echo.Context) error {
				sess, _ := session.Get(auth.SessionCookie, c)
				sess.Options.MaxAge = -1
				sess.Save(c.Request(), c.Response())
				return echo.ErrUnauthorized
			},
			Skipper: func(c echo.Context) bool {
				return slices.Contains([]string{
					"/api/auth/login",
					"/api/auth/launch",
					"/api/auth/callback",
				}, c.Path())
			},
		}),
	}
}
