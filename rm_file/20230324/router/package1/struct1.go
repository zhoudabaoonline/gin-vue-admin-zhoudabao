package package1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type Struct1Router struct {
}

// InitStruct1Router 初始化 Struct1 路由信息
func (s *Struct1Router) InitStruct1Router(Router *gin.RouterGroup) {
	struct1Router := Router.Group("struct1").Use(middleware.OperationRecord())
	struct1RouterWithoutRecord := Router.Group("struct1")
	var struct1Api = v1.ApiGroupApp.Package1ApiGroup.Struct1Api
	{
		struct1Router.POST("createStruct1", struct1Api.CreateStruct1)   // 新建Struct1
		struct1Router.DELETE("deleteStruct1", struct1Api.DeleteStruct1) // 删除Struct1
		struct1Router.DELETE("deleteStruct1ByIds", struct1Api.DeleteStruct1ByIds) // 批量删除Struct1
		struct1Router.PUT("updateStruct1", struct1Api.UpdateStruct1)    // 更新Struct1
	}
	{
		struct1RouterWithoutRecord.GET("findStruct1", struct1Api.FindStruct1)        // 根据ID获取Struct1
		struct1RouterWithoutRecord.GET("getStruct1List", struct1Api.GetStruct1List)  // 获取Struct1列表
	}
}
