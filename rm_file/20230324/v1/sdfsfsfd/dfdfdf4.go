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

type Dfdfdf4Api struct {
}

var dfdfdf4Service = service.ServiceGroupApp.SdfsfsfdServiceGroup.Dfdfdf4Service


// CreateDfdfdf4 创建Dfdfdf4
// @Tags Dfdfdf4
// @Summary 创建Dfdfdf4
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sdfsfsfd.Dfdfdf4 true "创建Dfdfdf4"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dfdfdf4/createDfdfdf4 [post]
func (dfdfdf4Api *Dfdfdf4Api) CreateDfdfdf4(c *gin.Context) {
	var dfdfdf4 sdfsfsfd.Dfdfdf4
	err := c.ShouldBindJSON(&dfdfdf4)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := dfdfdf4Service.CreateDfdfdf4(dfdfdf4); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteDfdfdf4 删除Dfdfdf4
// @Tags Dfdfdf4
// @Summary 删除Dfdfdf4
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sdfsfsfd.Dfdfdf4 true "删除Dfdfdf4"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dfdfdf4/deleteDfdfdf4 [delete]
func (dfdfdf4Api *Dfdfdf4Api) DeleteDfdfdf4(c *gin.Context) {
	var dfdfdf4 sdfsfsfd.Dfdfdf4
	err := c.ShouldBindJSON(&dfdfdf4)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := dfdfdf4Service.DeleteDfdfdf4(dfdfdf4); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDfdfdf4ByIds 批量删除Dfdfdf4
// @Tags Dfdfdf4
// @Summary 批量删除Dfdfdf4
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Dfdfdf4"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /dfdfdf4/deleteDfdfdf4ByIds [delete]
func (dfdfdf4Api *Dfdfdf4Api) DeleteDfdfdf4ByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := dfdfdf4Service.DeleteDfdfdf4ByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDfdfdf4 更新Dfdfdf4
// @Tags Dfdfdf4
// @Summary 更新Dfdfdf4
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sdfsfsfd.Dfdfdf4 true "更新Dfdfdf4"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dfdfdf4/updateDfdfdf4 [put]
func (dfdfdf4Api *Dfdfdf4Api) UpdateDfdfdf4(c *gin.Context) {
	var dfdfdf4 sdfsfsfd.Dfdfdf4
	err := c.ShouldBindJSON(&dfdfdf4)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := dfdfdf4Service.UpdateDfdfdf4(dfdfdf4); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindDfdfdf4 用id查询Dfdfdf4
// @Tags Dfdfdf4
// @Summary 用id查询Dfdfdf4
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query sdfsfsfd.Dfdfdf4 true "用id查询Dfdfdf4"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dfdfdf4/findDfdfdf4 [get]
func (dfdfdf4Api *Dfdfdf4Api) FindDfdfdf4(c *gin.Context) {
	var dfdfdf4 sdfsfsfd.Dfdfdf4
	err := c.ShouldBindQuery(&dfdfdf4)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if redfdfdf4, err := dfdfdf4Service.GetDfdfdf4(dfdfdf4.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redfdfdf4": redfdfdf4}, c)
	}
}

// GetDfdfdf4List 分页获取Dfdfdf4列表
// @Tags Dfdfdf4
// @Summary 分页获取Dfdfdf4列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query sdfsfsfdReq.Dfdfdf4Search true "分页获取Dfdfdf4列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dfdfdf4/getDfdfdf4List [get]
func (dfdfdf4Api *Dfdfdf4Api) GetDfdfdf4List(c *gin.Context) {
	var pageInfo sdfsfsfdReq.Dfdfdf4Search
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := dfdfdf4Service.GetDfdfdf4InfoList(pageInfo); err != nil {
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
