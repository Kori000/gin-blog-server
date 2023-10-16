package menu_api

import (
	"github.com/gin-gonic/gin"

	"gin-blog-server/global"
	"gin-blog-server/models"
	"gin-blog-server/models/ctype"
	"gin-blog-server/models/res"
)

type ImageSort struct {
	ImageID uint `json:"image_id"`
	Sort    int  `json:"sort"`
}

type MenuRequest struct {
	MenuTitle     string      `json:"menu_title" binding:"required" msg:"请完善菜单名称"`
	MenuTitleEn   string      `json:"menu_title_en" binding:"required" msg:"请完善菜单英文名称"`
	Slogan        string      `json:"slogan"`
	Abstract      ctype.Array `json:"abstract"`      // 简介
	AbstractTime  int         `json:"abstract_time"` // 简介的切换时间
	BannerTime    int         `json:"banner_time"`   // 菜单图片的切换时间 0 表示不切换
	Sort          int         `json:"sort" binding:"required" msg:"请完善菜单序号"`
	ImageSortList []ImageSort `json:"image_sort_list"`
}

// @Tags 菜单管理
// @Summary 菜单新增
// @Description 菜单新增
// @Router /api/menu/create [POST]
// @Accept json
// @Produce json
// @Param data body MenuRequest true "菜单信息"
// @Success 200
func (MenuApi) MenuCreateView(c *gin.Context) {
	var cr MenuRequest
	err := c.ShouldBindJSON(&cr)

	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	// 创建banner数据入库

	menuModel := models.MenuModel{
		MenuTitle:    cr.MenuTitle,
		MenuTitleEn:  cr.MenuTitleEn,
		Slogan:       cr.Slogan,
		Abstract:     cr.Abstract,
		AbstractTime: cr.AbstractTime,
		BannerTime:   cr.BannerTime,
		Sort:         cr.Sort,
	}

	err = global.DB.Model(&models.MenuModel{}).Create(&menuModel).Error

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}

	if len(cr.ImageSortList) == 0 {
		res.OkWithNoData("菜单添加成功", c)
		return
	}

	// 如果请求中带有图片, 需要给给第三张表数据入库

	var menuBannerModel []models.MenuBannerModel

	for _, imageSort := range cr.ImageSortList {
		menuBannerModel = append(menuBannerModel, models.MenuBannerModel{
			MenuID:   menuModel.ID,
			BannerID: imageSort.ImageID,
			Sort:     imageSort.Sort,
		})
	}

	err = global.DB.Debug().Model(&models.MenuBannerModel{}).Create(&menuBannerModel).Error

	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}

	res.OkWithNoData("菜单添加成功", c)
}
