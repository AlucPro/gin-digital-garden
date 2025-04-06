package initialize

import (
	"gin-digital-garden/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	systemRouter := router.RouterGroupApp.System
	blogRouter := router.RouterGroupApp.Blog

	PublicGroup := Router.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}

	{
		systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
		blogRouter.InitBaseRouter(PublicGroup)   // 注册博客功能路由 不做鉴权
	}

	return Router

}
