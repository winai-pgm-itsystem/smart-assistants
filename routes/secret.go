package handler

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func CheckEnvirontmentTarget(r *gin.RouterGroup) {
	r.GET("/secret", func(c *gin.Context) {
		veriflyId := c.Query("veriflyId")
		veriflyIdFromEnv := os.Getenv("VERIFLY_ID")
		if veriflyId == veriflyIdFromEnv {
			c.JSON(http.StatusOK, map[string]interface{}{
				"veriflyId":              veriflyIdFromEnv,
				"apiKey":                 os.Getenv("API_KEY"),
				"lineChannelAccessToken": os.Getenv("LINE_CHANNEL_ACCESS_TOKEN"),
			})

		} else {
			c.JSON(http.StatusOK, map[string]interface{}{
				"message": "verifly id is invalid.",
			})
			return
		}
	})
}

// func SecretHandler(w http.ResponseWriter, r *http.Request) {
// 	app.ServeHTTP(w, r)
// }
