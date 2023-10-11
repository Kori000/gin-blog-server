package advertisement_api

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"

	"gin-blog-server/global"
	"gin-blog-server/models"
	"gin-blog-server/models/res"
)

type AdvertEditBody struct {
	models.MODEL `structs:"-"`
	Title        string `json:"title" binding:"required" msg:"标题非法" structs:"title"`
	Href         string `json:"href" binding:"required,url" msg:"跳转链接非法" structs:"href"`
	Images       string `json:"images" binding:"required,url" msg:"图片路径非法" structs:"images"`
	IsShow       *bool  `json:"is_show" binding:"required" msg:"缺少is_show字段" structs:"is_show"`
}

func (AdvertisementApi) AdvertisementEditView(c *gin.Context) {

	var commonList models.AdvertModel
	var cr AdvertEditBody
	var count int64

	err := c.ShouldBindJSON(&cr)

	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	maps := structs.Map(&cr)

	global.DB.Debug().Model(&commonList).Find(&commonList, cr.ID).Updates(maps).Count(&count)

	if count <= 0 {
		res.FailWithCode(res.RecordNotFoundError, c)
		return
	}

	res.OkWithData(commonList, c)
}
