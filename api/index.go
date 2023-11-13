package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	app *gin.Engine
)

func HealthCheck(r *gin.RouterGroup) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"app":     "smart-assistants",
			"version": "v0.0.1",
		})

	})
}

// comment this when deploy to vercel
// func SetupRoutes(app *gin.Engine) {
// 	router := app.Group("/api")

// 	HealthCheck(router)
// 	CheckEnvirontmentTarget(router)
// 	LineWebHookHandler(router)

// }

func init() {
	app = gin.New()
	router := app.Group("/api")
	HealthCheck(router)
	CheckEnvirontmentTarget(router)
	LineWebHookHandler(router)

}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
