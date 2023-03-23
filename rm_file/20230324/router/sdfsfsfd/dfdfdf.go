package sdfsfsfd

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DfdfdfRouter struct {
}

// InitDfdfdfRouter 初始化 Dfdfdf 路由信息
func (s *DfdfdfRouter) InitDfdfdfRouter(Router *gin.RouterGroup) {
	dfdfdfRouter := Router.Group("dfdfdf").Use(middleware.OperationRecord())
	dfdfdfRouterWithoutRecord := Router.Group("dfdfdf")
	var dfdfdfApi = v1.ApiGroupApp.SdfsfsfdApiGroup.DfdfdfApi
	{
		dfdfdfRouter.POST("createDfdfdf", dfdfdfApi.CreateDfdfdf)   // 新建Dfdfdf
		dfdfdfRouter.DELETE("deleteDfdfdf", dfdfdfApi.DeleteDfdfdf) // 删除Dfdfdf
		dfdfdfRouter.DELETE("deleteDfdfdfByIds", dfdfdfApi.DeleteDfdfdfByIds) // 批量删除Dfdfdf
		dfdfdfRouter.PUT("updateDfdfdf", dfdfdfApi.UpdateDfdfdf)    // 更新Dfdfdf
	}
	{
		dfdfdfRouterWithoutRecord.GET("findDfdfdf", dfdfdfApi.FindDfdfdf)        // 根据ID获取Dfdfdf
		dfdfdfRouterWithoutRecord.GET("getDfdfdfList", dfdfdfApi.GetDfdfdfList)  // 获取Dfdfdf列表
	}
}
