package fail

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Unauthorized(c echo.Context, msg string) error {
	return c.String(http.StatusBadRequest, msg)
}
