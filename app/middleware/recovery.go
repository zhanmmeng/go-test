package middleware

import (
	"github.com/gin-gonic/gin"
	"go-test/app/common/response"
	"go-test/conf"
	"gopkg.in/natefinch/lumberjack.v2"
)

func CustomRecovery() gin.HandlerFunc {
	return gin.RecoveryWithWriter(
		&lumberjack.Logger{
			Filename:   conf.Global.ConfigStruct.Log.RootDir + "/" + conf.Global.ConfigStruct.Log.Filename,
			MaxSize:    conf.Global.ConfigStruct.Log.MaxSize,
			MaxBackups: conf.Global.ConfigStruct.Log.MaxBackups,
			MaxAge:     conf.Global.ConfigStruct.Log.MaxAge,
			Compress:   conf.Global.ConfigStruct.Log.Compress,
		},
		response.ServerError)
}