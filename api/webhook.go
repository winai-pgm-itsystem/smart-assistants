package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func postWebHook(r *gin.RouterGroup) {
	r.POST("/webhook", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]any{
			"webhook": "webhook is ok",
		})

	})
}

func WebHookHandler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
