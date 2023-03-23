package sdfsfsfd

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DfdfdfzzRouter struct {
}

// InitDfdfdfzzRouter 初始化 Dfdfdfzz 路由信息
func (s *DfdfdfzzRouter) InitDfdfdfzzRouter(Router *gin.RouterGroup) {
	dfdfdfzzzzRouter := Router.Group("dfdfdfzzzz").Use(middleware.OperationRecord())
	dfdfdfzzzzRouterWithoutRecord := Router.Group("dfdfdfzzzz")
	var dfdfdfzzzzApi = v1.ApiGroupApp.SdfsfsfdApiGroup.DfdfdfzzApi
	{
		dfdfdfzzzzRouter.POST("createDfdfdfzz", dfdfdfzzzzApi.CreateDfdfdfzz)   // 新建Dfdfdfzz
		dfdfdfzzzzRouter.DELETE("deleteDfdfdfzz", dfdfdfzzzzApi.DeleteDfdfdfzz) // 删除Dfdfdfzz
		dfdfdfzzzzRouter.DELETE("deleteDfdfdfzzByIds", dfdfdfzzzzApi.DeleteDfdfdfzzByIds) // 批量删除Dfdfdfzz
		dfdfdfzzzzRouter.PUT("updateDfdfdfzz", dfdfdfzzzzApi.UpdateDfdfdfzz)    // 更新Dfdfdfzz
	}
	{
		dfdfdfzzzzRouterWithoutRecord.GET("findDfdfdfzz", dfdfdfzzzzApi.FindDfdfdfzz)        // 根据ID获取Dfdfdfzz
		dfdfdfzzzzRouterWithoutRecord.GET("getDfdfdfzzList", dfdfdfzzzzApi.GetDfdfdfzzList)  // 获取Dfdfdfzz列表
	}
}
