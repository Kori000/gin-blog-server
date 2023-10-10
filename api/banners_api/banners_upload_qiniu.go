package banners_api

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"

	"gin-blog-server/global"
	"gin-blog-server/models"
	"gin-blog-server/models/res"
	"gin-blog-server/utils"
)

// 对象存储方式的请求体
type OssUploadBody struct {
	ImageList []string `json:"image_list" binding:"required"`
}

// 上传图片 - 七牛云
func (BannerApi) BannersUploadQiNiuView(c *gin.Context) {

	var responseList []UploadResponse // 响应体

	var fileList OssUploadBody

	err := c.ShouldBindJSON(&fileList)

	if err != nil {
		res.FailWithMessage(err.Error(), c)
	}

	for _, file := range fileList.ImageList {

		fileName := strings.Replace(file, "\u0026", "&", -1)
		fmt.Printf("✅: %v\n", fileName)

		// md5 图片内容
		md5Hash := utils.Md5([]byte(fileName))

		// 必须每次循环的时候创建
		var bannerModel models.BannerModel

		// 查询数据库是否已经存在该图片
		err = global.DB.Debug().Take(&bannerModel, "hash = ?", md5Hash).Error

		// 如果没报错, 就是找到了
		if err == nil {
			responseList = append(responseList, UploadResponse{
				FileName:  bannerModel.Name,
				IsSuccess: false,
				Msg:       "图片已存在",
			})
			continue
		}

		// 上传成功
		responseList = append(responseList, UploadResponse{
			FileName:  utils.TruncateStr(fileName, 28),
			IsSuccess: true,
			Msg:       "上传成功",
		})

		// 图片入库
		global.DB.Create(&models.BannerModel{
			Path:       fileName,
			Hash:       md5Hash,
			Name:       utils.TruncateStr(fileName, 28),
			OriginType: 2,
		})

	}

	res.OkWithData(responseList, c)
}
