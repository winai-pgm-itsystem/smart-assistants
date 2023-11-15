package handler

import (
	"encoding/json"
	"net/http"
	"os"
	"smart-assistants/entities"
	"smart-assistants/handlers"

	"github.com/gin-gonic/gin"
)

func LineWebHookHandler(r *gin.RouterGroup) {
	r.POST("/linewebhook",
		func(c *gin.Context) {
			apiKey := c.Query("apikey")
			if !isValidAPIKey(apiKey) {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "API key is invalid.",
				})
				return
			}
			webHookEvent := &entities.WebHookEvent{}
			if err := c.BindJSON(webHookEvent); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "body error",
				})
				return
			}

			_, err := json.Marshal(webHookEvent)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "error converting to JSON",
				})
				return
			}

			if len(webHookEvent.Events) != 0 {
				handlers.HandleMessageEvent(webHookEvent)

			}
			c.JSON(http.StatusOK, gin.H{
				"message": "successfuly.",
			})

		},
	)
}

func isValidAPIKey(apiKey string) bool {
	apiKeyFromEnv := os.Getenv("API_KEY")
	return apiKey == apiKeyFromEnv
}

// func LineHandler(w http.ResponseWriter, r *http.Request) {
// 	app.ServeHTTP(w, r)
// }
