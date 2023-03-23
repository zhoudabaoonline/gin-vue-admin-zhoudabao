package ceshiplug

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/ceshiplug/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/ceshiplug/router"
	"github.com/gin-gonic/gin"
)

type CeshiPlugPlugin struct {
}

func CreateCeshiPlugPlug(globalProperty string) *CeshiPlugPlugin {
	global.GlobalConfig.GlobalProperty = globalProperty

	return &CeshiPlugPlugin{}
}

func (*CeshiPlugPlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitCeshiPlugRouter(group)
}

func (*CeshiPlugPlugin) RouterPath() string {
	return "/routeGroup"
}
