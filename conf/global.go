package conf

import (
	"github.com/go-redis/redis/v8"
	"github.com/jassue/go-storage/storage"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type GlobalConfig struct {
	Viper        *viper.Viper
	ConfigStruct ConfigStruct
	Log          *zap.Logger
	DB           *gorm.DB
	Redis        *redis.Client
}

var Global = new(GlobalConfig)

// Disk 获取文件系统实例的统一入口
func (app *GlobalConfig) Disk(disk... string) storage.Storage {
	// 若未传参，默认使用配置文件驱动
	diskName := app.ConfigStruct.Storage.Default
	if len(disk) > 0 {
		diskName = storage.DiskName(disk[0])
	}
	s, err := storage.Disk(diskName)
	if err != nil {
		panic(err)
	}
	return s
}