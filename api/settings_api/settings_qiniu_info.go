package settings_api

import (
	"github.com/gin-gonic/gin"

	"gin-blog-server/config"
	"gin-blog-server/core"
	"gin-blog-server/global"
	"gin-blog-server/models/res"
)

// @Tags 系统配置管理
// @Summary 关于七牛云 - 获取
// @Description 关于七牛云 - 获取
// @Router /api/settings/qiniu [GET]
// @Accept json
// @Success 200
func (SettingsApi) SettingsQiNiuInfoView(c *gin.Context) {
	res.OkWithData(global.Config.QiNiu, c)
}

// @Tags 系统配置管理
// @Summary 关于七牛云 - 编辑
// @Description 关于七牛云 - 编辑
// @Router /api/settings/qiniu [PUT]
// @Accept json
// @Produce json
// @Param data body config.QiNiu true " "
// @Success 200
func (SettingsApi) SettingsQiNiuInfoUpdateView(c *gin.Context) {
	var cr config.QiNiu
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	global.Config.QiNiu = cr

	err = core.SetYaml()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
	}

	res.Ok(c)
}
