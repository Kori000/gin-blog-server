package qiniu

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"

	"gin-blog-server/global"
	"gin-blog-server/models/res"
)

// 服务器直传 - 暂未实现
func QiNiuUpload(c *gin.Context) {

	var (
		accessKey = global.Config.QiNiu.AccessKey
		secretKey = global.Config.QiNiu.SecretKey
		bucket    = global.Config.QiNiu.Bucket
		localFile = "uploads/images/2023-10-09/door----2023-10-09T20:01:33.png"
		key       = "github-xx.png"
	)

	putPolicy := storage.PutPolicy{
		Scope: bucket,
		// 上传成功七牛云的回复
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
		// 格式化回复为 JSON 格式
		CallbackBodyType: "application/json",
	}

	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	// cfg.Region = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	// 可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		global.Log.Errorln("上传错误  ", err)
		return
	}

	global.Log.Infoln("上传成功  ", ret)

	res.OkWithData(ret, c)
}
