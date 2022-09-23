package services

import (
	"errors"
	"go-test/app/common/request"
	"go-test/app/models"
	"go-test/conf"
)

type adminService struct {
}

var AdminService = new(adminService)

func (admin *adminService) Login(params request.AdminLogin) (err error, adminUser *models.Admin)  {
	err = conf.Global.DB.Table("admin_user").Where("account = ?", params.Account).First(&adminUser).Error
	if err != nil || params.Password != adminUser.Password {
		err = errors.New("管理员不存在或密码错误")
	}
	return
}
