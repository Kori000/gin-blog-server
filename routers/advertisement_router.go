package routers

import (
	"gin-blog-server/api"
)

// @Tags 广告管理
func (router RouterGroup) AdvertisementRouter() {
	app := api.ApiGroupApp.AdvertisementApi

	// 广告列表
	router.GET("advertisement/list", app.AdvertisementListView)
	// 广告创建
	router.POST("advertisement/create", app.AdvertisementCreateView)
	// 广告编辑
	router.POST("advertisement/edit", app.AdvertisementEditView)
	// 广告删除
	router.POST("advertisement/remove", app.AdvertisementRemoveView)
}
