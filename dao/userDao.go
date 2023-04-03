package dao

import (
	"fmt"

	"github.com/go-admin/common"
	"github.com/go-admin/model"
)

func GetUSerInfoById(userId string) (*model.User, error) {
	var user *model.User
	result := common.DB.Model(&model.User{}).Where("id = ?", userId).First(&user)

	if result.Error != nil {
		return nil, fmt.Errorf("failed to get user info: %v", result.Error)
	}

	return user, nil
}
