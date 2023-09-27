package main

import (
	"gin-blog-server/core"
	"gin-blog-server/global"
	"gin-blog-server/routers"
)

func main() {
	// 加载环境变量
	core.InitCore()
	// 初始化日志
	global.Log = core.InitLogger()
	// 连接数据库
	global.DB = core.InitGorm()

	// 初始化路由
	router := routers.InitRouter()
	addr := global.Config.System.Addr()
	global.Log.Infof("gin-blog-server运行在: %s", addr)
	router.Run(addr)

}
