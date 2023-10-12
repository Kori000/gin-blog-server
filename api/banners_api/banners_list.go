package banners_api

import (
	"github.com/gin-gonic/gin"

	"gin-blog-server/models"
	"gin-blog-server/models/res"
	"gin-blog-server/service/common"
)

// @Tags 图片管理
// @Summary 图片列表
// @Description 图片列表
// @Router /api/banner/list [GET]
// @Accept json
// @Produce json
// @Param data query models.PageInfo true " "
// @Success 200
func (BannerApi) BannersListView(c *gin.Context) {

	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)

	if err != nil {
		res.FailWithMessage(err.Error(), c)
		// res.FailWithCode(res.ArgumentError, c)
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
