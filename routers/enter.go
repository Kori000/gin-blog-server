package routers

import (
	"gin-blog-server/global"

	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	// 设置路由模式
	gin.SetMode(global.Config.System.Env)
	// 初始化路由
	router := gin.Default()

	apiRouterGourp := router.Group("api")

	routerGroupApp := RouterGroup{apiRouterGourp}
	// 系统配置api
	routerGroupApp.SettingsRouter()

	return router
}
