package auth

import (
	"net/http"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

const SessionCookie = "session"

type handler struct {
	service Service
	// oauthProvider oauth.Provider
	sessionTTL time.Duration
}

func NewHandler(service Service, sessionTTL time.Duration) handler {
	return handler{
		service: service,
		// oauthProvider: oauthProvider,
		sessionTTL: sessionTTL,
	}
}

func (h handler) RegisterRoutes(router *echo.Group) {
	r := router.Group("/auth")
	r.POST("/login", h.login)
	r.GET("/logout", h.logout)
	r.GET("/session", h.session)
	// r.GET("/launch", h.launch)
	// r.GET("/callback", h.callback)
}

func (h handler) session(c echo.Context) error {
	sess, err := session.Get(SessionCookie, c)
	if err != nil {
		return err
	}
	return c.JSON(200, user{
		Username: sess.Values["username"].(string),
	})
}

func (h handler) login(c echo.Context) error {
	req := loginRequest{}
	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := h.service.Authenticate(req.Username, req.Password); err != nil {
		return c.NoContent(400)
	}

	sess, err := session.Get(SessionCookie, c)
	if err != nil {
		return err
	}

	sess.Values["username"] = req.Username

	if err := h.setSession(c, sess); err != nil {
		return err
	}

	return c.JSON(200, user{Username: req.Username})
}

func (h handler) logout(c echo.Context) error {
	sess, err := session.Get(SessionCookie, c)
	if err != nil {
		return err
	}

	sess.Options.MaxAge = -1
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return err
	}

	return c.NoContent(204)
}

// func (h handler) launch(c echo.Context) error {
// 	url := h.oauthProvider.GetAuthURL()
// 	return c.Redirect(http.StatusFound, url)
// }

// func (h handler) callback(c echo.Context) error {
// 	code := c.QueryParam("code")

// 	username, err := h.oauthProvider.GetIdentity(code)
// 	if err != nil {
// 		return err
// 	}

// 	sess, err := session.Get(SessionCookie, c)
// 	if err != nil {
// 		return err
// 	}

// 	sess.Values["username"] = username

// 	if err := h.setSession(c, sess); err != nil {
// 		return err
// 	}

// 	return c.Redirect(http.StatusFound, "/")
// }

func (h handler) setSession(c echo.Context, sess *sessions.Session) error {
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   int(h.sessionTTL.Seconds()),
		SameSite: http.SameSiteLaxMode,
		Secure:   true,
		HttpOnly: true,
	}

	return sess.Save(c.Request(), c.Response())
}
