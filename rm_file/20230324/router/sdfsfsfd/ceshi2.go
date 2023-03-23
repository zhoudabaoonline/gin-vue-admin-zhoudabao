package sdfsfsfd

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type Ceshi2Router struct {
}

// InitCeshi2Router 初始化 Ceshi2 路由信息
func (s *Ceshi2Router) InitCeshi2Router(Router *gin.RouterGroup) {
	ceshi2Router := Router.Group("ceshi2").Use(middleware.OperationRecord())
	ceshi2RouterWithoutRecord := Router.Group("ceshi2")
	var ceshi2Api = v1.ApiGroupApp.SdfsfsfdApiGroup.Ceshi2Api
	{
		ceshi2Router.POST("createCeshi2", ceshi2Api.CreateCeshi2)   // 新建Ceshi2
		ceshi2Router.DELETE("deleteCeshi2", ceshi2Api.DeleteCeshi2) // 删除Ceshi2
		ceshi2Router.DELETE("deleteCeshi2ByIds", ceshi2Api.DeleteCeshi2ByIds) // 批量删除Ceshi2
		ceshi2Router.PUT("updateCeshi2", ceshi2Api.UpdateCeshi2)    // 更新Ceshi2
	}
	{
		ceshi2RouterWithoutRecord.GET("findCeshi2", ceshi2Api.FindCeshi2)        // 根据ID获取Ceshi2
		ceshi2RouterWithoutRecord.GET("getCeshi2List", ceshi2Api.GetCeshi2List)  // 获取Ceshi2列表
	}
}
