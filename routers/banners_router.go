package routers

import (
	"gin-blog-server/api"
)

func (router RouterGroup) ImagesRouter() {
	app := api.ApiGroupApp.BannerApi

	qiniu := api.ApiGroupApp.QiNiuApi

	// 获取 banner 图片列表
	router.GET("banner/list", app.BannersListView)

	// 上传图片
	router.POST("banner/upload", app.BannersUploadView)

	// 删除图片
	router.POST("banner/remove", app.BannersRemoveView)

	// 编辑图片
	router.POST("banner/edit", app.BannersEditView)

	// 获取七牛云 token
	router.GET("token/qiniu", qiniu.GetQiNiuToken)

	// 上传图片 - 七牛云地址
	router.POST("banner/upload/url", app.BannersUploadQiNiuView)
}
