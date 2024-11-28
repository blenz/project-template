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
				return echo.NewHTTPError(http.StatusBadRequest, "Could not get session")
			}

			keyFunc := func(t *jwt.Token) (interface{}, error) { return []byte(jwtSecret), nil }

			if _, err = jwt.Parse(token.Value, keyFunc); err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, "Invalid token")
			}

			return next(c)
		}
	}
}
