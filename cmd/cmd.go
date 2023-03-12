package cmd

import (
	"github.com/go-admin/common"
	"github.com/go-admin/config"
)

func BeforeStart() {
	config.InitConfig()
	common.InitLogger()
	common.InitDB()
}

func AfterStart() {

}
