package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-admin/dao"
)

func GetUSerInfoById(ctx *gin.Context) {

	result, error := dao.GetUSerInfoById("1")
	if error != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": error,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "OK",
		"data":    result,
	})

}
