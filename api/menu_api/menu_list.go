package menu_api

import (
	"github.com/gin-gonic/gin"

	"gin-blog-server/global"
	"gin-blog-server/models"
	"gin-blog-server/models/res"
)

type Banner struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
}

type MenuResponse struct {
	models.MenuModel
	Banners []Banner
}

// @Tags 菜单管理
// @Summary 菜单列表
// @Description 菜单列表
// @Router /api/menu/ list [POST]
// @Accept json
// @Produce json
// @Success 200
func (MenuApi) MenuListView(c *gin.Context) {
	var menuList []models.MenuModel // 菜单model
	var menuIDList []uint           // 取出菜单model的 id切片, 手动查询第三张表
	err := global.DB.Model(menuList).Order("sort desc").Find(&menuList).Select("id").Scan(&menuIDList).Error

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	var menuBanners []models.MenuBannerModel // 查询[菜单-banner 连接表]

	err = global.DB.Model(menuBanners).Preload("BannerModel").Order("sort").Find(&menuBanners, "menu_id in ?", menuIDList).Error

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	// 接口返回格式
	var menus []MenuResponse

	// 循环整个菜单列表
	for _, model := range menuList {
		// 定义接口需要的 Banner 格式
		var banners []Banner
		// 循环 连接表的数据
		for _, banner := range menuBanners {
			if model.ID != banner.MenuID {
				continue
			}
			banners = append(banners, Banner{
				ID:   banner.BannerID,
				Path: banner.BannerModel.Path,
			})
			menus = append(menus, MenuResponse{
				MenuModel: model,
				Banners:   banners,
			})
		}
	}

	res.OkWithData(menus, c)
}
