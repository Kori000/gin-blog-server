package settings_api

import (
	"github.com/gin-gonic/gin"

	"gin-blog-server/config"
	"gin-blog-server/core"
	"gin-blog-server/global"
	"gin-blog-server/models/res"
)

// @Tags 系统配置管理
// @Summary 关于QQ - 获取
// @Description 关于QQ - 获取
// @Router /api/settings/qq [GET]
// @Accept json
// @Success 200
func (SettingsApi) SettingsQQInfoView(c *gin.Context) {
	res.OkWithData(global.Config.QQ, c)
}

// @Tags 系统配置管理
// @Summary 关于QQ - 编辑
// @Description 关于QQ - 编辑
// @Router /api/settings/qq [PUT]
// @Accept json
// @Produce json
// @Param data body config.QQ true " "
// @Success 200
func (SettingsApi) SettingsQQInfoUpdateView(c *gin.Context) {
	var cr config.QQ
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	global.Config.QQ = cr

	err = core.SetYaml()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
	}

	res.Ok(c)
}
