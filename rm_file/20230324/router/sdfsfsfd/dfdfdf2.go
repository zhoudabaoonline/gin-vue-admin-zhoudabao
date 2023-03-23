package sdfsfsfd

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type Dfdfdf22Router struct {
}

// InitDfdfdf22Router 初始化 Dfdfdf22 路由信息
func (s *Dfdfdf22Router) InitDfdfdf22Router(Router *gin.RouterGroup) {
	dfdfdf22Router := Router.Group("dfdfdf22").Use(middleware.OperationRecord())
	dfdfdf22RouterWithoutRecord := Router.Group("dfdfdf22")
	var dfdfdf22Api = v1.ApiGroupApp.SdfsfsfdApiGroup.Dfdfdf22Api
	{
		dfdfdf22Router.POST("createDfdfdf22", dfdfdf22Api.CreateDfdfdf22)   // 新建Dfdfdf22
		dfdfdf22Router.DELETE("deleteDfdfdf22", dfdfdf22Api.DeleteDfdfdf22) // 删除Dfdfdf22
		dfdfdf22Router.DELETE("deleteDfdfdf22ByIds", dfdfdf22Api.DeleteDfdfdf22ByIds) // 批量删除Dfdfdf22
		dfdfdf22Router.PUT("updateDfdfdf22", dfdfdf22Api.UpdateDfdfdf22)    // 更新Dfdfdf22
	}
	{
		dfdfdf22RouterWithoutRecord.GET("findDfdfdf22", dfdfdf22Api.FindDfdfdf22)        // 根据ID获取Dfdfdf22
		dfdfdf22RouterWithoutRecord.GET("getDfdfdf22List", dfdfdf22Api.GetDfdfdf22List)  // 获取Dfdfdf22列表
	}
}
