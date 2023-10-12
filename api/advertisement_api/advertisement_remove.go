package advertisement_api

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"gin-blog-server/global"
	"gin-blog-server/models"
	"gin-blog-server/models/res"
)

// @Tags 广告管理
// @Summary 广告删除
// @Description 广告删除
// @Router /api/advertisement/remove [POST]
// @Accept json
// @Produce json
// @Param data body models.RemoveRequest true " "
// @Success 200
func (AdvertisementApi) AdvertisementRemoveView(c *gin.Context) {
	var count int64
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)

	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var commonList []models.AdvertModel

	fmt.Printf("cr.IDList: %v\n", cr.IDList)
	global.DB.Model(&commonList).Find(&commonList, cr.IDList).Count(&count)
	if count <= 0 {
		res.FailWithMessage("数据不存在", c)
		return
	}

	global.DB.Model(&commonList).Delete(&commonList)
	fmt.Printf("count: %v\n", count)

	res.OkWithNoData(fmt.Sprintf("成功删除%v个数据", count), c)
}
