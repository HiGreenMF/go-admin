package cmd

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-admin/common"
	"github.com/go-admin/config"
	"github.com/go-admin/router"
)

func initService() {

	mode := os.Getenv("GIN_MODE")
	if mode == "" {
		mode = gin.ReleaseMode
	}
	gin.SetMode(mode)

	r := gin.Default()
	router.Register(r)

	r.Run(config.Conf.Service.Port)
}

func BeforeStart() {

	config.InitConfig()
	common.InitLogger()
	common.InitDB()
}

func AfterStart() {

	initService()
}
