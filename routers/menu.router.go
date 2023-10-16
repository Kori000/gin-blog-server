package routers

import (
	"gin-blog-server/api"
)

func (router RouterGroup) MenuRouter() {
	app := api.ApiGroupApp.MenuApi

	// 菜单列表
	router.GET("menu/list", app.MenuListView)
	// 菜单创建
	router.POST("menu/create", app.MenuCreateView)
	// 菜单编辑
	// router.POST("menu/edit", app.MenuListView)
	// 菜单删除
	// router.POST("menu/remove", app.MenuListView)
}
