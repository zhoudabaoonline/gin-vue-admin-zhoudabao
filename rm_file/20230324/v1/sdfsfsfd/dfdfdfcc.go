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

type DfdfdfccApi struct {
}

var dfdfdfccService = service.ServiceGroupApp.SdfsfsfdServiceGroup.DfdfdfccService


// CreateDfdfdfcc 创建Dfdfdfcc
// @Tags Dfdfdfcc
// @Summary 创建Dfdfdfcc
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sdfsfsfd.Dfdfdfcc true "创建Dfdfdfcc"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dfdfdfcc/createDfdfdfcc [post]
func (dfdfdfccApi *DfdfdfccApi) CreateDfdfdfcc(c *gin.Context) {
	var dfdfdfcc sdfsfsfd.Dfdfdfcc
	err := c.ShouldBindJSON(&dfdfdfcc)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := dfdfdfccService.CreateDfdfdfcc(dfdfdfcc); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteDfdfdfcc 删除Dfdfdfcc
// @Tags Dfdfdfcc
// @Summary 删除Dfdfdfcc
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sdfsfsfd.Dfdfdfcc true "删除Dfdfdfcc"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dfdfdfcc/deleteDfdfdfcc [delete]
func (dfdfdfccApi *DfdfdfccApi) DeleteDfdfdfcc(c *gin.Context) {
	var dfdfdfcc sdfsfsfd.Dfdfdfcc
	err := c.ShouldBindJSON(&dfdfdfcc)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := dfdfdfccService.DeleteDfdfdfcc(dfdfdfcc); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDfdfdfccByIds 批量删除Dfdfdfcc
// @Tags Dfdfdfcc
// @Summary 批量删除Dfdfdfcc
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Dfdfdfcc"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /dfdfdfcc/deleteDfdfdfccByIds [delete]
func (dfdfdfccApi *DfdfdfccApi) DeleteDfdfdfccByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := dfdfdfccService.DeleteDfdfdfccByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDfdfdfcc 更新Dfdfdfcc
// @Tags Dfdfdfcc
// @Summary 更新Dfdfdfcc
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sdfsfsfd.Dfdfdfcc true "更新Dfdfdfcc"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dfdfdfcc/updateDfdfdfcc [put]
func (dfdfdfccApi *DfdfdfccApi) UpdateDfdfdfcc(c *gin.Context) {
	var dfdfdfcc sdfsfsfd.Dfdfdfcc
	err := c.ShouldBindJSON(&dfdfdfcc)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := dfdfdfccService.UpdateDfdfdfcc(dfdfdfcc); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindDfdfdfcc 用id查询Dfdfdfcc
// @Tags Dfdfdfcc
// @Summary 用id查询Dfdfdfcc
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query sdfsfsfd.Dfdfdfcc true "用id查询Dfdfdfcc"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dfdfdfcc/findDfdfdfcc [get]
func (dfdfdfccApi *DfdfdfccApi) FindDfdfdfcc(c *gin.Context) {
	var dfdfdfcc sdfsfsfd.Dfdfdfcc
	err := c.ShouldBindQuery(&dfdfdfcc)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if redfdfdfcc, err := dfdfdfccService.GetDfdfdfcc(dfdfdfcc.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redfdfdfcc": redfdfdfcc}, c)
	}
}

// GetDfdfdfccList 分页获取Dfdfdfcc列表
// @Tags Dfdfdfcc
// @Summary 分页获取Dfdfdfcc列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query sdfsfsfdReq.DfdfdfccSearch true "分页获取Dfdfdfcc列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dfdfdfcc/getDfdfdfccList [get]
func (dfdfdfccApi *DfdfdfccApi) GetDfdfdfccList(c *gin.Context) {
	var pageInfo sdfsfsfdReq.DfdfdfccSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := dfdfdfccService.GetDfdfdfccInfoList(pageInfo); err != nil {
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
