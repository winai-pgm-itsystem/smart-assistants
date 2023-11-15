package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var srv http.Handler

func main() {

	e := echo.New()
	e.GET("/", HealthCheck)
	e.GET("/books", Book)
	srv = e

	e.Start(":8080")
}

func Hanlder(w http.ResponseWriter, r *http.Request) {
	srv.ServeHTTP(w, r)
}
