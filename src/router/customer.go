package router

import (
	"example/src/service"

	"github.com/gin-gonic/gin"
)

func CustomerRoute(r *gin.Engine) {
	customerGroup := r.Group("/customer")
	{
		customerGroup.POST("/export", service.ExportExcel)
	}
}
