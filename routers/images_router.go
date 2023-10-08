package routers

import (
	"gin-blog-server/api"
)

func (router RouterGroup) ImagesRouter() {
	app := api.ApiGroupApp.ImagesApi

	// 上传图片
	router.POST("upload/images", app.UploadImage)

	router.GET("banner/list", app.ImageListView)
}
