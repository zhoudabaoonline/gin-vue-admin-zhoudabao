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

type DfdfdfApi struct {
}

var dfdfdfService = service.ServiceGroupApp.SdfsfsfdServiceGroup.DfdfdfService


// CreateDfdfdf 创建Dfdfdf
// @Tags Dfdfdf
// @Summary 创建Dfdfdf
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sdfsfsfd.Dfdfdf true "创建Dfdfdf"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dfdfdf/createDfdfdf [post]
func (dfdfdfApi *DfdfdfApi) CreateDfdfdf(c *gin.Context) {
	var dfdfdf sdfsfsfd.Dfdfdf
	err := c.ShouldBindJSON(&dfdfdf)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := dfdfdfService.CreateDfdfdf(dfdfdf); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteDfdfdf 删除Dfdfdf
// @Tags Dfdfdf
// @Summary 删除Dfdfdf
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sdfsfsfd.Dfdfdf true "删除Dfdfdf"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dfdfdf/deleteDfdfdf [delete]
func (dfdfdfApi *DfdfdfApi) DeleteDfdfdf(c *gin.Context) {
	var dfdfdf sdfsfsfd.Dfdfdf
	err := c.ShouldBindJSON(&dfdfdf)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := dfdfdfService.DeleteDfdfdf(dfdfdf); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDfdfdfByIds 批量删除Dfdfdf
// @Tags Dfdfdf
// @Summary 批量删除Dfdfdf
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Dfdfdf"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /dfdfdf/deleteDfdfdfByIds [delete]
func (dfdfdfApi *DfdfdfApi) DeleteDfdfdfByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := dfdfdfService.DeleteDfdfdfByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDfdfdf 更新Dfdfdf
// @Tags Dfdfdf
// @Summary 更新Dfdfdf
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sdfsfsfd.Dfdfdf true "更新Dfdfdf"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dfdfdf/updateDfdfdf [put]
func (dfdfdfApi *DfdfdfApi) UpdateDfdfdf(c *gin.Context) {
	var dfdfdf sdfsfsfd.Dfdfdf
	err := c.ShouldBindJSON(&dfdfdf)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := dfdfdfService.UpdateDfdfdf(dfdfdf); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindDfdfdf 用id查询Dfdfdf
// @Tags Dfdfdf
// @Summary 用id查询Dfdfdf
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query sdfsfsfd.Dfdfdf true "用id查询Dfdfdf"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dfdfdf/findDfdfdf [get]
func (dfdfdfApi *DfdfdfApi) FindDfdfdf(c *gin.Context) {
	var dfdfdf sdfsfsfd.Dfdfdf
	err := c.ShouldBindQuery(&dfdfdf)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if redfdfdf, err := dfdfdfService.GetDfdfdf(dfdfdf.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redfdfdf": redfdfdf}, c)
	}
}

// GetDfdfdfList 分页获取Dfdfdf列表
// @Tags Dfdfdf
// @Summary 分页获取Dfdfdf列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query sdfsfsfdReq.DfdfdfSearch true "分页获取Dfdfdf列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dfdfdf/getDfdfdfList [get]
func (dfdfdfApi *DfdfdfApi) GetDfdfdfList(c *gin.Context) {
	var pageInfo sdfsfsfdReq.DfdfdfSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := dfdfdfService.GetDfdfdfInfoList(pageInfo); err != nil {
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
