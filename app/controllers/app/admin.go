package app

import (
	"github.com/gin-gonic/gin"
	"go-test/app/common/request"
	"go-test/app/common/response"
	"go-test/app/services"
	"log"
)

// AdminLogin 进行入参校验，并调用 UserService 和 JwtService 服务，颁发 Token
func AdminLogin(c *gin.Context) {
	var form request.AdminLogin
	if err := c.ShouldBindJSON(&form); err != nil {
		log.Printf("1 %v",err)
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, admin := services.AdminService.Login(form); err != nil {
		log.Printf("2 %v",err)
		response.BusinessFail(c, err.Error())
	} else {
		tokenData, err, _ := services.JwtService.CreateToken(services.AppGuardName, admin)
		if err != nil {
			log.Printf("3 %v",err)
			response.BusinessFail(c, err.Error())
			return
		}
		response.Success(c, tokenData)
	}
}