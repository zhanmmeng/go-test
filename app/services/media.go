package services

import (
	"context"
	"errors"
	"github.com/jassue/go-storage/storage"
	"github.com/satori/go.uuid"
	"go-test/app/common/request"
	"go-test/app/models"
	"go-test/conf"
	"path"
	"strconv"
	"time"
)

//图片上传

type mediaService struct {
}

var MediaService = new(mediaService)

type outPut struct {
	Id int64 `json:"id"`
	Path string `json:"path"`
	Url string `json:"url"`
}

const mediaCacheKeyPre = "media:"

// 文件存储目录
func (mediaService *mediaService) makeFaceDir(business string) string {
	return conf.Global.ConfigStruct.App.Env + "/" + business
}

// HashName 生成文件名称（使用 uuid）
func (mediaService *mediaService) HashName(fileName string) string {
	fileSuffix := path.Ext(fileName)
	return uuid.NewV4().String() + fileSuffix
}

// SaveImage 保存图片（公共读）
func (mediaService *mediaService) SaveImage(params request.ImageUpload) (result outPut, err error) {
	file, err := params.Image.Open()
	defer file.Close()
	if err != nil {
		err = errors.New("上传失败")
		return
	}

	localPrefix := ""
	// 本地文件存放路径为 storage/app/public，由于在『（五）静态资源处理 & 优雅重启服务器』中，
	// 配置了静态资源处理路由 router.Static("/storage", "./storage/app/public")
	// 所以此处不需要将 public/ 存入到 mysql 中，防止后续拼接文件 Url 错误
	if conf.Global.ConfigStruct.Storage.Default == storage.Local {
		localPrefix = "public" + "/"
	}
	key := mediaService.makeFaceDir(params.Business) + "/" + mediaService.HashName(params.Image.Filename)
	disk := conf.Global.Disk()
	err = disk.Put(localPrefix + key, file, params.Image.Size)
	if err != nil {
		return
	}

	image := models.Media{
		DiskType: string(conf.Global.ConfigStruct.Storage.Default),
		SrcType:    1,
		Src:        key,
	}
	err = conf.Global.DB.Create(&image).Error
	if err != nil {
		return
	}

	result = outPut{int64(image.ID.ID), key, disk.Url(key)}
	return
}

// GetUrlById 通过 id 获取文件 url
func (mediaService *mediaService) GetUrlById(id int64) string {
	if id == 0 {
		return ""
	}

	var url string
	cacheKey := mediaCacheKeyPre + strconv.FormatInt(id,10)

	exist := conf.Global.Redis.Exists(context.Background(), cacheKey).Val()
	if exist == 1 {
		url = conf.Global.Redis.Get(context.Background(), cacheKey).Val()
	} else {
		media := models.Media{}
		err := conf.Global.DB.First(&media, id).Error
		if err != nil {
			return ""
		}
		url = conf.Global.Disk(media.DiskType).Url(media.Src)
		conf.Global.Redis.Set(context.Background(), cacheKey, url, time.Second*3*24*3600)
	}

	return url
}