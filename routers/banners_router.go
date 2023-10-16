package routers

import (
	"gin-blog-server/api"
)

func (router RouterGroup) ImagesRouter() {
	app := api.ApiGroupApp.BannerApi

	// 获取 banner 图片列表
	router.GET("banner/list", app.BannersListView)

	// 获取 banner 图片名称列表
	router.GET("banner/name/list", app.BannersNameListView)

	// 上传图片
	router.POST("banner/upload", app.BannersUploadView)

	// 删除图片
	router.POST("banner/remove", app.BannersRemoveView)

	// 编辑图片
	router.POST("banner/edit", app.BannersEditView)

	// 上传图片 - 七牛云地址
	router.POST("banner/upload/url", app.BannersUploadQiNiuView)

}
