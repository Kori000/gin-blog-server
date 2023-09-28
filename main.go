package main

import (
	"gin-blog-server/core"
	"gin-blog-server/flag"
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

	// 命令行参数绑定
	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}

	// 初始化路由
	router := routers.InitRouter()
	addr := global.Config.System.Addr()
	global.Log.Infof("gin-blog-server 运行在: %s", addr)

	err := router.Run(addr)
	if err != nil {
		global.Log.Fatalf(err.Error())
	}

}
