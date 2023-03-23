package sdfsfsfd

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DfdfdfccRouter struct {
}

// InitDfdfdfccRouter 初始化 Dfdfdfcc 路由信息
func (s *DfdfdfccRouter) InitDfdfdfccRouter(Router *gin.RouterGroup) {
	dfdfdfccRouter := Router.Group("dfdfdfcc").Use(middleware.OperationRecord())
	dfdfdfccRouterWithoutRecord := Router.Group("dfdfdfcc")
	var dfdfdfccApi = v1.ApiGroupApp.SdfsfsfdApiGroup.DfdfdfccApi
	{
		dfdfdfccRouter.POST("createDfdfdfcc", dfdfdfccApi.CreateDfdfdfcc)   // 新建Dfdfdfcc
		dfdfdfccRouter.DELETE("deleteDfdfdfcc", dfdfdfccApi.DeleteDfdfdfcc) // 删除Dfdfdfcc
		dfdfdfccRouter.DELETE("deleteDfdfdfccByIds", dfdfdfccApi.DeleteDfdfdfccByIds) // 批量删除Dfdfdfcc
		dfdfdfccRouter.PUT("updateDfdfdfcc", dfdfdfccApi.UpdateDfdfdfcc)    // 更新Dfdfdfcc
	}
	{
		dfdfdfccRouterWithoutRecord.GET("findDfdfdfcc", dfdfdfccApi.FindDfdfdfcc)        // 根据ID获取Dfdfdfcc
		dfdfdfccRouterWithoutRecord.GET("getDfdfdfccList", dfdfdfccApi.GetDfdfdfccList)  // 获取Dfdfdfcc列表
	}
}
