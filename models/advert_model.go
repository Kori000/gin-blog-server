package models

type AdvertModel struct {
	MODEL
	Title  string `gorm:"size:32" json:"title" default:"名称"`
	Href   string `json:"href" default:"https://avatars.githubusercontent.com/u/9919?s=80&v=4"`
	Images string `json:"images" default:"https://avatars.githubusercontent.com/u/9919?s=80&v=4"`
	IsShow bool   `json:"is_show" default:"true"`
}
