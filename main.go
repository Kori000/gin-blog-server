package main

import (
	"gin-blog-server/core"
	"gin-blog-server/global"
)

func main() {
	// 加载环境变量
	core.InitCore()
	//初始化日志
	global.Log = core.InitLogger()
	// 连接数据库
	global.DB = core.InitGorm()

	global.Log.Warnln("哈哈哈")
	global.Log.Errorln("哈哈哈1")
	global.Log.Infoln("哈哈哈2")

}
