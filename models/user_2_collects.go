package models

import "time"

type User2Collects struct {
	UserID       uint         `gorm:"primaryKey"`
	UserModel    UserModel    `gorm:"foreignKey:UserID"`
	AritcleID    uint         `gorm:"primaryKey"`
	ArticleModel ArticleModel `gorm:"primaryKey:AritcleID"`
	CreateAt     time.Time
}
