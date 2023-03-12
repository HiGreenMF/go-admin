package common

import (
	"fmt"
	"os"
	"time"

	"github.com/go-admin/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Log *zap.SugaredLogger

func InitLogger() {
	now := time.Now()
	logFileName := fmt.Sprintf("%s/%04d-%02d-%02d-%02d.log", config.Conf.Logs.Path, now.Year(), now.Month(), now.Day(), now.Hour())

	var coreList []zapcore.Core

	encodingConfig := zapcore.EncoderConfig{
		MessageKey:    "msg",
		LevelKey:      "level",
		TimeKey:       "time",
		NameKey:       "name",
		CallerKey:     "file",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.CapitalColorLevelEncoder,
		EncodeTime: func(t time.Time, pae zapcore.PrimitiveArrayEncoder) {
			pae.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	encode := zapcore.NewConsoleEncoder(encodingConfig)

	enable := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return true
	})

	logFileWriteAsync := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logFileName,
		MaxSize:    config.Conf.Logs.MaxSize,
		MaxBackups: config.Conf.Logs.MaxBackups,
		LocalTime:  false,
		Compress:   config.Conf.Logs.Gzip,
	})

	logFileCore := zapcore.NewCore(encode, zapcore.NewMultiWriteSyncer(logFileWriteAsync, zapcore.AddSync(os.Stdout)), enable)

	coreList = append(coreList, logFileCore)
	logger := zap.New(zapcore.NewTee(coreList...), zap.AddCaller())
	Log = logger.Sugar()
	Log.Info("Init Log Done")
}
