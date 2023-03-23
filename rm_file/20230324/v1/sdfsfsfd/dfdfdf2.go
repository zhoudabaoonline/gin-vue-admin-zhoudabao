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

type Dfdfdf22Api struct {
}

var dfdfdf22Service = service.ServiceGroupApp.SdfsfsfdServiceGroup.Dfdfdf22Service


// CreateDfdfdf22 创建Dfdfdf22
// @Tags Dfdfdf22
// @Summary 创建Dfdfdf22
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sdfsfsfd.Dfdfdf22 true "创建Dfdfdf22"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dfdfdf22/createDfdfdf22 [post]
func (dfdfdf22Api *Dfdfdf22Api) CreateDfdfdf22(c *gin.Context) {
	var dfdfdf22 sdfsfsfd.Dfdfdf22
	err := c.ShouldBindJSON(&dfdfdf22)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := dfdfdf22Service.CreateDfdfdf22(dfdfdf22); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteDfdfdf22 删除Dfdfdf22
// @Tags Dfdfdf22
// @Summary 删除Dfdfdf22
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sdfsfsfd.Dfdfdf22 true "删除Dfdfdf22"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dfdfdf22/deleteDfdfdf22 [delete]
func (dfdfdf22Api *Dfdfdf22Api) DeleteDfdfdf22(c *gin.Context) {
	var dfdfdf22 sdfsfsfd.Dfdfdf22
	err := c.ShouldBindJSON(&dfdfdf22)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := dfdfdf22Service.DeleteDfdfdf22(dfdfdf22); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDfdfdf22ByIds 批量删除Dfdfdf22
// @Tags Dfdfdf22
// @Summary 批量删除Dfdfdf22
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Dfdfdf22"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /dfdfdf22/deleteDfdfdf22ByIds [delete]
func (dfdfdf22Api *Dfdfdf22Api) DeleteDfdfdf22ByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := dfdfdf22Service.DeleteDfdfdf22ByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDfdfdf22 更新Dfdfdf22
// @Tags Dfdfdf22
// @Summary 更新Dfdfdf22
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sdfsfsfd.Dfdfdf22 true "更新Dfdfdf22"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dfdfdf22/updateDfdfdf22 [put]
func (dfdfdf22Api *Dfdfdf22Api) UpdateDfdfdf22(c *gin.Context) {
	var dfdfdf22 sdfsfsfd.Dfdfdf22
	err := c.ShouldBindJSON(&dfdfdf22)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := dfdfdf22Service.UpdateDfdfdf22(dfdfdf22); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindDfdfdf22 用id查询Dfdfdf22
// @Tags Dfdfdf22
// @Summary 用id查询Dfdfdf22
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query sdfsfsfd.Dfdfdf22 true "用id查询Dfdfdf22"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dfdfdf22/findDfdfdf22 [get]
func (dfdfdf22Api *Dfdfdf22Api) FindDfdfdf22(c *gin.Context) {
	var dfdfdf22 sdfsfsfd.Dfdfdf22
	err := c.ShouldBindQuery(&dfdfdf22)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if redfdfdf22, err := dfdfdf22Service.GetDfdfdf22(dfdfdf22.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redfdfdf22": redfdfdf22}, c)
	}
}

// GetDfdfdf22List 分页获取Dfdfdf22列表
// @Tags Dfdfdf22
// @Summary 分页获取Dfdfdf22列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query sdfsfsfdReq.Dfdfdf22Search true "分页获取Dfdfdf22列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dfdfdf22/getDfdfdf22List [get]
func (dfdfdf22Api *Dfdfdf22Api) GetDfdfdf22List(c *gin.Context) {
	var pageInfo sdfsfsfdReq.Dfdfdf22Search
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := dfdfdf22Service.GetDfdfdf22InfoList(pageInfo); err != nil {
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
