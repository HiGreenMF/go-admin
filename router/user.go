package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-admin/controller"
	"github.com/go-admin/middleware"
)

func User(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	v1.Use(middleware.AuthMiddleware())
	{
		v1.GET("/get-user-info", controller.GetUSerInfoById)
	}

}
