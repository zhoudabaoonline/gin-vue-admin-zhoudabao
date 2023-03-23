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

type DfdfdfvvvvApi struct {
}

var dfdfdfvvService = service.ServiceGroupApp.SdfsfsfdServiceGroup.DfdfdfvvvvService


// CreateDfdfdfvvvv 创建Dfdfdfvvvv
// @Tags Dfdfdfvvvv
// @Summary 创建Dfdfdfvvvv
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sdfsfsfd.Dfdfdfvvvv true "创建Dfdfdfvvvv"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dfdfdfvv/createDfdfdfvvvv [post]
func (dfdfdfvvApi *DfdfdfvvvvApi) CreateDfdfdfvvvv(c *gin.Context) {
	var dfdfdfvv sdfsfsfd.Dfdfdfvvvv
	err := c.ShouldBindJSON(&dfdfdfvv)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := dfdfdfvvService.CreateDfdfdfvvvv(dfdfdfvv); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteDfdfdfvvvv 删除Dfdfdfvvvv
// @Tags Dfdfdfvvvv
// @Summary 删除Dfdfdfvvvv
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sdfsfsfd.Dfdfdfvvvv true "删除Dfdfdfvvvv"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dfdfdfvv/deleteDfdfdfvvvv [delete]
func (dfdfdfvvApi *DfdfdfvvvvApi) DeleteDfdfdfvvvv(c *gin.Context) {
	var dfdfdfvv sdfsfsfd.Dfdfdfvvvv
	err := c.ShouldBindJSON(&dfdfdfvv)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := dfdfdfvvService.DeleteDfdfdfvvvv(dfdfdfvv); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDfdfdfvvvvByIds 批量删除Dfdfdfvvvv
// @Tags Dfdfdfvvvv
// @Summary 批量删除Dfdfdfvvvv
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Dfdfdfvvvv"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /dfdfdfvv/deleteDfdfdfvvvvByIds [delete]
func (dfdfdfvvApi *DfdfdfvvvvApi) DeleteDfdfdfvvvvByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := dfdfdfvvService.DeleteDfdfdfvvvvByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDfdfdfvvvv 更新Dfdfdfvvvv
// @Tags Dfdfdfvvvv
// @Summary 更新Dfdfdfvvvv
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sdfsfsfd.Dfdfdfvvvv true "更新Dfdfdfvvvv"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dfdfdfvv/updateDfdfdfvvvv [put]
func (dfdfdfvvApi *DfdfdfvvvvApi) UpdateDfdfdfvvvv(c *gin.Context) {
	var dfdfdfvv sdfsfsfd.Dfdfdfvvvv
	err := c.ShouldBindJSON(&dfdfdfvv)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := dfdfdfvvService.UpdateDfdfdfvvvv(dfdfdfvv); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindDfdfdfvvvv 用id查询Dfdfdfvvvv
// @Tags Dfdfdfvvvv
// @Summary 用id查询Dfdfdfvvvv
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query sdfsfsfd.Dfdfdfvvvv true "用id查询Dfdfdfvvvv"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dfdfdfvv/findDfdfdfvvvv [get]
func (dfdfdfvvApi *DfdfdfvvvvApi) FindDfdfdfvvvv(c *gin.Context) {
	var dfdfdfvv sdfsfsfd.Dfdfdfvvvv
	err := c.ShouldBindQuery(&dfdfdfvv)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if redfdfdfvv, err := dfdfdfvvService.GetDfdfdfvvvv(dfdfdfvv.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redfdfdfvv": redfdfdfvv}, c)
	}
}

// GetDfdfdfvvvvList 分页获取Dfdfdfvvvv列表
// @Tags Dfdfdfvvvv
// @Summary 分页获取Dfdfdfvvvv列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query sdfsfsfdReq.DfdfdfvvvvSearch true "分页获取Dfdfdfvvvv列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dfdfdfvv/getDfdfdfvvvvList [get]
func (dfdfdfvvApi *DfdfdfvvvvApi) GetDfdfdfvvvvList(c *gin.Context) {
	var pageInfo sdfsfsfdReq.DfdfdfvvvvSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := dfdfdfvvService.GetDfdfdfvvvvInfoList(pageInfo); err != nil {
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
