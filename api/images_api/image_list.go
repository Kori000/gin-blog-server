package images_api

import (
	"fmt"
	"math"

	"github.com/gin-gonic/gin"

	"gin-blog-server/global"
	"gin-blog-server/models"
	"gin-blog-server/models/res"
)

type Page struct {
	Page     int    `form:"page"`
	Key      string `form:"key"`
	PageSize int    `form:"page_size"`
	Sort     string `form:"sort"`
}

func (ImagesApi) ImageListView(c *gin.Context) {
	var total int64                    // 总条数
	var total_page int64               // 总页数
	var imageList []models.BannerModel // 数据模型
	var cr Page
	err := c.ShouldBindQuery(&cr)

	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	// 总条数 -- 最佳实践是前面都指定Model;
	global.DB.Model(&imageList).Find(&imageList).Count(&total)

	// 计算总页数
	total_page = int64(math.Ceil(float64(total) / float64(cr.PageSize)))

	if int64(cr.Page) > total_page {
		res.FailWithMessage(fmt.Sprintf("超过最大页数,最大页数为 %v", total_page), c)
		return
	}

	offset := (cr.Page - 1) * cr.PageSize
	if offset < 0 {
		offset = 0
	}

	// 查询
	global.DB.Model(&imageList).Limit(cr.PageSize).Offset(offset).Find(&imageList)

	res.ListResult(imageList, total, total_page, c)
}
