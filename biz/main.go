package biz

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-admin/config"
	"github.com/go-admin/register_router"
)

func InitService() {

	mode := os.Getenv("GIN_MODE")
	if mode == "" {
		mode = gin.ReleaseMode
	}
	gin.SetMode(mode)

	r := gin.Default()
	register_router.Register(r)

	r.Run(config.Conf.Service.Port)
}
