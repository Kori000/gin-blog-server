package images_api

import (
	"path"
	"path/filepath"
	"strconv"

	"time"

	"github.com/gin-gonic/gin"

	"gin-blog-server/global"
	"gin-blog-server/models/res"
)

type UploadResponse struct {
	FileName  string `json:"file_name"`
	IsSuccess bool   `json:"is_success"`
	Msg       string `json:"msg"`
}

// 上传单个图片, 返回图片 url
func (ImagesApi) UploadImage(c *gin.Context) {

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
		res.FailWithMessage("images文件不存在", c)
	}

	for _, file := range fileList {

		// 拆分文件名和后缀
		filenameWithoutExt := filepath.Base(file.Filename)
		ext := filepath.Ext(file.Filename)
		// 移除后缀
		filenameWithoutExt = filenameWithoutExt[:len(filenameWithoutExt)-len(ext)]
		// 使用Join构建新文件名
		newFilename := filenameWithoutExt + "----" + fileSuffix + ext

		// 判断大小
		if file.Size > maxSize {
			responseList = append(responseList, UploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       "文件大小超出限制," + "最大" + strconv.Itoa(global.Config.Upload.Size) + "mb",
			})
			continue
		}

		newFilePath := path.Join(global.Config.Upload.ImagePath+"/"+bucket, newFilename)

		err := c.SaveUploadedFile(file, newFilePath)
		if err != nil {
			global.Log.Error(err)
			responseList = append(responseList, UploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       "服务器储存失败," + err.Error(),
			})
			continue
		}

		responseList = append(responseList, UploadResponse{
			FileName:  file.Filename,
			IsSuccess: true,
			Msg:       "上传成功",
		})
	}

	res.OkWithData(responseList, c)
}
