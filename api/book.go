package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Book(c echo.Context) error {

	return c.String(http.StatusOK, "Hello, book!")

}
