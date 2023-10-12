package models

import "time"

type MODEL struct {
	ID        uint      `gorm:"primary_key" json:"id" msg:"请输入id"`
	CreatedAt time.Time `json:"created_at" swaggerignore:"true"`
	UpdatedAt time.Time `json:"-" swaggerignore:"true"`
}

// 列表查询
type PageInfo struct {
	Page     int    `form:"page" binding:"required" default:"1"`
	PageSize int    `form:"page_size" binding:"required" default:"10"`
	Sort     string `form:"sort" swaggerignore:"true"`
	Key      string `form:"key" swaggerignore:"true"`
}

// 删除请求体
type RemoveRequest struct {
	IDList []int `json:"id_list,omitempty" binding:"required,min=1" `
}
