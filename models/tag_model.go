package models

type TagModel struct {
	MODEL
	Title  string `gorm:"size:16" json:"title"`
	Articl string `gorm:"many2many:article_tag" json:"-"`
}
