package common

import (
	"fmt"
	"math"

	"gorm.io/gorm"

	"gin-blog-server/global"
	"gin-blog-server/models"
)

type Options struct {
	models.PageInfo
	Debug bool
}

func CommonList[T any](model T, option Options) (list []T, total int64, total_page int64, err error) {
	var DB *gorm.DB
	fmt.Printf("option.Debug: %v\n", option.Debug)
	if option.Debug {
		DB = global.DB.Session(&gorm.Session{
			Logger: global.MysqlLog,
		})
	} else {
		DB = global.DB
	}

	global.Log.Infoln("✅", model)
	query := DB.Where(model)
	// 总条数 -- 最佳实践是前面都指定Model;
	query.Model(&list).Select("id").Count(&total)

	// Select 会改变 model, 这里再一次 Where 进行复位
	query = DB.Where(model)
	// 计算总页数
	total_page = int64(math.Ceil(float64(total) / float64(option.PageInfo.PageSize)))

	// 是否超出总页数
	if int64(option.PageInfo.Page) > total_page && int64(option.PageInfo.Page) != 1 {
		err = fmt.Errorf(fmt.Sprintf("超过最大页数,最大页数为 %v", total_page))
		return
	}

	//  计算偏移量
	offset := (option.PageInfo.Page - 1) * option.PageInfo.PageSize

	// 如果小于 0 就从 0 开始
	if offset < 0 {
		offset = 0
	}

	if option.Sort == "" {
		option.Sort = "created_at desc" // 默认时间排序
	}

	// 查询
	err = query.Model(&list).Limit(option.PageInfo.PageSize).Offset(offset).Order(option.Sort).Find(&list).Error

	return list, total, total_page, err

}
