package banners_api

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"

	"gin-blog-server/global"
	"gin-blog-server/models"
	"gin-blog-server/models/ctype"
	"gin-blog-server/models/res"
)

type BannersEditBody struct {
	models.MODEL `structs:"-"`
	Name         string          `json:"name" binding:"required" msg:"请输入名称"`          // 图片名称
	OriginType   ctype.ImageType `json:"origin_type" binding:"required" msg:"请输入来源类型"` // 图片来源
}

func (BannerApi) BannersEditView(c *gin.Context) {

	var count int64
	var imageList []models.BannerModel

	var cr BannersEditBody
	err := c.ShouldBindJSON(&cr)

	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	maps := structs.Map(&cr)

	global.DB.Model(&imageList).Find(&imageList, cr.ID).Updates(maps).Count(&count)

	if count <= 0 {
		res.FailWithMessage("文件不存在", c)
		return
	}

	res.OkWith(imageList, "修改成功", c)

}
