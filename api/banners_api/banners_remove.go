package banners_api

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"gin-blog-server/global"
	"gin-blog-server/models"
	"gin-blog-server/models/res"
)

func (BannerApi) BannersRemoveView(c *gin.Context) {
	var count int64
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)

	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var imageList []models.BannerModel

	fmt.Printf("cr.IDList: %v\n", cr.IDList)
	global.DB.Model(&imageList).Find(&imageList, cr.IDList).Count(&count)
	if count <= 0 {
		res.FailWithMessage("文件不存在", c)
		return
	}

	global.DB.Model(&imageList).Delete(&imageList)
	fmt.Printf("count: %v\n", count)

	res.OkWith(imageList, fmt.Sprintf("成功删除%v个文件", count), c)
}
