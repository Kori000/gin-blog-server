package menu_api

import (
	"github.com/gin-gonic/gin"

	"gin-blog-server/global"
	"gin-blog-server/models"
	"gin-blog-server/models/res"
)

type MenuNameResponse struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Path  string `json:"path"`
}

// @Tags 菜单管理
// @Summary 菜单列表
// @Description 菜单列表
// @Router /api/menu/name/list [GET]
// @Accept json
// @Produce json
// @Success 200
func (MenuApi) MenuNameListView(c *gin.Context) {

	var menuNameList []MenuNameResponse

	global.DB.Model(models.MenuModel{}).Select("id", "title", "path").Scan(&menuNameList)

	res.OkWithData(menuNameList, c)
}
