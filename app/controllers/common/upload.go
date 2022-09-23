package common

import (
	"github.com/gin-gonic/gin"
	"go-test/app/common/request"
	"go-test/app/common/response"
	"go-test/app/services"
)

// ImageUpload 校验入参，调用 MediaService
func ImageUpload(c *gin.Context) {
	var form request.ImageUpload
	if err := c.ShouldBind(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	outPut, err := services.MediaService.SaveImage(form)
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, outPut)
}