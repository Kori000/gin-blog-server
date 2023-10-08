package settings_api

import (
	"github.com/gin-gonic/gin"

	"gin-blog-server/config"
	"gin-blog-server/core"
	"gin-blog-server/global"
	"gin-blog-server/models/res"
)

func (SettingsApi) SettingsJwtInfoView(c *gin.Context) {
	res.OkWithData(global.Config.JWT, c)
}

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
