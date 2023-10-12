package settings_api

import (
	"github.com/gin-gonic/gin"

	"gin-blog-server/config"
	"gin-blog-server/core"
	"gin-blog-server/global"
	"gin-blog-server/models/res"
)

// @Tags 系统管理
// @Summary 关于邮件 - 获取
// @Description 关于邮件 - 获取
// @Router /api/settings/email [GET]
// @Accept json
// @Success 200
func (SettingsApi) SettingsEmailInfoView(c *gin.Context) {
	res.OkWithData(global.Config.Email, c)
}

// @Tags 系统管理
// @Summary 关于邮件 - 编辑
// @Description 关于邮件 - 编辑
// @Router /api/settings/email [PUT]
// @Accept json
// @Produce json
// @Param data body config.Email true " "
// @Success 200
func (SettingsApi) SettingsEmailInfoUpdateView(c *gin.Context) {
	var cr config.Email

	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	global.Config.Email = cr

	err = core.SetYaml()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}

	res.Ok(c)
}
