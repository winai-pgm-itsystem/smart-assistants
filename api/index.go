// package handler

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// var (
// 	app *gin.Engine
// )

// func HelthCheck(r *gin.RouterGroup) {
// 	r.GET("/", func(c *gin.Context) {

// 		c.JSON(http.StatusOK, map[string]any{
// 			"app":     "smart-assistants",
// 			"version": "v0.0.1",
// 		})
// 	})
// }

// func main() {
// 	app = gin.New()
// 	router := app.Group("/api")
// 	HelthCheck(router)

// 	//secret
// 	getSecret(router)

// 	//user
// 	postWebHook(router)

// }

// func Handler(w http.ResponseWriter, r *http.Request) {
// 	app.ServeHTTP(w, r)
// }

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	StatusBadRequest = http.StatusBadRequest
	StatusOK         = http.StatusOK
)

func healthCheck(c *gin.Context) {
	c.JSON(StatusOK, gin.H{
		"message": "server running..",
	})
}

func main() {
	router := gin.Default()
	router.GET("/", healthCheck)

	router.Run()

}
