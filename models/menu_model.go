package models

import "gin-blog-server/models/ctype"

type MenuModel struct {
	MODEL
	MenuTitle    string        `gorm:"size:32" json:"menu_title"`
	MenuTitleEn  string        `gorm:"size:32" json:"menu_title_en"`
	Slogan       string        `gorm:"size:64" json:"slogan"`
	Abstract     ctype.Array   `gorm:"type:string" json:"abstract"`                                                               // 简介
	AbstractTime int           `json:"abstract_time"`                                                                             // 简介的切换时间
	Banners      []BannerModel `gorm:"many2many:menu_banner_models;joinForeignKey:MenuID;joinReferences:BannerID" json:"banners"` // 菜单图片列表
	BannerTime   int           `json:"banner_time"`                                                                               // 菜单图片的切换时间 0 表示不切换
	Sort         int           `gorm:"size:10" json:"sort"`                                                                       // 菜单的顺序
}
