package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-admin/controller"
)

func Ping(r *gin.Engine) {

	r.GET("/", controller.Ping)
}
