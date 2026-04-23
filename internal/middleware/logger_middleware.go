package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func CustomLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next() // proses request (seperti next() di express / nestjs)

		// setelah request selesai
		latency := time.Since(start)
		status := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path

		log.Printf("[CUSTOM LOGGER] - [%d] %s %s (%v)", status, method, path, latency)
	}
}
