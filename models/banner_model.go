package models

import (
	"fmt"
	"os"

	"gorm.io/gorm"

	"gin-blog-server/global"
	"gin-blog-server/models/ctype"
)

type BannerModel struct {
	MODEL
	Path       string          `json:"path" `                        // 图片路径
	Hash       string          `json:"hash"`                         // 图片 hash 值,判断重复文件
	Name       string          `json:"name" gorm:"size:255"`         // 图片名称
	OriginType ctype.ImageType `json:"origin_type" gorm:"default:1"` // 图片来源
}

func (b *BannerModel) BeforeDelete(*gorm.DB) (err error) {

	global.Log.Infof("来源对比: %v | %v", b.OriginType, ctype.Local)

	if b.OriginType == ctype.Local {
		fmt.Printf("删除本地图片: %v\n", b.Path)
		err := os.Remove(b.Path)

		if err != nil {
			global.Log.Error(err)
			return err
		}
	}
	return nil
}
