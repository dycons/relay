package main

import (
	"github.com/dycons/relay/app"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/consents", app.ConsentsGet)
	r.Run(":3000")
}
