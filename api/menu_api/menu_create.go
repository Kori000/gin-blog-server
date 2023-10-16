package menu_api

import (
	"github.com/gin-gonic/gin"

	"gin-blog-server/models/ctype"
)

type ImageSort struct {
	ImageID uint `json:"image_id"`
	Sort    int  `json:"sort"`
}

type MenuRequest struct {
	MenuTitle     string      `json:"menu_title"`
	MenuTitleEn   string      `json:"menu_title_en"`
	Slogan        string      `json:"slogan"`
	Abstract      ctype.Array `json:"abstract"`      // 简介
	AbstractTime  int         `json:"abstract_time"` // 简介的切换时间
	BannerTime    int         `json:"banner_time"`   // 菜单图片的切换时间 0 表示不切换
	Sort          int         `json:"sort"`
	ImageSortList []ImageSort `json:"image_sort_list"`
}

func (MenuApi) MenuCreateView(c *gin.Context) {

}
