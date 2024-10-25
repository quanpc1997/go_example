package router

import (
	"example/src/service"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) {
	userGroup := r.Group("/user")
	{
		userGroup.GET("/", service.GetAllUsers)
	}
}
