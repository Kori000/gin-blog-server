package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (router RouterGroup) TestRouter() {

	router.GET("/test", FnTest)

}

// @Tags 测试模块
// @Summary 功能测试
// @Description 功能测试
// @Router /api/test [GET]
// @Accept json
// @Produce json
// @Success 200
func FnTest(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "https://google.com")
}
