package cmd

import (
	"github.com/go-admin/common"
	"github.com/go-admin/config"
	"github.com/go-admin/model"
)

func BeforeStart() {
	config.InitConfig()
	common.InitLogger()
	common.InitDB()
}

func AfterStart() {
	err := common.DB.Create(&model.User{
		UserName:      "Nicholas Zhao Si",
		NickName:      "亚洲舞王",
		Name:          "尼古拉斯·赵四",
		Mail:          "xxx@163.com",
		JobNumber:     "5481230",
		MobileNumber:  "13129000000",
		Avatar:        "https://bkimg.cdn.bcebos.com/pic/8ad4b31c8701a18b87d62e146c65100828381f30501f?x-bce-process=image/watermark,image_d2F0ZXIvYmFpa2U5Mg==,g_7,xp_5,yp_5",
		PostalAddress: "辽宁省开原市松山镇象牙山村001号",
		Department:    "董事长",
		Position:      "舞王",
		Status:        1,
		DepartmentId:  "100001",
		OpenId:        "k95wxiangyashan1",
	}).Error

	if err != nil {
		common.Log.Errorf("Create User Error: %s", err)
	}
}
