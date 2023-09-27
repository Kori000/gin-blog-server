package routers

import (
	"gin-blog-server/api"
)

func (router RouterGroup) SettingsRouter() {
	settingsApi := api.ApiGroupApp.SettingsApi
	router.GET("", settingsApi.SettingsInfoView)
}
