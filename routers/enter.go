package routers

import (
	"github.com/gin-gonic/gin"

	"gin-blog-server/global"
	"gin-blog-server/middleware"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	// 设置路由模式
	gin.SetMode(global.Config.System.Env)
	// 初始化路由
	router := gin.Default()

	router.Use(middleware.Cors())

	apiRouterGourp := router.Group("api")

	routerGroupApp := RouterGroup{apiRouterGourp}

	// 系统配置api
	routerGroupApp.SettingsRouter()
	// 图片管理
	routerGroupApp.ImagesRouter()
	// 广告管理
	routerGroupApp.AdvertisementRouter()

	return router
}
