package sdfsfsfd

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/sdfsfsfd"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    sdfsfsfdReq "github.com/flipped-aurora/gin-vue-admin/server/model/sdfsfsfd/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type CeshiApi struct {
}

var ceshiService = service.ServiceGroupApp.SdfsfsfdServiceGroup.CeshiService


// CreateCeshi 创建Ceshi
// @Tags Ceshi
// @Summary 创建Ceshi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sdfsfsfd.Ceshi true "创建Ceshi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ceshi/createCeshi [post]
func (ceshiApi *CeshiApi) CreateCeshi(c *gin.Context) {
	var ceshi sdfsfsfd.Ceshi
	err := c.ShouldBindJSON(&ceshi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ceshiService.CreateCeshi(ceshi); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCeshi 删除Ceshi
// @Tags Ceshi
// @Summary 删除Ceshi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sdfsfsfd.Ceshi true "删除Ceshi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ceshi/deleteCeshi [delete]
func (ceshiApi *CeshiApi) DeleteCeshi(c *gin.Context) {
	var ceshi sdfsfsfd.Ceshi
	err := c.ShouldBindJSON(&ceshi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ceshiService.DeleteCeshi(ceshi); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCeshiByIds 批量删除Ceshi
// @Tags Ceshi
// @Summary 批量删除Ceshi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Ceshi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /ceshi/deleteCeshiByIds [delete]
func (ceshiApi *CeshiApi) DeleteCeshiByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ceshiService.DeleteCeshiByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCeshi 更新Ceshi
// @Tags Ceshi
// @Summary 更新Ceshi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sdfsfsfd.Ceshi true "更新Ceshi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ceshi/updateCeshi [put]
func (ceshiApi *CeshiApi) UpdateCeshi(c *gin.Context) {
	var ceshi sdfsfsfd.Ceshi
	err := c.ShouldBindJSON(&ceshi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ceshiService.UpdateCeshi(ceshi); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCeshi 用id查询Ceshi
// @Tags Ceshi
// @Summary 用id查询Ceshi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query sdfsfsfd.Ceshi true "用id查询Ceshi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ceshi/findCeshi [get]
func (ceshiApi *CeshiApi) FindCeshi(c *gin.Context) {
	var ceshi sdfsfsfd.Ceshi
	err := c.ShouldBindQuery(&ceshi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if receshi, err := ceshiService.GetCeshi(ceshi.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"receshi": receshi}, c)
	}
}

// GetCeshiList 分页获取Ceshi列表
// @Tags Ceshi
// @Summary 分页获取Ceshi列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query sdfsfsfdReq.CeshiSearch true "分页获取Ceshi列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ceshi/getCeshiList [get]
func (ceshiApi *CeshiApi) GetCeshiList(c *gin.Context) {
	var pageInfo sdfsfsfdReq.CeshiSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := ceshiService.GetCeshiInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
