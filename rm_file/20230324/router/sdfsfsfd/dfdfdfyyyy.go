package sdfsfsfd

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DfdfdfyyyyRouter struct {
}

// InitDfdfdfyyyyRouter 初始化 Dfdfdfyyyy 路由信息
func (s *DfdfdfyyyyRouter) InitDfdfdfyyyyRouter(Router *gin.RouterGroup) {
	dfdfdfyyyRouter := Router.Group("dfdfdfyyy").Use(middleware.OperationRecord())
	dfdfdfyyyRouterWithoutRecord := Router.Group("dfdfdfyyy")
	var dfdfdfyyyApi = v1.ApiGroupApp.SdfsfsfdApiGroup.DfdfdfyyyyApi
	{
		dfdfdfyyyRouter.POST("createDfdfdfyyyy", dfdfdfyyyApi.CreateDfdfdfyyyy)   // 新建Dfdfdfyyyy
		dfdfdfyyyRouter.DELETE("deleteDfdfdfyyyy", dfdfdfyyyApi.DeleteDfdfdfyyyy) // 删除Dfdfdfyyyy
		dfdfdfyyyRouter.DELETE("deleteDfdfdfyyyyByIds", dfdfdfyyyApi.DeleteDfdfdfyyyyByIds) // 批量删除Dfdfdfyyyy
		dfdfdfyyyRouter.PUT("updateDfdfdfyyyy", dfdfdfyyyApi.UpdateDfdfdfyyyy)    // 更新Dfdfdfyyyy
	}
	{
		dfdfdfyyyRouterWithoutRecord.GET("findDfdfdfyyyy", dfdfdfyyyApi.FindDfdfdfyyyy)        // 根据ID获取Dfdfdfyyyy
		dfdfdfyyyRouterWithoutRecord.GET("getDfdfdfyyyyList", dfdfdfyyyApi.GetDfdfdfyyyyList)  // 获取Dfdfdfyyyy列表
	}
}
