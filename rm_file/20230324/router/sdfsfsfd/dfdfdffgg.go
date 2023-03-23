package sdfsfsfd

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DfdfdffgRouter struct {
}

// InitDfdfdffgRouter 初始化 Dfdfdffg 路由信息
func (s *DfdfdffgRouter) InitDfdfdffgRouter(Router *gin.RouterGroup) {
	dfdfdffgRouter := Router.Group("dfdfdffg").Use(middleware.OperationRecord())
	dfdfdffgRouterWithoutRecord := Router.Group("dfdfdffg")
	var dfdfdffgApi = v1.ApiGroupApp.SdfsfsfdApiGroup.DfdfdffgApi
	{
		dfdfdffgRouter.POST("createDfdfdffg", dfdfdffgApi.CreateDfdfdffg)   // 新建Dfdfdffg
		dfdfdffgRouter.DELETE("deleteDfdfdffg", dfdfdffgApi.DeleteDfdfdffg) // 删除Dfdfdffg
		dfdfdffgRouter.DELETE("deleteDfdfdffgByIds", dfdfdffgApi.DeleteDfdfdffgByIds) // 批量删除Dfdfdffg
		dfdfdffgRouter.PUT("updateDfdfdffg", dfdfdffgApi.UpdateDfdfdffg)    // 更新Dfdfdffg
	}
	{
		dfdfdffgRouterWithoutRecord.GET("findDfdfdffg", dfdfdffgApi.FindDfdfdffg)        // 根据ID获取Dfdfdffg
		dfdfdffgRouterWithoutRecord.GET("getDfdfdffgList", dfdfdffgApi.GetDfdfdffgList)  // 获取Dfdfdffg列表
	}
}
