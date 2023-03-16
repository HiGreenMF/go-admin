package common

import (
	"fmt"

	"github.com/go-admin/config"
	"github.com/go-admin/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// 初始化数据库
func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&collation=%s&%s",
		config.Conf.Mysql.Username,
		config.Conf.Mysql.Password,
		config.Conf.Mysql.Host,
		config.Conf.Mysql.Port,
		config.Conf.Mysql.Database,
		config.Conf.Mysql.Charset,
		config.Conf.Mysql.Collation,
		config.Conf.Mysql.Query,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		Log.Panicf("Init Mysql Error: %v", err)
		panic(fmt.Errorf("Init Mysql Error: %v", err))
	}

	// 开启mysql日志
	if config.Conf.Mysql.LogMode {
		db.Debug()
	}

	DB = db
	// 自动迁移表结构
	dbAutoMigrate()
	Log.Info("Init DB Done")
}

// 自动迁移表结构
func dbAutoMigrate() {
	_ = DB.AutoMigrate(
		&model.User{},
	)
}
