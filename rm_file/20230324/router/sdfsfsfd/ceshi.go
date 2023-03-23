package sdfsfsfd

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CeshiRouter struct {
}

// InitCeshiRouter 初始化 Ceshi 路由信息
func (s *CeshiRouter) InitCeshiRouter(Router *gin.RouterGroup) {
	ceshiRouter := Router.Group("ceshi").Use(middleware.OperationRecord())
	ceshiRouterWithoutRecord := Router.Group("ceshi")
	var ceshiApi = v1.ApiGroupApp.SdfsfsfdApiGroup.CeshiApi
	{
		ceshiRouter.POST("createCeshi", ceshiApi.CreateCeshi)   // 新建Ceshi
		ceshiRouter.DELETE("deleteCeshi", ceshiApi.DeleteCeshi) // 删除Ceshi
		ceshiRouter.DELETE("deleteCeshiByIds", ceshiApi.DeleteCeshiByIds) // 批量删除Ceshi
		ceshiRouter.PUT("updateCeshi", ceshiApi.UpdateCeshi)    // 更新Ceshi
	}
	{
		ceshiRouterWithoutRecord.GET("findCeshi", ceshiApi.FindCeshi)        // 根据ID获取Ceshi
		ceshiRouterWithoutRecord.GET("getCeshiList", ceshiApi.GetCeshiList)  // 获取Ceshi列表
	}
}
