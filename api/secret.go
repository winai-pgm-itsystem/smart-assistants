package handler

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func getSecret(r *gin.RouterGroup) {
	r.GET("/secret",
		func(c *gin.Context) {

			veriflyId := c.Query("veriflyId")
			veriflyEnvId := os.Getenv("VERIFLY_ID")
			apiKey := os.Getenv("API_KEY")

			if veriflyId == veriflyEnvId {
				c.JSON(http.StatusOK, map[string]any{
					"apiKey": apiKey,
				})
			}

			c.JSON(http.StatusOK, map[string]any{
				"message": "varifly id is invalid.",
			})
		})
}

func SecretHandler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
