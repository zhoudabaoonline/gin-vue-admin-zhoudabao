package sdfsfsfd

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DfdfdfvvvRouter struct {
}

// InitDfdfdfvvvRouter 初始化 Dfdfdfvvv 路由信息
func (s *DfdfdfvvvRouter) InitDfdfdfvvvRouter(Router *gin.RouterGroup) {
	dfdfdfvvvvRouter := Router.Group("dfdfdfvvvv").Use(middleware.OperationRecord())
	dfdfdfvvvvRouterWithoutRecord := Router.Group("dfdfdfvvvv")
	var dfdfdfvvvvApi = v1.ApiGroupApp.SdfsfsfdApiGroup.DfdfdfvvvApi
	{
		dfdfdfvvvvRouter.POST("createDfdfdfvvv", dfdfdfvvvvApi.CreateDfdfdfvvv)   // 新建Dfdfdfvvv
		dfdfdfvvvvRouter.DELETE("deleteDfdfdfvvv", dfdfdfvvvvApi.DeleteDfdfdfvvv) // 删除Dfdfdfvvv
		dfdfdfvvvvRouter.DELETE("deleteDfdfdfvvvByIds", dfdfdfvvvvApi.DeleteDfdfdfvvvByIds) // 批量删除Dfdfdfvvv
		dfdfdfvvvvRouter.PUT("updateDfdfdfvvv", dfdfdfvvvvApi.UpdateDfdfdfvvv)    // 更新Dfdfdfvvv
	}
	{
		dfdfdfvvvvRouterWithoutRecord.GET("findDfdfdfvvv", dfdfdfvvvvApi.FindDfdfdfvvv)        // 根据ID获取Dfdfdfvvv
		dfdfdfvvvvRouterWithoutRecord.GET("getDfdfdfvvvList", dfdfdfvvvvApi.GetDfdfdfvvvList)  // 获取Dfdfdfvvv列表
	}
}
