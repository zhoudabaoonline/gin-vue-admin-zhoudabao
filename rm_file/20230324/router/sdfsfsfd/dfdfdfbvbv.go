package sdfsfsfd

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DfdfdfvvvvRouter struct {
}

// InitDfdfdfvvvvRouter 初始化 Dfdfdfvvvv 路由信息
func (s *DfdfdfvvvvRouter) InitDfdfdfvvvvRouter(Router *gin.RouterGroup) {
	dfdfdfvvRouter := Router.Group("dfdfdfvv").Use(middleware.OperationRecord())
	dfdfdfvvRouterWithoutRecord := Router.Group("dfdfdfvv")
	var dfdfdfvvApi = v1.ApiGroupApp.SdfsfsfdApiGroup.DfdfdfvvvvApi
	{
		dfdfdfvvRouter.POST("createDfdfdfvvvv", dfdfdfvvApi.CreateDfdfdfvvvv)   // 新建Dfdfdfvvvv
		dfdfdfvvRouter.DELETE("deleteDfdfdfvvvv", dfdfdfvvApi.DeleteDfdfdfvvvv) // 删除Dfdfdfvvvv
		dfdfdfvvRouter.DELETE("deleteDfdfdfvvvvByIds", dfdfdfvvApi.DeleteDfdfdfvvvvByIds) // 批量删除Dfdfdfvvvv
		dfdfdfvvRouter.PUT("updateDfdfdfvvvv", dfdfdfvvApi.UpdateDfdfdfvvvv)    // 更新Dfdfdfvvvv
	}
	{
		dfdfdfvvRouterWithoutRecord.GET("findDfdfdfvvvv", dfdfdfvvApi.FindDfdfdfvvvv)        // 根据ID获取Dfdfdfvvvv
		dfdfdfvvRouterWithoutRecord.GET("getDfdfdfvvvvList", dfdfdfvvApi.GetDfdfdfvvvvList)  // 获取Dfdfdfvvvv列表
	}
}
