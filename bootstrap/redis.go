package bootstrap

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go-test/conf"
	"go.uber.org/zap"
)

func InitializeRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     conf.Global.ConfigStruct.Redis.Host + ":" + conf.Global.ConfigStruct.Redis.Port,
		Password: conf.Global.ConfigStruct.Redis.Password, // no password set
		DB:       conf.Global.ConfigStruct.Redis.DB,       // use default DB
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		conf.Global.Log.Error("Redis connect ping failed, err:", zap.Any("err", err))
		return nil
	}
	return client
}