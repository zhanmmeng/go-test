package app

import (
	"github.com/gin-gonic/gin"
	"go-test/app/common/request"
	"go-test/app/common/response"
	"go-test/app/services"
)

func ArticleCreate(c *gin.Context) {
	var form request.ArticleCreate
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, articleId := services.ArticleService.ArticleCreate(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, articleId)
	}
}

func ArticleDel(c *gin.Context) {
	var form request.ArticleDel
	if err := c.ShouldBindQuery(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
	}

	if err := services.ArticleService.ArticleDel(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, "删除成功")

	}

}

func ArticleModify(c *gin.Context) {
	var form request.ArticleModify
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	if err := services.ArticleService.ArticleModify(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, "修改成功")
	}
}

func ArticleList(c *gin.Context) {
	var form request.ArticleList

	c.ShouldBindQuery(&form)

	if err, list, page := services.ArticleService.ArticleList(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		c.AbortWithStatusJSON(200,
			gin.H{
				"data": list, "page": page,
			})
		//response.Success(c,list)
	}

}

func ArticleDetail(c *gin.Context) {
	var form request.ArticleDel
	c.ShouldBindQuery(&form)

	if err, article := services.ArticleService.ArticleDetail(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, article)
	}
}
