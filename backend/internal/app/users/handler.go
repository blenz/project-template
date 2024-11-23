package users

import (
	"github.com/labstack/echo/v4"
)

type handler struct {
	repo Repository
}

func NewHandler(repo Repository) *handler {
	return &handler{
		repo: repo,
	}
}

func (h handler) RegisterRoutes(router *echo.Echo) {
	r := router.Group("/users")
	r.POST("", h.create)
	r.GET("/:id", h.find)
	r.PATCH("/:id", h.update)
	r.DELETE("/:id", h.delete)
}

func (h handler) create(c echo.Context) error {
	return nil
}

func (h handler) find(c echo.Context) error {
	return nil
}

func (h handler) update(c echo.Context) error {
	return nil
}

func (h handler) delete(c echo.Context) error {
	return nil
}
