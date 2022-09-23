package routes

import (
	"github.com/gin-gonic/gin"
	"go-test/app/common/request"
	"go-test/app/controllers/app"
	"go-test/app/controllers/common"
	"go-test/app/middleware"
	"go-test/app/services"
	"net/http"
	"time"
)

// SetApiGroupRoutes 定义 api 分组路由
func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.GET("/test", func(c *gin.Context) {
		time.Sleep(5*time.Second)
		c.String(http.StatusOK, "success")
	})

	router.POST("/user/register", func(c *gin.Context) {
		var form request.Register
		if err := c.ShouldBindJSON(&form); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": request.GetErrorMsg(form, err),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})



	router.POST("/auth/register",app.Register)
	router.POST("/auth/login", app.Login)
	router.POST("/admin/login", app.AdminLogin)

	authRouter := router.Group("").Use(middleware.JWTAuth(services.AppGuardName))
	{
		authRouter.POST("/auth/info", app.Info)
		authRouter.POST("/auth/logout", app.Logout)
	}
	router.POST("/image_upload", common.ImageUpload)

	categoryRouter := router.Group("/category")
	{
		categoryRouter.POST("/create",app.CategoryCreate)
		categoryRouter.GET("/list",app.CategoryList)
		categoryRouter.DELETE("/del",app.CategoryDel)
		categoryRouter.POST("/modify",app.CategoryModify)
	}

	articleRouter := router.Group("/article")
	{
		articleRouter.POST("/create",app.ArticleCreate)
		articleRouter.GET("/list",app.ArticleList)
		articleRouter.DELETE("/del",app.ArticleDel)
		articleRouter.POST("/modify",app.ArticleModify)
		articleRouter.GET("/detail",app.ArticleDetail)
	}

}