package auth

import (
	"net/http"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

const sessionCookie = "session"

type handler struct {
	service    Service
	sessionTTL time.Duration
}

func NewHandler(service Service, sessionTTL time.Duration) handler {
	return handler{
		service: service,
	}
}

func (h handler) RegisterRoutes(router *echo.Group) {
	r := router.Group("/auth")
	r.POST("/login", h.login)
	r.GET("/logout", h.logout)
	r.GET("/session", h.session)
}

func (h handler) session(c echo.Context) error {
	sess, err := session.Get(sessionCookie, c)
	if err != nil || sess.IsNew {
		return err
	}
	if sess.IsNew {
		return c.NoContent(http.StatusUnauthorized)
	}
	return c.JSON(200, user{})
}

func (h handler) login(c echo.Context) error {
	req := loginRequest{}
	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := h.service.Authenticate(req.Username, req.Password); err != nil {
		return c.NoContent(400)
	}

	sess, err := session.Get(sessionCookie, c)
	if err != nil {
		return err
	}
	sess.Options = &sessions.Options{
		Path:   "/",
		MaxAge: int(h.sessionTTL.Seconds()),
	}

	sess.Values["id"] = "123"

	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return err
	}

	return c.JSON(200, user{Username: req.Username})
}

func (h handler) logout(c echo.Context) error {
	sess, err := session.Get(sessionCookie, c)
	if err != nil {
		return err
	}

	sess.Options.MaxAge = -1
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return err
	}

	return c.NoContent(204)
}
