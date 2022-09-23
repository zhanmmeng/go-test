package bootstrap

import (
	"github.com/jassue/go-storage/kodo"
	"github.com/jassue/go-storage/local"
	"github.com/jassue/go-storage/oss"
	"go-test/conf"
)

func InitializeStorage() {
	_, _ = local.Init(conf.Global.ConfigStruct.Storage.Disks.Local)
	_, _ = kodo.Init(conf.Global.ConfigStruct.Storage.Disks.QiNiu)
	_, _ = oss.Init(conf.Global.ConfigStruct.Storage.Disks.AliOss)
}