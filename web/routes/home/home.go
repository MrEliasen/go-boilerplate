package home

import (
	"github.com/labstack/echo/v4"
	"github.com/placeholder/boiler/web/renderer"
	"github.com/placeholder/boiler/web/views"
)

func Route(c echo.Context) error {
	return renderer.Templ(c, views.Home("Home", map[string]string{
		"github": "GitHub",
	}))
}
