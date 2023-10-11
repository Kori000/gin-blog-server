package advertisement_api

import (
	"github.com/gin-gonic/gin"

	"gin-blog-server/global"
	"gin-blog-server/models"
	"gin-blog-server/models/res"
)

type AdvertEdit struct {
	Title  string `json:"title" binding:"required" msg:"标题非法"`
	Href   string `json:"href" binding:"required,url" msg:"跳转链接非法"`
	Images string `json:"images" binding:"required,url" msg:"图片路径非法"`
	IsShow *bool  `json:"is_show" binding:"required" msg:"缺少is_show字段"`
}

type AdvertEditBody struct {
	ID int `json:"id" binding:"required" msg:"请输入id"`
	AdvertEdit
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

	global.DB.Debug().Model(&commonList).Find(&commonList, cr.ID).Updates(AdvertEdit{
		Title:  cr.Title,
		Href:   cr.Href,
		Images: cr.Images,
		IsShow: cr.IsShow,
	}).Count(&count)

	if count <= 0 {
		res.FailWithCode(res.RecordNotFoundError, c)
		return
	}

	res.OkWithData(commonList, c)
}
