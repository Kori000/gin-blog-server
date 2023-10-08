package images_api

import (
	"github.com/gin-gonic/gin"

	"gin-blog-server/global"
	"gin-blog-server/models"
	"gin-blog-server/models/res"
)

type Page struct {
	Page  int    `form:"page"`
	Key   string `form:"key"`
	Limit int    `form:"limit"`
	Sort  string `form:"sort"`
}

func (ImagesApi) ImageListView(c *gin.Context) {
	var cr Page
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var imageList []models.BannerModel

	var count int64
	global.DB.Find(&imageList).Count(&count)

	offset := (cr.Page - 1) * cr.Limit
	if offset < 0 {
		offset = 0
	}

	global.DB.Model(&models.BannerModel{}).Limit(cr.Limit).Offset(offset)

	res.OkWithData(gin.H{
		"data":  imageList,
		"count": count,
	}, c)

}
