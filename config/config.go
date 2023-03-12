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

type config struct {
	Logs *LogsConfig `mapstructure:"logs" json:"logs"`
}

func InitConfig() {
	workDir, err := os.Getwd()
	if err != nil {
		panic(fmt.Errorf("fatal error read app directory: %s", err))
	}
	// 配置文件名称
	viper.SetConfigFile("config.yml")
	// 添加搜素路径
	viper.AddConfigPath(workDir)
	// 读取配置信息
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error read config: %s", err))
	}

	// 热更新配置
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		if err := viper.Unmarshal(Conf); err != nil {
			panic(fmt.Errorf("fatal error hot update config: %s", err))
		}
	})

	// 读取的配置信息保存至全局Conf中
	if err := viper.Unmarshal(Conf); err != nil {
		panic(fmt.Errorf("fatal error init config: %s", err))
	}
}
