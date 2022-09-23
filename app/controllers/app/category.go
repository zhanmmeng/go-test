package app

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"go-test/app/common/request"
	"go-test/app/common/response"
	"go-test/app/services"
	_ "go-test/utils"
)

func CategoryCreate(c *gin.Context) {

	var form request.CategoryCreate
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, user := services.CategoryService.CategoryCreate(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, user)
	}
}

func CategoryList(c *gin.Context) {
	if err, category := services.CategoryService.CategoryList();err != nil{
		response.BusinessFail(c,err.Error())
	}else {
		response.Success(c,category)
	}
}

func CategoryDel(c *gin.Context) {
	var form request.CategoryDel
	if err := c.ShouldBindQuery(&form); err != nil{
		response.ValidateFail(c,request.GetErrorMsg(form,err))
		return
	}

	if err := services.CategoryService.CategoryDel(form); err!=nil {
		response.BusinessFail(c,err.Error())
	}else {
		response.Success(c,"删除成功")
	}
}

func CategoryModify(c *gin.Context) {
	var form request.CategoryModify
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c,request.GetErrorMsg(form,err))
		return
	}

	if err, categoryModify := services.CategoryService.CategoryModify(form); err != nil {
		response.BusinessFail(c,err.Error())
	}else {
		response.Success(c,categoryModify)
	}
}