package api

import "github.com/gin-gonic/gin"

func GetStripePlan(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ping pong",
	})
}

func CanSubscribe(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "dot not anything",
	})
}
