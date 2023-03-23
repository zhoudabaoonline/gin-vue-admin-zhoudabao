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

type DfdfdfzzApi struct {
}

var dfdfdfzzzzService = service.ServiceGroupApp.SdfsfsfdServiceGroup.DfdfdfzzService


// CreateDfdfdfzz 创建Dfdfdfzz
// @Tags Dfdfdfzz
// @Summary 创建Dfdfdfzz
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sdfsfsfd.Dfdfdfzz true "创建Dfdfdfzz"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dfdfdfzzzz/createDfdfdfzz [post]
func (dfdfdfzzzzApi *DfdfdfzzApi) CreateDfdfdfzz(c *gin.Context) {
	var dfdfdfzzzz sdfsfsfd.Dfdfdfzz
	err := c.ShouldBindJSON(&dfdfdfzzzz)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := dfdfdfzzzzService.CreateDfdfdfzz(dfdfdfzzzz); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteDfdfdfzz 删除Dfdfdfzz
// @Tags Dfdfdfzz
// @Summary 删除Dfdfdfzz
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sdfsfsfd.Dfdfdfzz true "删除Dfdfdfzz"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dfdfdfzzzz/deleteDfdfdfzz [delete]
func (dfdfdfzzzzApi *DfdfdfzzApi) DeleteDfdfdfzz(c *gin.Context) {
	var dfdfdfzzzz sdfsfsfd.Dfdfdfzz
	err := c.ShouldBindJSON(&dfdfdfzzzz)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := dfdfdfzzzzService.DeleteDfdfdfzz(dfdfdfzzzz); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDfdfdfzzByIds 批量删除Dfdfdfzz
// @Tags Dfdfdfzz
// @Summary 批量删除Dfdfdfzz
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Dfdfdfzz"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /dfdfdfzzzz/deleteDfdfdfzzByIds [delete]
func (dfdfdfzzzzApi *DfdfdfzzApi) DeleteDfdfdfzzByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := dfdfdfzzzzService.DeleteDfdfdfzzByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDfdfdfzz 更新Dfdfdfzz
// @Tags Dfdfdfzz
// @Summary 更新Dfdfdfzz
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sdfsfsfd.Dfdfdfzz true "更新Dfdfdfzz"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dfdfdfzzzz/updateDfdfdfzz [put]
func (dfdfdfzzzzApi *DfdfdfzzApi) UpdateDfdfdfzz(c *gin.Context) {
	var dfdfdfzzzz sdfsfsfd.Dfdfdfzz
	err := c.ShouldBindJSON(&dfdfdfzzzz)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := dfdfdfzzzzService.UpdateDfdfdfzz(dfdfdfzzzz); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindDfdfdfzz 用id查询Dfdfdfzz
// @Tags Dfdfdfzz
// @Summary 用id查询Dfdfdfzz
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query sdfsfsfd.Dfdfdfzz true "用id查询Dfdfdfzz"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dfdfdfzzzz/findDfdfdfzz [get]
func (dfdfdfzzzzApi *DfdfdfzzApi) FindDfdfdfzz(c *gin.Context) {
	var dfdfdfzzzz sdfsfsfd.Dfdfdfzz
	err := c.ShouldBindQuery(&dfdfdfzzzz)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if redfdfdfzzzz, err := dfdfdfzzzzService.GetDfdfdfzz(dfdfdfzzzz.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redfdfdfzzzz": redfdfdfzzzz}, c)
	}
}

// GetDfdfdfzzList 分页获取Dfdfdfzz列表
// @Tags Dfdfdfzz
// @Summary 分页获取Dfdfdfzz列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query sdfsfsfdReq.DfdfdfzzSearch true "分页获取Dfdfdfzz列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dfdfdfzzzz/getDfdfdfzzList [get]
func (dfdfdfzzzzApi *DfdfdfzzApi) GetDfdfdfzzList(c *gin.Context) {
	var pageInfo sdfsfsfdReq.DfdfdfzzSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := dfdfdfzzzzService.GetDfdfdfzzInfoList(pageInfo); err != nil {
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
