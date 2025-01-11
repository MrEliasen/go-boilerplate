package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/markbates/goth/gothic"
	"github.com/placeholder/boiler/web/routes/fail"
)

func callback(c echo.Context) error {
	user, err := gothic.CompleteUserAuth(c.Response().Writer, c.Request())
	if err != nil {
		return fail.BadRequest(c, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}
