package controller

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-admin/common"
	"github.com/go-admin/constants"
	"github.com/go-admin/util"
	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	larkauth "github.com/larksuite/oapi-sdk-go/v3/service/auth/v3"
)

type TokenInfo struct {
	AppAccessToken    string `json:"app_access_token"`
	Code              int    `json:"code"`
	Expire            int    `json:"expire"`
	Msg               string `json:"msg"`
	TenantAccessToken string `json:"-"`
}

func Login(ctx *gin.Context) {

	// 创建 Client
	client := lark.NewClient(constants.AppID, constants.AppSecret)

	// 创建请求对象、
	req := larkauth.NewInternalAppAccessTokenReqBuilder().
		Body(larkauth.NewInternalAppAccessTokenReqBodyBuilder().
			AppId(constants.AppID).
			AppSecret(constants.AppSecret).
			Build()).
		Build()
	// 发起请求
	resp, err := client.Auth.AppAccessToken.Internal(context.Background(), req)
	response := util.Gin{C: ctx}

	// 处理错误
	if err != nil {
		response.ResponseBadRequest(util.Error, nil, err.Error())
		common.Log.Error(err)
		return
	}

	// 服务端错误处理
	if !resp.Success() {
		response.ResponseBadRequest(resp.Code, resp.RequestId(), resp.Msg)
		common.Log.Error(resp.Code, resp.Msg, resp.RequestId())
		return
	}
	// 读取Body
	var tokenInfo TokenInfo
	// 	err = json.Unmarshal(resp.ApiResp.RawBody, &tokenInfo)

	err = resp.ApiResp.JSONUnmarshalBody(&tokenInfo, &larkcore.Config{
		Serializable: new(larkcore.DefaultSerialization),
	})
	if err != nil {
		common.Log.Error(err)
		response.ResponseBadRequest(util.Error, nil, err.Error())
		return
	}
	tokenStr, _ := util.CreateToken(tokenInfo.AppAccessToken)
	// 生成jwtToken
	ctx.Header("Authorization", fmt.Sprintf("Bearer %v", tokenStr))
	response.ResponseOk(tokenInfo.Code, tokenInfo, tokenInfo.Msg)

}
