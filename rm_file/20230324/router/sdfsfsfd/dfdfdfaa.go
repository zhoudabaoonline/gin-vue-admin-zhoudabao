package sdfsfsfd

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DfdfdfaaaRouter struct {
}

// InitDfdfdfaaaRouter 初始化 Dfdfdfaaa 路由信息
func (s *DfdfdfaaaRouter) InitDfdfdfaaaRouter(Router *gin.RouterGroup) {
	dfdfdfaaRouter := Router.Group("dfdfdfaa").Use(middleware.OperationRecord())
	dfdfdfaaRouterWithoutRecord := Router.Group("dfdfdfaa")
	var dfdfdfaaApi = v1.ApiGroupApp.SdfsfsfdApiGroup.DfdfdfaaaApi
	{
		dfdfdfaaRouter.POST("createDfdfdfaaa", dfdfdfaaApi.CreateDfdfdfaaa)   // 新建Dfdfdfaaa
		dfdfdfaaRouter.DELETE("deleteDfdfdfaaa", dfdfdfaaApi.DeleteDfdfdfaaa) // 删除Dfdfdfaaa
		dfdfdfaaRouter.DELETE("deleteDfdfdfaaaByIds", dfdfdfaaApi.DeleteDfdfdfaaaByIds) // 批量删除Dfdfdfaaa
		dfdfdfaaRouter.PUT("updateDfdfdfaaa", dfdfdfaaApi.UpdateDfdfdfaaa)    // 更新Dfdfdfaaa
	}
	{
		dfdfdfaaRouterWithoutRecord.GET("findDfdfdfaaa", dfdfdfaaApi.FindDfdfdfaaa)        // 根据ID获取Dfdfdfaaa
		dfdfdfaaRouterWithoutRecord.GET("getDfdfdfaaaList", dfdfdfaaApi.GetDfdfdfaaaList)  // 获取Dfdfdfaaa列表
	}
}
