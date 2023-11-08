package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func myRoute(r *gin.RouterGroup) {
	r.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]any{
			"userId":   "U001",
			"userName": "john doe",
		})
	})
}
