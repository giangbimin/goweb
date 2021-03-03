package giangweb

import (
	"net/http"

	"github.com/labstack/echo"
)

func showHome(c echo.Context) error {
	return c.String(http.StatusOK, "damn good!")
}
