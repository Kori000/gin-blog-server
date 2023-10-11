package models

import "time"

type MODEL struct {
	ID        uint      `gorm:"primary_key" json:"id" msg:"请输入id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"-"`
}

// 列表查询
type PageInfo struct {
	Page     int    `form:"page" binding:"required"`
	PageSize int    `form:"page_size" binding:"required"`
	Sort     string `form:"sort"`
	Key      string `form:"key"`
}

// 删除请求体
type RemoveRequest struct {
	IDList []int `json:"id_list,omitempty" binding:"required,min=1"`
}
