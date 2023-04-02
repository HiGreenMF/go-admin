package register_router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-admin/biz/router"
)

func Register(r *gin.Engine) {

	customizedRegister(r)

	router.Ping(r)

}
