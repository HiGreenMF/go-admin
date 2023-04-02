package cmd

import (
	service "github.com/go-admin/biz"
	"github.com/go-admin/common"
	"github.com/go-admin/config"
)

func BeforeStart() {

	config.InitConfig()
	common.InitLogger()
	common.InitDB()
}

func AfterStart() {

	service.InitService()
}
