package api

import (
	"gin-blog-server/api/banners_api"
	"gin-blog-server/api/settings_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	BannerApi   banners_api.BannerApi
}

var ApiGroupApp = new(ApiGroup)
