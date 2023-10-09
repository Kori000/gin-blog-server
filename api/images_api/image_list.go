package images_api

import (
	"github.com/gin-gonic/gin"

	"gin-blog-server/models"
	"gin-blog-server/models/res"
	"gin-blog-server/service/common"
)

func (ImagesApi) ImageListView(c *gin.Context) {

	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)

	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	imageList, total, total_page, err := common.CommonList(models.BannerModel{}, common.Options{
		PageInfo: cr,
		Debug:    true,
	})

	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}

	res.ListResult(imageList, total, total_page, c)
}
