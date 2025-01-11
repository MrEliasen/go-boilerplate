package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/markbates/goth/gothic"
)

func login(c echo.Context) error {
	res := c.Response().Writer
	req := c.Request()

	// try to get the user without re-authenticating
	if gothUser, err := gothic.CompleteUserAuth(res, req); err == nil {
		return c.JSON(http.StatusOK, gothUser)
	}

	gothic.BeginAuthHandler(res, req)
	return nil
}
