package settings_api

import (
	"github.com/gin-gonic/gin"

	"gin-blog-server/config"
	"gin-blog-server/core"
	"gin-blog-server/global"
	"gin-blog-server/models/res"
)

func (SettingsApi) SettingsEmailInfoView(c *gin.Context) {
	res.OkWithData(global.Config.Email, c)
}

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
