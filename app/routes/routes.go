package routes

import (
	"smart-assistants/app/handler"

	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(app *gin.Engine) {
	app.NoRoute(ErrRouter)

	app.GET("/ping", handler.Ping)

	route := app.Group("/api")
	{
		route.GET("/hello/:name", handler.Hello)

	}
}

func ErrRouter(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"errors": "this page could not be found",
	})
}
