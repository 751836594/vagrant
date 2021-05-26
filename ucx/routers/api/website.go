package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ucx/pkg/models"
)

func GetStripePlan(c *gin.Context) {
	models.Setup()
	data, err := models.ExistUser(2)
	fmt.Printf("%+v", data)
	fmt.Printf("%s", err)
	c.JSON(200, gin.H{
		"message": "ping pong",
	})
}

func CanSubscribe(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "dot not anything",
	})
}
