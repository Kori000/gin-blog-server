package advertisement_api

import (
	"github.com/gin-gonic/gin"

	_ "gin-blog-server/docs"
	"gin-blog-server/global"
	"gin-blog-server/models"
	"gin-blog-server/models/res"
)

type AdvertRequest struct {
	Title  string `json:"title" binding:"required" msg:"标题非法" default:"名称"`
	Href   string `json:"href" binding:"required,url" msg:"跳转链接非法" default:"https://avatars.githubusercontent.com/u/9919?s=80&v=4"`
	Images string `json:"images" binding:"required,url" msg:"图片路径非法" default:"https://avatars.githubusercontent.com/u/9919?s=80&v=4"`
	IsShow *bool  `json:"is_show" binding:"required" msg:"缺少is_show字段" default:"true"`
}

// @Tags 广告管理
// @Summary 广告新增
// @Description 广告新增
// @Router /api/advertisement/create [POST]
// @Accept json
// @Produce json
// @Param data body AdvertRequest true "广告信息"
// @Success 200
func (AdvertisementApi) AdvertisementCreateView(c *gin.Context) {
	var cr AdvertRequest

	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	var model models.AdvertModel
	// 查询是否存在同名 title
	err = global.DB.Model(&model).Take(&model, "title = ?", cr.Title).Error

	if err == nil {
		res.FailWithMessage("广告名称已存在", c)
		return
	}

	// 新建数据模型
	newModel := models.AdvertModel{
		Title:  cr.Title,
		Href:   cr.Href,
		Images: cr.Images,
		IsShow: *cr.IsShow,
	}

	err = global.DB.Model(&newModel).Create(&newModel).Error
	if err != nil {
		res.FailWithMessage("数据库错误", c)
	}

	res.OkWith(newModel, "新增成功", c)
}
