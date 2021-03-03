package goweb

import (
	"net/http"

	"github.com/labstack/echo"
)

func indexItems(c echo.Context) error {
	if c.QueryParam("json") != "" {
		return c.JSON(http.StatusOK, []string{"a", "b"})
	}
	return c.String(http.StatusOK, "list of Item")
}

func showItem(c echo.Context) error {
	return c.String(http.StatusOK, c.Param("item_name"))
}
