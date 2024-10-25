package service

import "github.com/gin-gonic/gin"

func GetAllUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "List of all products",
	})
}
