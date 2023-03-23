package sdfsfsfd

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type Dfdfdf4Router struct {
}

// InitDfdfdf4Router 初始化 Dfdfdf4 路由信息
func (s *Dfdfdf4Router) InitDfdfdf4Router(Router *gin.RouterGroup) {
	dfdfdf4Router := Router.Group("dfdfdf4").Use(middleware.OperationRecord())
	dfdfdf4RouterWithoutRecord := Router.Group("dfdfdf4")
	var dfdfdf4Api = v1.ApiGroupApp.SdfsfsfdApiGroup.Dfdfdf4Api
	{
		dfdfdf4Router.POST("createDfdfdf4", dfdfdf4Api.CreateDfdfdf4)   // 新建Dfdfdf4
		dfdfdf4Router.DELETE("deleteDfdfdf4", dfdfdf4Api.DeleteDfdfdf4) // 删除Dfdfdf4
		dfdfdf4Router.DELETE("deleteDfdfdf4ByIds", dfdfdf4Api.DeleteDfdfdf4ByIds) // 批量删除Dfdfdf4
		dfdfdf4Router.PUT("updateDfdfdf4", dfdfdf4Api.UpdateDfdfdf4)    // 更新Dfdfdf4
	}
	{
		dfdfdf4RouterWithoutRecord.GET("findDfdfdf4", dfdfdf4Api.FindDfdfdf4)        // 根据ID获取Dfdfdf4
		dfdfdf4RouterWithoutRecord.GET("getDfdfdf4List", dfdfdf4Api.GetDfdfdf4List)  // 获取Dfdfdf4列表
	}
}
