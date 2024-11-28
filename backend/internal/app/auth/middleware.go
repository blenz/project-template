package auth

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func ValidateSession(jwtSecret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token, err := c.Cookie("token")
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
			}

			keyFunc := func(t *jwt.Token) (interface{}, error) { return []byte(jwtSecret), nil }

			if _, err = jwt.Parse(token.Value, keyFunc); err != nil {
				token.MaxAge = -1
				token.Path = "/"
				c.SetCookie(token)
				return echo.NewHTTPError(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
			}

			return next(c)
		}
	}
}
