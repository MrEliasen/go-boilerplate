package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/placeholder/boiler/web/routes/auth"
	"github.com/placeholder/boiler/web/routes/home"
)

func Mount(e *echo.Echo) {
	e.GET("/", home.Route)

	// auth endpoints
	auth.Mount(e.Group("/auth"))
}
