package app

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	RegisterRoutes(parent *echo.Group)
}

type handler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) Handler {
	return &handler{
		db: db,
	}
}

func (h handler) RegisterRoutes(router *echo.Group) {
	router.GET("/health", h.health)
}

func (h handler) health(c echo.Context) error {
	code := 200
	if err := h.db.Ping(); err != nil {
		code = 400
	}
	return c.NoContent(code)
}
