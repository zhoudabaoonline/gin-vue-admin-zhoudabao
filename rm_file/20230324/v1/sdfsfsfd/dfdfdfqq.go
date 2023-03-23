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

type DfdfdfqqqApi struct {
}

var dfdfdfqqqService = service.ServiceGroupApp.SdfsfsfdServiceGroup.DfdfdfqqqService


// CreateDfdfdfqqq 创建Dfdfdfqqq
// @Tags Dfdfdfqqq
// @Summary 创建Dfdfdfqqq
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sdfsfsfd.Dfdfdfqqq true "创建Dfdfdfqqq"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dfdfdfqqq/createDfdfdfqqq [post]
func (dfdfdfqqqApi *DfdfdfqqqApi) CreateDfdfdfqqq(c *gin.Context) {
	var dfdfdfqqq sdfsfsfd.Dfdfdfqqq
	err := c.ShouldBindJSON(&dfdfdfqqq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := dfdfdfqqqService.CreateDfdfdfqqq(dfdfdfqqq); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteDfdfdfqqq 删除Dfdfdfqqq
// @Tags Dfdfdfqqq
// @Summary 删除Dfdfdfqqq
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sdfsfsfd.Dfdfdfqqq true "删除Dfdfdfqqq"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dfdfdfqqq/deleteDfdfdfqqq [delete]
func (dfdfdfqqqApi *DfdfdfqqqApi) DeleteDfdfdfqqq(c *gin.Context) {
	var dfdfdfqqq sdfsfsfd.Dfdfdfqqq
	err := c.ShouldBindJSON(&dfdfdfqqq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := dfdfdfqqqService.DeleteDfdfdfqqq(dfdfdfqqq); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDfdfdfqqqByIds 批量删除Dfdfdfqqq
// @Tags Dfdfdfqqq
// @Summary 批量删除Dfdfdfqqq
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Dfdfdfqqq"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /dfdfdfqqq/deleteDfdfdfqqqByIds [delete]
func (dfdfdfqqqApi *DfdfdfqqqApi) DeleteDfdfdfqqqByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := dfdfdfqqqService.DeleteDfdfdfqqqByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDfdfdfqqq 更新Dfdfdfqqq
// @Tags Dfdfdfqqq
// @Summary 更新Dfdfdfqqq
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sdfsfsfd.Dfdfdfqqq true "更新Dfdfdfqqq"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dfdfdfqqq/updateDfdfdfqqq [put]
func (dfdfdfqqqApi *DfdfdfqqqApi) UpdateDfdfdfqqq(c *gin.Context) {
	var dfdfdfqqq sdfsfsfd.Dfdfdfqqq
	err := c.ShouldBindJSON(&dfdfdfqqq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := dfdfdfqqqService.UpdateDfdfdfqqq(dfdfdfqqq); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindDfdfdfqqq 用id查询Dfdfdfqqq
// @Tags Dfdfdfqqq
// @Summary 用id查询Dfdfdfqqq
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query sdfsfsfd.Dfdfdfqqq true "用id查询Dfdfdfqqq"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dfdfdfqqq/findDfdfdfqqq [get]
func (dfdfdfqqqApi *DfdfdfqqqApi) FindDfdfdfqqq(c *gin.Context) {
	var dfdfdfqqq sdfsfsfd.Dfdfdfqqq
	err := c.ShouldBindQuery(&dfdfdfqqq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if redfdfdfqqq, err := dfdfdfqqqService.GetDfdfdfqqq(dfdfdfqqq.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redfdfdfqqq": redfdfdfqqq}, c)
	}
}

// GetDfdfdfqqqList 分页获取Dfdfdfqqq列表
// @Tags Dfdfdfqqq
// @Summary 分页获取Dfdfdfqqq列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query sdfsfsfdReq.DfdfdfqqqSearch true "分页获取Dfdfdfqqq列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dfdfdfqqq/getDfdfdfqqqList [get]
func (dfdfdfqqqApi *DfdfdfqqqApi) GetDfdfdfqqqList(c *gin.Context) {
	var pageInfo sdfsfsfdReq.DfdfdfqqqSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := dfdfdfqqqService.GetDfdfdfqqqInfoList(pageInfo); err != nil {
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
