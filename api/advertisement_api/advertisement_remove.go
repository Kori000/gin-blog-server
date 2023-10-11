package advertisement_api

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"gin-blog-server/global"
	"gin-blog-server/models"
	"gin-blog-server/models/res"
)

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

	res.OkWith(commonList, fmt.Sprintf("成功删除%v个数据", count), c)
}
