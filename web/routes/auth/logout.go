package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/markbates/goth/gothic"
)

func logout(c echo.Context) error {
	res := c.Response().Writer
	req := c.Request()

	gothic.Logout(res, req)
	res.Header().Set("Location", "/")
	res.WriteHeader(http.StatusTemporaryRedirect)
	return nil
}
