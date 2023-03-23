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

type DfdfdfvvvApi struct {
}

var dfdfdfvvvvService = service.ServiceGroupApp.SdfsfsfdServiceGroup.DfdfdfvvvService


// CreateDfdfdfvvv 创建Dfdfdfvvv
// @Tags Dfdfdfvvv
// @Summary 创建Dfdfdfvvv
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sdfsfsfd.Dfdfdfvvv true "创建Dfdfdfvvv"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dfdfdfvvvv/createDfdfdfvvv [post]
func (dfdfdfvvvvApi *DfdfdfvvvApi) CreateDfdfdfvvv(c *gin.Context) {
	var dfdfdfvvvv sdfsfsfd.Dfdfdfvvv
	err := c.ShouldBindJSON(&dfdfdfvvvv)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := dfdfdfvvvvService.CreateDfdfdfvvv(dfdfdfvvvv); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteDfdfdfvvv 删除Dfdfdfvvv
// @Tags Dfdfdfvvv
// @Summary 删除Dfdfdfvvv
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sdfsfsfd.Dfdfdfvvv true "删除Dfdfdfvvv"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dfdfdfvvvv/deleteDfdfdfvvv [delete]
func (dfdfdfvvvvApi *DfdfdfvvvApi) DeleteDfdfdfvvv(c *gin.Context) {
	var dfdfdfvvvv sdfsfsfd.Dfdfdfvvv
	err := c.ShouldBindJSON(&dfdfdfvvvv)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := dfdfdfvvvvService.DeleteDfdfdfvvv(dfdfdfvvvv); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDfdfdfvvvByIds 批量删除Dfdfdfvvv
// @Tags Dfdfdfvvv
// @Summary 批量删除Dfdfdfvvv
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Dfdfdfvvv"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /dfdfdfvvvv/deleteDfdfdfvvvByIds [delete]
func (dfdfdfvvvvApi *DfdfdfvvvApi) DeleteDfdfdfvvvByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := dfdfdfvvvvService.DeleteDfdfdfvvvByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDfdfdfvvv 更新Dfdfdfvvv
// @Tags Dfdfdfvvv
// @Summary 更新Dfdfdfvvv
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sdfsfsfd.Dfdfdfvvv true "更新Dfdfdfvvv"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dfdfdfvvvv/updateDfdfdfvvv [put]
func (dfdfdfvvvvApi *DfdfdfvvvApi) UpdateDfdfdfvvv(c *gin.Context) {
	var dfdfdfvvvv sdfsfsfd.Dfdfdfvvv
	err := c.ShouldBindJSON(&dfdfdfvvvv)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := dfdfdfvvvvService.UpdateDfdfdfvvv(dfdfdfvvvv); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindDfdfdfvvv 用id查询Dfdfdfvvv
// @Tags Dfdfdfvvv
// @Summary 用id查询Dfdfdfvvv
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query sdfsfsfd.Dfdfdfvvv true "用id查询Dfdfdfvvv"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dfdfdfvvvv/findDfdfdfvvv [get]
func (dfdfdfvvvvApi *DfdfdfvvvApi) FindDfdfdfvvv(c *gin.Context) {
	var dfdfdfvvvv sdfsfsfd.Dfdfdfvvv
	err := c.ShouldBindQuery(&dfdfdfvvvv)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if redfdfdfvvvv, err := dfdfdfvvvvService.GetDfdfdfvvv(dfdfdfvvvv.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redfdfdfvvvv": redfdfdfvvvv}, c)
	}
}

// GetDfdfdfvvvList 分页获取Dfdfdfvvv列表
// @Tags Dfdfdfvvv
// @Summary 分页获取Dfdfdfvvv列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query sdfsfsfdReq.DfdfdfvvvSearch true "分页获取Dfdfdfvvv列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dfdfdfvvvv/getDfdfdfvvvList [get]
func (dfdfdfvvvvApi *DfdfdfvvvApi) GetDfdfdfvvvList(c *gin.Context) {
	var pageInfo sdfsfsfdReq.DfdfdfvvvSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := dfdfdfvvvvService.GetDfdfdfvvvInfoList(pageInfo); err != nil {
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
