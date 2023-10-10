package qiniu_api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"

	"gin-blog-server/global"
	"gin-blog-server/models/res"
)

type TokenReponse struct {
	Token   string `json:"token"`
	Expires uint64 `json:"expires"`
}

func (QiNiuApi) GetQiNiuToken(c *gin.Context) {

	accessKey := global.Config.QiNiu.AccessKey
	secretKey := global.Config.QiNiu.SecretKey

	bucket := global.Config.QiNiu.Bucket

	putPolicy := storage.PutPolicy{
		Scope: bucket,
		// 上传成功七牛云的回复
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
		// 格式化回复为 JSON 格式
		CallbackBodyType: "application/json",
	}

	mac := qbox.NewMac(accessKey, secretKey)

	token := putPolicy.UploadToken(mac)

	putPolicy.Expires = 7200

	response := TokenReponse{
		Token:   token,
		Expires: putPolicy.Expires,
	}

	fmt.Printf("过期时间是 %vs", putPolicy.Expires)
	res.OkWithData(response, c)
}
