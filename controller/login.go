package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-admin/util"
)

func Login(ctx *gin.Context) {

	tokenStr, _ := util.CreateToken("20230402")
	ctx.Header("Authorization", fmt.Sprintf("Bearer %v", tokenStr))

	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    tokenStr,
	})

}
