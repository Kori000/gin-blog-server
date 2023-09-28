package models

type MenuImageMmodel struct {
	MenuID      uint        `json:"menu_id"`
	MenuTitleEn MenuModel   `gorm:"foreignKey:MenuID"`
	BannerID    uint        `json:"banner_id"`
	BannerModel BannerModel `gorm:"foreignKey:BannerID"`
	Sort        int         `gorm:"size:10" json:"sort"`
}
