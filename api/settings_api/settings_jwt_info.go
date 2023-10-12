package settings_api

import (
	"github.com/gin-gonic/gin"

	"gin-blog-server/config"
	"gin-blog-server/core"
	"gin-blog-server/global"
	"gin-blog-server/models/res"
)

// @Tags 系统管理
// @Summary 关于JWT - 获取
// @Description 关于JWT - 获取
// @Router /api/settings/jwt [GET]
// @Accept json
// @Success 200
func (SettingsApi) SettingsJwtInfoView(c *gin.Context) {
	res.OkWithData(global.Config.JWT, c)
}

// @Tags 系统管理
// @Summary 关于JWT - 编辑
// @Description 关于JWT - 编辑
// @Router /api/settings/jwt [PUT]
// @Accept json
// @Produce json
// @Param data body config.JWT true " "
// @Success 200
func (SettingsApi) SettingsJwtInfoUpdateView(c *gin.Context) {
	var cr config.JWT
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	global.Config.JWT = cr

	err = core.SetYaml()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
	}

	res.Ok(c)
}
