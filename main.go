package main

import (
	"go-test/bootstrap"
	"go-test/conf"
	_ "go-test/conf"
)


func main() {
	conf.Init()

	// 初始化日志
	conf.Global.Log = bootstrap.InitializeLog()

	// 初始化数据库
	conf.Global.DB = bootstrap.InitializeDB()
	// 程序关闭前，释放数据库连接
	defer func() {
		if conf.Global.DB != nil {
			db, _ := conf.Global.DB.DB()
			db.Close()
		}
	}()

	// 初始化验证器
	bootstrap.InitializeValidator()

	// 初始化Redis
	conf.Global.Redis = bootstrap.InitializeRedis()

	// 初始化文件系统
	bootstrap.InitializeStorage()

	bootstrap.RunServer()
}