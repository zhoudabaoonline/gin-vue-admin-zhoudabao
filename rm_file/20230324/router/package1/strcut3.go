package package1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type Strcut3Router struct {
}

// InitStrcut3Router 初始化 Strcut3 路由信息
func (s *Strcut3Router) InitStrcut3Router(Router *gin.RouterGroup) {
	strcut3Router := Router.Group("strcut3").Use(middleware.OperationRecord())
	strcut3RouterWithoutRecord := Router.Group("strcut3")
	var strcut3Api = v1.ApiGroupApp.Package1ApiGroup.Strcut3Api
	{
		strcut3Router.POST("createStrcut3", strcut3Api.CreateStrcut3)   // 新建Strcut3
		strcut3Router.DELETE("deleteStrcut3", strcut3Api.DeleteStrcut3) // 删除Strcut3
		strcut3Router.DELETE("deleteStrcut3ByIds", strcut3Api.DeleteStrcut3ByIds) // 批量删除Strcut3
		strcut3Router.PUT("updateStrcut3", strcut3Api.UpdateStrcut3)    // 更新Strcut3
	}
	{
		strcut3RouterWithoutRecord.GET("findStrcut3", strcut3Api.FindStrcut3)        // 根据ID获取Strcut3
		strcut3RouterWithoutRecord.GET("getStrcut3List", strcut3Api.GetStrcut3List)  // 获取Strcut3列表
	}
}
