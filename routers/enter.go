package routers

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

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

	// swagger router
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Use(middleware.Cors())

	apiRouterGourp := router.Group("api")

	routerGroupApp := RouterGroup{apiRouterGourp}

	// 系统配置
	routerGroupApp.SettingsRouter()
	// 功能测试
	routerGroupApp.TestRouter()
	// Token管理
	routerGroupApp.TokenRouter()
	// 图片管理
	routerGroupApp.ImagesRouter()
	// 广告管理
	routerGroupApp.AdvertisementRouter()
	// 菜单管理
	routerGroupApp.MenuRouter()

	return router
}
