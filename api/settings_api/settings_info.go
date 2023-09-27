package settings_api

import (
	"gin-blog-server/models/res"

	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	res.FailWithCode(1001, c)
}
