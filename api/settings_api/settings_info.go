package settings_api

import (
	"github.com/gin-gonic/gin"

	"gin-blog-server/config"
	"gin-blog-server/core"
	"gin-blog-server/global"
	"gin-blog-server/models/res"
)

// @Tags 系统配置管理
// @Summary 关于站点 - 获取
// @Description 关于站点 - 获取
// @Router /api/settings/site [GET]
// @Accept json
// @Success 200
func (SettingsApi) SettingsInfoView(c *gin.Context) {
	res.OkWithData(global.Config.SiteInfo, c)
}

// @Tags 系统配置管理
// @Summary 关于站点 - 编辑
// @Description 关于站点 - 编辑
// @Router /api/settings/site [PUT]
// @Accept json
// @Produce json
// @Param data body config.SiteInfo true " "
// @Success 200
func (SettingsApi) SettingsInfoUpdateView(c *gin.Context) {
	var cr config.SiteInfo
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	global.Config.SiteInfo = cr

	err = core.SetYaml()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
	}
	res.Ok(c)
}
