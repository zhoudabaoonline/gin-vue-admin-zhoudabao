package package1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type Struct2Router struct {
}

// InitStruct2Router 初始化 Struct2 路由信息
func (s *Struct2Router) InitStruct2Router(Router *gin.RouterGroup) {
	struct2Router := Router.Group("struct2").Use(middleware.OperationRecord())
	struct2RouterWithoutRecord := Router.Group("struct2")
	var struct2Api = v1.ApiGroupApp.Package1ApiGroup.Struct2Api
	{
		struct2Router.POST("createStruct2", struct2Api.CreateStruct2)   // 新建Struct2
		struct2Router.DELETE("deleteStruct2", struct2Api.DeleteStruct2) // 删除Struct2
		struct2Router.DELETE("deleteStruct2ByIds", struct2Api.DeleteStruct2ByIds) // 批量删除Struct2
		struct2Router.PUT("updateStruct2", struct2Api.UpdateStruct2)    // 更新Struct2
	}
	{
		struct2RouterWithoutRecord.GET("findStruct2", struct2Api.FindStruct2)        // 根据ID获取Struct2
		struct2RouterWithoutRecord.GET("getStruct2List", struct2Api.GetStruct2List)  // 获取Struct2列表
	}
}
