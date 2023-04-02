package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-admin/biz/handler"
)

func Ping(r *gin.Engine) {

	r.GET("/", handler.Ping)
}
