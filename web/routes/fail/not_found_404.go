package fail

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func NotFound(c echo.Context, msg string) error {
	return c.String(http.StatusNotFound, msg)
}
