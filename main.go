package main

import (
	"gin-blog-server/core"
	_ "gin-blog-server/docs"
	"gin-blog-server/flag"
	"gin-blog-server/global"
	"gin-blog-server/routers"
)

// @Tittle gin-blog-server Api文档
// @Version 1.0.0
// @Host 127.0.0.1:4000
// @BasePath /
func main() {
	// 加载环境变量
	core.InitCore()
	// 初始化日志
	global.Log = core.InitLogger()
	// 连接数据库
	global.DB = core.InitGorm()

	// 命令行参数绑定
	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}
	flag.RemoveFile(option)

	// 初始化路由
	router := routers.InitRouter()
	addr := global.Config.System.Addr()
	global.Log.Infof("gin-blog-server 运行在: %s", addr)

	err := router.Run(addr)
	if err != nil {
		global.Log.Fatalf(err.Error())
	}

}
