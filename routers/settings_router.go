package routers

import (
	"gin-blog-server/api"
)

// 系统配置
func (router RouterGroup) SettingsRouter() {
	settingsApi := api.ApiGroupApp.SettingsApi
	// 站点信息
	router.GET("settings/site", settingsApi.SettingsInfoView)
	router.PUT("settings/site", settingsApi.SettingsInfoUpdateView)
	// 邮件
	router.GET("settings/email", settingsApi.SettingsEmailInfoView)
	router.PUT("settings/email", settingsApi.SettingsEmailInfoUpdateView)
	// QQ
	router.GET("settings/qq", settingsApi.SettingsQQInfoView)
	router.PUT("settings/qq", settingsApi.SettingsQQInfoUpdateView)
	// 七牛云
	router.GET("settings/qiniu", settingsApi.SettingsQiNiuInfoView)
	router.PUT("settings/qiniu", settingsApi.SettingsQiNiuInfoUpdateView)
	// jwt
	router.GET("settings/jwt", settingsApi.SettingsJwtInfoView)
	router.PUT("settings/jwt", settingsApi.SettingsJwtInfoUpdateView)
}
