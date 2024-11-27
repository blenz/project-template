package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type handler struct{}

func NewHandler() *handler {
	return &handler{}
}

func (h handler) RegisterRoutes(router *echo.Group) {
	r := router.Group("/auth")
	r.POST("/login", h.login)
	r.GET("/logout", h.logout)
	r.GET("/session", h.session)
}

func (h handler) session(c echo.Context) error {
	cookie, err := c.Cookie("session")
	if err != nil {
		return c.JSON(400, sessionResponse{})
	}

	resp := sessionResponse{}
	if cookie.Value == "123" {
		resp = sessionResponse{
			Username: "test",
			Token:    "test",
		}
	}

	return c.JSON(200, resp)
}

func (h handler) login(c echo.Context) error {
	req := loginRequest{}

	if err := c.Bind(&req); err != nil {
		return err
	}

	// simple auth for now
	if req.Username != "test" && req.Password != "test" {
		return c.NoContent(400)
	}

	c.SetCookie(&http.Cookie{
		Name:  "session",
		Value: "123",
	})

	return c.NoContent(200)
}

func (h handler) logout(c echo.Context) error {
	c.SetCookie(&http.Cookie{
		Name:   "session",
		MaxAge: -1,
	})
	return c.NoContent(200)
}
