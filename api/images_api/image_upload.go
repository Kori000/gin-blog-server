package images_api

import (
	"fmt"
	"mime/multipart"

	"github.com/gin-gonic/gin"

	"gin-blog-server/models/res"
)

// 上传单个图片, 返回图片 url
func (ImagesApi) UploadImage(c *gin.Context) {

	form, err := c.MultipartForm()

	if err != nil {
		res.FailWithMessage(err.Error(), c)
	}

	fileList, ok := form.File["images"]

	if !ok {
		res.FailWithMessage("不存在", c)
	}

	var files []*multipart.FileHeader

	for _, file := range fileList {
		fmt.Println(file.Header)
		fmt.Println(file.Size)
		fmt.Println(file.Filename)
		files = append(files, file)

	}
	res.OkWithData(files, c)

}
