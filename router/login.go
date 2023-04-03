package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-admin/controller"
)

func Login(r *gin.Engine) {
	r.POST("/login", controller.Login)

}
