package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/ceshiplug/api"
	"github.com/gin-gonic/gin"
)

type CeshiPlugRouter struct {
}

func (s *CeshiPlugRouter) InitCeshiPlugRouter(Router *gin.RouterGroup) {
	plugRouter := Router
	plugApi := api.ApiGroupApp.CeshiPlugApi
	{
		plugRouter.POST("routerName", plugApi.ApiName)
	}
}
