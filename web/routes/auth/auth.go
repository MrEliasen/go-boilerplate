package auth

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/github"
)

func Mount(g *echo.Group) {
	goth.UseProviders(
		github.New(
			os.Getenv("GITHUB_KEY"),
			os.Getenv("GITHUB_SECRET"),
			os.Getenv("APP_URL")+"/auth/github/callback",
		),
	)

	// bind the "provider" value to the request URL, so Goth has access to it
	bindProvider := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			q := c.Request().URL.Query()
			q.Add("provider", c.Param("provider"))
			c.Request().URL.RawQuery = q.Encode()
			return next(c)
		}
	}

	g.Use(bindProvider)

	g.GET("/:provider", login)
	g.GET("/logout/:provider", logout)
	g.GET("/:provider/callback", callback)
}
