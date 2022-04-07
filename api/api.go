package api

import (
	"github.com/gin-gonic/gin"
	"time"
)

func RunAPI() {
	router := gin.Default()
	router.GET("/users/:id", getUserByID, SleepMiddleware())
	router.Run("localhost:2323")
}

// delay requests manually since we don't have a real DB
func SleepMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		time.Sleep(750 * time.Millisecond)
	}
}
