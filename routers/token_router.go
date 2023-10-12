package routers

import (
	"gin-blog-server/api"
)

func (router RouterGroup) TokenRouter() {
	app := api.ApiGroupApp.QiNiuApi

	// 获取七牛云 token
	router.GET("token/qiniu", app.GetQiNiuToken)

}
