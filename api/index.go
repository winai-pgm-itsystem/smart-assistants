package handler

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	app *gin.Engine
)

func helthCheck(r *gin.RouterGroup) {
	r.GET("/", func(c *gin.Context) {
		apiKey := os.Getenv("API_KEY")
		c.JSON(http.StatusOK, map[string]any{
			"app":     "smart-assistants",
			"version": "v0.0.1",
			"apiKey":  apiKey,
		})
	})
}

func init() {
	app = gin.New()
	r := app.Group("/api")
	helthCheck(r)
	getUser(r)
	postUser(r)

}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
