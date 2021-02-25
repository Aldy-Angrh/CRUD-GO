package service

import (
	"os"

	"github.com/gin-gonic/gin"
)

//CORSMiddleware function to allow access control
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		clientUrl := os.Getenv("*")
		c.Header("Access-Control-Allow-Origin", clientUrl)
		c.Header("Access-Control-Allow-Credentials", "false")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
