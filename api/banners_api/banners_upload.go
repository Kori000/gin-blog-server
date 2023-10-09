package banners_api

import (
	"fmt"
	"io"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"gin-blog-server/global"
	"gin-blog-server/models"
	"gin-blog-server/models/res"
	"gin-blog-server/utils"
)

var (
	WhiteImageSuffixList = []string{
		".jpg",
		".jpeg",
		".png",
		".ico",
		".tiff",
		".svg",
		".webp",
	}
)

type UploadResponse struct {
	FileName  string `json:"file_name"`
	IsSuccess bool   `json:"is_success"`
	Msg       string `json:"msg"`
}

// 上传图片
func (BannerApi) BannersUploadView(c *gin.Context) {

	maxSize := int64(global.Config.Upload.Size * 1024 * 1024) // 文件限制大小
	bucket := time.Now().Format("2006-01-02")                 // 文件夹夹储存桶
	fileSuffix := time.Now().Format("2006-01-02T15:04:05")    // 文件时期后缀

	var responseList []UploadResponse // 响应体

	form, err := c.MultipartForm()

	if err != nil {
		res.FailWithMessage(err.Error(), c)
	}

	fileList, ok := form.File["images"]

	if !ok {
		res.FailWithMessage("文件不存在", c)
		return
	}

	for _, file := range fileList {

		fileName := file.Filename

		// 拆分文件名和后缀
		filenameWithoutExt := filepath.Base(fileName)
		ext := strings.ToLower(filepath.Ext(fileName))
		// 移除后缀
		filenameWithoutExt = filenameWithoutExt[:len(filenameWithoutExt)-len(ext)]
		// 使用Join构建新文件名
		newFilename := filenameWithoutExt + "----" + fileSuffix + ext

		// 判断后缀
		res := utils.Contains(WhiteImageSuffixList, ext)
		if !res {
			responseList = append(responseList, UploadResponse{
				FileName:  fileName,
				IsSuccess: false,
				Msg:       fmt.Sprintf("上传失败, 仅支持%s类型", strings.Join(WhiteImageSuffixList, " | ")),
			})
			continue
		}

		// 判断大小
		if file.Size > maxSize {
			responseList = append(responseList, UploadResponse{
				FileName:  fileName,
				IsSuccess: false,
				Msg:       fmt.Sprintf("文件大小超出限制, 最大 %vMB, 当前 %.2fMB", global.Config.Upload.Size, (float64(file.Size) / float64((1024 * 1024)))),
			})
			continue
		}

		fileObj, err := file.Open()
		if err != nil {
			global.Log.Error(err.Error())
		}

		byteData, err := io.ReadAll(fileObj)
		if err != nil {
			global.Log.Error(err)
			responseList = append(responseList, UploadResponse{
				FileName:  fileName,
				IsSuccess: false,
				Msg:       "服务器储存失败," + err.Error(),
			})
			continue
		}
		// md5 图片内容
		md5Hash := utils.Md5(byteData)

		// 查询数据库是否已经存在该图片

		var bannerModel models.BannerModel

		err = global.DB.Take(&bannerModel, "hash = ?", md5Hash).Error

		// 如果没报错, 就是找到了
		if err == nil {
			responseList = append(responseList, UploadResponse{
				FileName:  bannerModel.Name,
				IsSuccess: false,
				Msg:       "图片已存在",
			})
			continue
		}

		newFilePath := path.Join(global.Config.Upload.ImagePath+"/"+bucket, newFilename)

		c.SaveUploadedFile(file, newFilePath)

		responseList = append(responseList, UploadResponse{
			FileName:  fileName,
			IsSuccess: true,
			Msg:       "上传成功",
		})

		// 图片入库
		global.DB.Create(&models.BannerModel{
			Path: newFilePath,
			Hash: md5Hash,
			Name: fileName,
		})

	}

	res.OkWithData(responseList, c)
}
