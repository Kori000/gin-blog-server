package banners_api

import (
	"github.com/gin-gonic/gin"

	"gin-blog-server/global"
	"gin-blog-server/models"
	"gin-blog-server/models/res"
)

type ImageResponse struct {
	ID   uint   `json:"id"`
	Path string `json:"path"` // 图片路径
	Name string `json:"name"` // 图片名称
}

// @Tags 图片管理
// @Summary 图片名称列表
// @Description 获取图片名称列表
// @Router /api/banner/name/list [GET]
// @Accept json
// @Produce json
// @Success 200
func (BannerApi) BannersNameListView(c *gin.Context) {

	var imageList []ImageResponse

	global.DB.Model(models.BannerModel{}).Select("id", "name", "path").Scan(&imageList)

	res.OkWithData(imageList, c)

}
