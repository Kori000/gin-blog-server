package advertisement_api

import (
	"gin-blog-server/models/res"

	"github.com/gin-gonic/gin"
)

func (AdvertisementApi) AdvertisementRemoveView(c *gin.Context) {
	res.OkWithData("删除成功", c)
}
