package renderer

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Templ(c echo.Context, t templ.Component) error {
	return t.Render(c.Request().Context(), c.Response().Writer)
}
