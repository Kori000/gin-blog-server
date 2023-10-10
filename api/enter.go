package api

import (
	"gin-blog-server/api/banners_api"
	"gin-blog-server/api/qiniu_api"
	"gin-blog-server/api/settings_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	BannerApi   banners_api.BannerApi
	QiNiuApi    qiniu_api.QiNiuApi
}

var ApiGroupApp = new(ApiGroup)
