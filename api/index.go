package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	app *gin.Engine
)

func myRoute(r *gin.RouterGroup) {
	r.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]any{
			"userId":   "U001",
			"userName": "john doe",
		})
	})
}

func helthCheck(r *gin.RouterGroup) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]any{
			"app":     "smart-assistants",
			"version": "v0.0.1",
		})
	})
}
func init() {
	app = gin.New()
	r := app.Group("/api")
	myRoute(r)
	helthCheck(r)

}

func Handler(w http.ResponseWriter, r *http.Request) {

	app.ServeHTTP(w, r)
}
