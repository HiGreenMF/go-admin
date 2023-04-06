package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-admin/dao"
	"github.com/go-admin/util"
)

func GetUSerInfoById(ctx *gin.Context) {

	result, error := dao.GetUSerInfoById("1")
	response := util.Gin{C: ctx}

	if error != nil {
		response.ResponseNotFound(util.Error, nil, error.Error())
	}

	response.ResponseOk(util.Success,result, "ok")

}
