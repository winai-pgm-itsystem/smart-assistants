package main

import (
	"smart-assistants/api"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.New()
	handler.SetupRoutes(app)

	app.Run(":8080")
}
