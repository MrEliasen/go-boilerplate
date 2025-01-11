package fail

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Forbidden(c echo.Context, msg string) error {
	return c.String(http.StatusForbidden, msg)
}
