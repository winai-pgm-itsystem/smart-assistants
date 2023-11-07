package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	app *gin.Engine
)

func heathCheck(r *gin.RouterGroup) {
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello from golang in vercel")
	})
}

func testRoute(r *gin.RouterGroup) {
	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "test  golang in vercel")
	})
}

func init() {
	app = gin.New()
	router := app.Group("/api")

	heathCheck(router)
	testRoute(router)

}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
