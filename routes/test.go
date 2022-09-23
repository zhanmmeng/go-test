package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetTestGroupRoutes(router *gin.RouterGroup) {
	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK,"login.html",gin.H{
			"title":"html/login",
		})
	})
}
