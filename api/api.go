package api

import (
	"github.com/gin-gonic/gin"
	"time"
)

func RunAPI() {
	router := gin.Default()
	router.Use(CORSMiddleware())
	router.GET("/users/:id", getUserByID, SleepMiddleware())
	router.Run("localhost:2323")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// delay requests manually since we don't have a real DB
func SleepMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		time.Sleep(750 * time.Millisecond)
	}
}
