package bootstrap

import (
	"go-test/conf"
	utils "go-test/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

var (
	level zapcore.Level // zap 日志等级
	options []zap.Option // zap 配置项
)

func InitializeLog() *zap.Logger {
	// 创建根目录
	createRootDir()

	// 设置日志等级
	setLogLevel()

	if conf.Global.ConfigStruct.Log.ShowLine {
		options = append(options, zap.AddCaller())
	}

	// 初始化 zap
	return zap.New(getZapCore(), options...)
}

func createRootDir() {
	if ok, _ := utils.PathExists(conf.Global.ConfigStruct.Log.RootDir); !ok {
		_ = os.Mkdir(conf.Global.ConfigStruct.Log.RootDir, os.ModePerm)
	}
}

func setLogLevel() {
	switch conf.Global.ConfigStruct.Log.Level {
	case "debug":
		level = zap.DebugLevel
		options = append(options, zap.AddStacktrace(level))
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
		options = append(options, zap.AddStacktrace(level))
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}
}

// 扩展 Zap
func getZapCore() zapcore.Core {
	var encoder zapcore.Encoder

	// 调整编码器默认配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(time.Format("[" + "2006-01-02 15:04:05.000" + "]"))
	}
	encoderConfig.EncodeLevel = func(l zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(conf.Global.ConfigStruct.App.Env + "." + l.String())
	}

	// 设置编码器
	if conf.Global.ConfigStruct.Log.Format == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	return zapcore.NewCore(encoder, getLogWriter(), level)
}

// 使用 lumberjack 作为日志写入器
func getLogWriter() zapcore.WriteSyncer {
	file := &lumberjack.Logger{
		Filename:   conf.Global.ConfigStruct.Log.RootDir + "/" + conf.Global.ConfigStruct.Log.Filename,
		MaxSize:    conf.Global.ConfigStruct.Log.MaxSize,
		MaxBackups: conf.Global.ConfigStruct.Log.MaxBackups,
		MaxAge:     conf.Global.ConfigStruct.Log.MaxAge,
		Compress:   conf.Global.ConfigStruct.Log.Compress,
	}

	return zapcore.AddSync(file)
}
