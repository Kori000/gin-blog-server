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
	Title        string `json:"title" binding:"required" msg:"标题非法" structs:"title" default:"名称"`
	Href         string `json:"href" binding:"required,url" msg:"跳转链接非法" structs:"href" default:"https://avatars.githubusercontent.com/u/9919?s=80&v=4"`
	Images       string `json:"images" binding:"required,url" msg:"图片路径非法" structs:"images" default:"https://avatars.githubusercontent.com/u/9919?s=80&v=4"`
	IsShow       *bool  `json:"is_show" binding:"required" msg:"缺少is_show字段" structs:"is_show" default:"true"`
}

// @Tags 广告管理
// @Summary 广告编辑
// @Description 广告编辑
// @Router /api/advertisement/edit [POST]
// @Accept json
// @Produce json
// @Param data body AdvertEditBody true " "
// @Success 200
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

	res.OkWith(commonList, "编辑成功", c)

}
