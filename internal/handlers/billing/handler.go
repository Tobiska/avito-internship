package billing

import "github.com/labstack/echo/v4"

type Handler struct {
	// todo env usecases
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Handle(c echo.Context) error {
	return c.String(200, "Hello world")
}

func (h *Handler) Register(e *echo.Echo) {
	e.GET("/hello-world", h.Handle)
}
