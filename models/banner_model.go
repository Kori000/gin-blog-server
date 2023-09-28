package models

type BannerModel struct {
	MODEL
	Path string `json:"path"`                 // 图片路径
	Hash string `json:"hash"`                 // 图片 hash 值,判断重复文件
	Name string `gorm:"size:255" json:"name"` // 图片名称
}
