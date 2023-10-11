package routers

import (
	"gin-blog-server/api"
)

func (router RouterGroup) AdvertisementRouter() {
	app := api.ApiGroupApp.AdvertisementApi

	// 上传图片
	router.POST("advertisement/create", app.AdvertisementCreateView)

}
