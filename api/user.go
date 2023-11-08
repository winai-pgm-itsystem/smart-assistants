package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getUser(r *gin.RouterGroup) {
	r.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]any{
			"userId":   "U001",
			"userName": "john doe",
		})
	})
}

func postUser(r *gin.RouterGroup) {
	r.POST("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]any{
			"userId":  "U001",
			"message": "updated successfuly!.",
		})

	})
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
