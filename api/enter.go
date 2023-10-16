package api

import (
	"gin-blog-server/api/advertisement_api"
	"gin-blog-server/api/banners_api"
	"gin-blog-server/api/menu_api"
	"gin-blog-server/api/settings_api"
	"gin-blog-server/api/token_api"
)

type ApiGroup struct {
	SettingsApi      settings_api.SettingsApi
	BannerApi        banners_api.BannerApi
	QiNiuApi         token_api.TokenApi
	AdvertisementApi advertisement_api.AdvertisementApi
	MenuApi          menu_api.MenuApi
}

var ApiGroupApp = new(ApiGroup)
