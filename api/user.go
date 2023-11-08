package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func User(r *gin.RouterGroup) {
	r.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]any{
			"userId":   "U001",
			"userName": "john doe",
		})
	})
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
