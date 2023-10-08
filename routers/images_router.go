package routers

import (
	"gin-blog-server/api"
)

func (router RouterGroup) ImagesRouter() {
	app := api.ApiGroupApp.ImagesApi

	router.POST("images", app.UploadImage)
}
