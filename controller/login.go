package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-admin/common"
	"github.com/go-admin/constants"
	"github.com/go-admin/util"
	"github.com/larksuite/oapi-sdk-go/v3"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	"github.com/larksuite/oapi-sdk-go/v3/service/auth/v3"
)

type Response struct {
	AppAccessToken    string `json:"app_access_token"`
	Code              int    `json:"code"`
	Expire            int    `json:"expire"`
	Message           string `json:"msg"`
	TenantAccessToken string `json:"tenant_access_token"`
}

type AppAccessTokenReqBody struct {
	AppId     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
}

func Login(ctx *gin.Context) {

	// 创建 Client

	// url := "https://open.feishu.cn/open-apis/auth/v3/app_access_token/internal"
	// method := "POST"
	// reqBody := &AppAccessTokenReqBody{
	// 	AppId:     constants.AppID,
	// 	AppSecret: constants.AppSecret,
	// }

	// client := &http.Client{}
	// reqBodyBytes, _ := json.Marshal(reqBody)
	// payload := bytes.NewReader(reqBodyBytes)

	// req, err := http.NewRequest(method, url, payload)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// req.Header.Add("Content-Type", "application/json")
	// res, err := client.Do(req)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer res.Body.Close()

	// body, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// responseJSON := string(body)
	// var resp Response
	// err1 := json.Unmarshal([]byte(responseJSON), &resp)
	// if err1 != nil {
	// 	panic(err)
	// }

	// tokenStr, _ := util.CreateToken(resp.AppAccessToken)
	// ctx.Header("Authorization", fmt.Sprintf("Bearer %v", tokenStr))



	client := lark.NewClient("appID", "appSecret")
	// 创建请求对象、

	req := larkauth.NewInternalAppAccessTokenReqBuilder().
		Body(larkauth.NewInternalAppAccessTokenReqBodyBuilder().
			AppId(constants.AppID).
			AppSecret(constants.AppSecret).
			Build()).
		Build()
	// 发起请求
	resp, err := client.Auth.AppAccessToken.Internal(context.Background(), req,larkcore.WithUserAccessToken("123"))
	response := util.Gin{C: ctx}
	// 处理错误
	if err != nil {
		response.ResponseBadRequest(util.Error, nil, err.Error())
		common.Log.Error(err)
		return
	}

	// 服务端错误处理
	if !resp.Success() {
		response.ResponseBadRequest(resp.Code, resp.RequestId(),resp.Msg)
		common.Log.Error(resp.Code, resp.Msg, resp.RequestId())
		return
	}

	// 业务处理
	response.ResponseOk(resp.Code, resp.ApiResp, resp.Msg)

}
