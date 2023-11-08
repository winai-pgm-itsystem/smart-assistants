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

		c.JSON(http.StatusOK, map[string]any{
			"app":     "smart-assistants",
			"version": "v0.0.1",
		})
	})
}

func init() {
	app = gin.New()
	router := app.Group("/api")
	helthCheck(router)

	//secret
	getSecret(router)

	//user
	getUser(router)
	postUser(router)

}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}

func middelwaresVerifly(c *gin.Context) {
	veriflyId := c.Query("veriflyId")
	veriflyEnvId := os.Getenv("VERIFLY_ID")

	if veriflyId == veriflyEnvId {
		c.Next()
	} else {
		c.JSON(http.StatusOK, map[string]any{
			"message": "verifly id is invalid.",
		})

	}

}
