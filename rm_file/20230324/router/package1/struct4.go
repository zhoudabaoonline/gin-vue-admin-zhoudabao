package package1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type Struct4Router struct {
}

// InitStruct4Router 初始化 Struct4 路由信息
func (s *Struct4Router) InitStruct4Router(Router *gin.RouterGroup) {
	struct4Router := Router.Group("struct4").Use(middleware.OperationRecord())
	struct4RouterWithoutRecord := Router.Group("struct4")
	var struct4Api = v1.ApiGroupApp.Package1ApiGroup.Struct4Api
	{
		struct4Router.POST("createStruct4", struct4Api.CreateStruct4)   // 新建Struct4
		struct4Router.DELETE("deleteStruct4", struct4Api.DeleteStruct4) // 删除Struct4
		struct4Router.DELETE("deleteStruct4ByIds", struct4Api.DeleteStruct4ByIds) // 批量删除Struct4
		struct4Router.PUT("updateStruct4", struct4Api.UpdateStruct4)    // 更新Struct4
	}
	{
		struct4RouterWithoutRecord.GET("findStruct4", struct4Api.FindStruct4)        // 根据ID获取Struct4
		struct4RouterWithoutRecord.GET("getStruct4List", struct4Api.GetStruct4List)  // 获取Struct4列表
	}
}
