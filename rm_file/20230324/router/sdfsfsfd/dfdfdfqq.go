package sdfsfsfd

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DfdfdfqqqRouter struct {
}

// InitDfdfdfqqqRouter 初始化 Dfdfdfqqq 路由信息
func (s *DfdfdfqqqRouter) InitDfdfdfqqqRouter(Router *gin.RouterGroup) {
	dfdfdfqqqRouter := Router.Group("dfdfdfqqq").Use(middleware.OperationRecord())
	dfdfdfqqqRouterWithoutRecord := Router.Group("dfdfdfqqq")
	var dfdfdfqqqApi = v1.ApiGroupApp.SdfsfsfdApiGroup.DfdfdfqqqApi
	{
		dfdfdfqqqRouter.POST("createDfdfdfqqq", dfdfdfqqqApi.CreateDfdfdfqqq)   // 新建Dfdfdfqqq
		dfdfdfqqqRouter.DELETE("deleteDfdfdfqqq", dfdfdfqqqApi.DeleteDfdfdfqqq) // 删除Dfdfdfqqq
		dfdfdfqqqRouter.DELETE("deleteDfdfdfqqqByIds", dfdfdfqqqApi.DeleteDfdfdfqqqByIds) // 批量删除Dfdfdfqqq
		dfdfdfqqqRouter.PUT("updateDfdfdfqqq", dfdfdfqqqApi.UpdateDfdfdfqqq)    // 更新Dfdfdfqqq
	}
	{
		dfdfdfqqqRouterWithoutRecord.GET("findDfdfdfqqq", dfdfdfqqqApi.FindDfdfdfqqq)        // 根据ID获取Dfdfdfqqq
		dfdfdfqqqRouterWithoutRecord.GET("getDfdfdfqqqList", dfdfdfqqqApi.GetDfdfdfqqqList)  // 获取Dfdfdfqqq列表
	}
}
