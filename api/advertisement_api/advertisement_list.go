package advertisement_api

import (
	"strings"

	"github.com/gin-gonic/gin"

	"gin-blog-server/models"
	"gin-blog-server/models/res"
	"gin-blog-server/service/common"
)

// @Tags 广告管理
// @Summary 广告列表
// @Description 广告列表
// @Router /api/advertisement/list [GET]
// @Accept json
// @Produce json
// @Param data query models.PageInfo true " "
// @Success 200
func (AdvertisementApi) AdvertisementListView(c *gin.Context) {

	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)

	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}

	// 判断请求头的 referer 来源, 若为前台, 只返回 is_show 为 true
	Referer := c.GetHeader("Referer")

	// Referer包含 admin 字符
	showAllList := strings.Contains(Referer, "admin")

	advertisementList, total, total_page, err := common.CommonList[models.AdvertModel](models.AdvertModel{
		IsShow: showAllList,
	}, common.Options{
		PageInfo: cr,
		Debug:    true,
	})

	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}

	res.ListResult(advertisementList, total, total_page, c)
}
