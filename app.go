package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// หน้าหลัก
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to the homepage!"})
	})

	// เส้นทางสำหรับการส่งคำขอ GET ไปยัง /hello
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})

	// เส้นทางสำหรับการรับค่าพารามิเตอร์
	router.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, gin.H{"message": "Hello " + name})
	})

	// Run เซิร์ฟเวอร์ที่พอร์ต 8080
	router.Run(":")
}
