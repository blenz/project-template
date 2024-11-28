package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type handler struct {
	service Service
}

func NewHandler(service Service) handler {
	return handler{
		service: service,
	}
}

func (h handler) RegisterRoutes(router *echo.Group, auth echo.MiddlewareFunc) {
	r := router.Group("/auth")
	r.POST("/login", h.login)
	r.GET("/logout", h.logout)
	r.GET("/session", auth(h.session))
}

func (h handler) session(c echo.Context) error {
	_, err := c.Cookie("token")
	if err != nil {
		return c.JSON(400, user{})
	}

	return c.JSON(200, user{})
}

func (h handler) login(c echo.Context) error {
	req := loginRequest{}
	if err := c.Bind(&req); err != nil {
		return err
	}

	token, err := h.service.Authenticate(req.Username, req.Password)
	if err != nil {
		return c.NoContent(400)
	}

	c.SetCookie(&http.Cookie{Name: "token", Value: token})

	return c.JSON(200, user{Username: req.Username})
}

func (h handler) logout(c echo.Context) error {
	c.SetCookie(&http.Cookie{Name: "token", MaxAge: -1})

	return c.NoContent(200)
}
