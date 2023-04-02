package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID            uint   `gorm:"primaryKey"`
	UserName      string `gorm:"type:varchar(50);not null;unique;comment:'用户名'" json:"userName"`     // 用户名
	NickName      string `gorm:"type:varchar(50);comment:'中文名'" json:"nickName"`                     // 昵称
	Name          string `gorm:"type:varchar(50);comment:'中文名'" json:"name"`                         // 中文名
	Mail          string `gorm:"type:varchar(100);comment:'邮箱'" json:"mail"`                         // 邮箱
	JobNumber     string `gorm:"type:varchar(20);comment:'工号'" json:"jobNumber"`                     // 工号
	MobileNumber  string `gorm:"type:varchar(15);not null;unique;comment:'手机号'" json:"mobileNumber"` // 手机号
	Avatar        string `gorm:"type:varchar(255);comment:'头像'" json:"avatar"`                       // 头像
	PostalAddress string `gorm:"type:varchar(255);comment:'地址'" json:"postalAddress"`                // 地址
	Department    string `gorm:"type:varchar(128);comment:'部门'" json:"department"`                   // 部门
	Position      string `gorm:"type:varchar(128);comment:'职位'" json:"position"`                     //  职位
	Status        uint   `gorm:"type:tinyint(1);default:1;comment:'状态:1在职, 2离职'" json:"status"`      // 状态
	DepartmentId  string `gorm:"type:varchar(100);not null;comment:'部门id'" json:"departmentId"`      // 部门id
	OpenId        string `gorm:"type:varchar(100);not null;comment:'用户open_id'" json:"openId"`       // 用户open_id
}
