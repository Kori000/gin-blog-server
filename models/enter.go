package models

import "time"

type MODEL struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"-"`
}

type PageInfo struct {
	Page     int    `form:"page"`
	Key      string `form:"key"`
	PageSize int    `form:"page_size"`
	Sort     string `form:"sort"`
}
