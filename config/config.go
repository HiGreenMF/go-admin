package config

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
)

var Conf = new(config)

type LogsConfig struct {
	Level      zapcore.Level `mapstructure:"level" json:"level"`
	Path       string        `mapstructure:"path" json:"path"`
	MaxSize    int           `mapstructure:"max-size" json:"maxSize"`
	MaxBackups int           `mapstructure:"max-backups" json:"maxBackups"`
	MaxAge     int           `mapstructure:"max-age" json:"maxAge"`
	Gzip       bool          `mapstructure:"gzip" json:"gzip"`
}

type MysqlConfig struct {
	Username    string `mapstructure:"username" json:"username"`
	Password    string `mapstructure:"password" json:"password"`
	Database    string `mapstructure:"database" json:"database"`
	Host        string `mapstructure:"host" json:"host"`
	Port        int    `mapstructure:"port" json:"port"`
	Query       string `mapstructure:"query" json:"query"`
	LogMode     bool   `mapstructure:"log-mode" json:"logMode"`
	TablePrefix string `mapstructure:"table-prefix" json:"tablePrefix"`
	Charset     string `mapstructure:"charset" json:"charset"`
	Collation   string `mapstructure:"collation" json:"collation"`
}

type ServiceConfig struct {
	Port string `mapstructure:"port" json:"port"`
}
type config struct {
	Logs    *LogsConfig    `mapstructure:"logs" json:"logs"`
	Mysql   *MysqlConfig   `mapstructure:"mysql" json:"mysql"`
	Service *ServiceConfig `mapstructure:"service" json:"service"`
}

func InitConfig() {
	workDir, err := os.Getwd()
	if err != nil {
		panic(fmt.Errorf("fatal error read app directory: %s", err).(any))
	}
	// 配置文件名称
	viper.SetConfigFile("config.yml")
	// 添加搜素路径
	viper.AddConfigPath(workDir)
	// 读取配置信息
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error read config: %s", err).(any))
	}

	// 热更新配置
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		if err := viper.Unmarshal(Conf); err != nil {
			panic(fmt.Errorf("fatal error hot update config: %s", err).(any))
		}
	})

	// 读取的配置信息保存至全局Conf中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Print(Conf)
		panic(fmt.Errorf("fatal error init config: %s", err).(any))
	}
}
