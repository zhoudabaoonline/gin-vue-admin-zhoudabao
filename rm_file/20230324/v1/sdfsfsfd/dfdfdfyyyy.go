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

type DfdfdfyyyyApi struct {
}

var dfdfdfyyyService = service.ServiceGroupApp.SdfsfsfdServiceGroup.DfdfdfyyyyService


// CreateDfdfdfyyyy 创建Dfdfdfyyyy
// @Tags Dfdfdfyyyy
// @Summary 创建Dfdfdfyyyy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sdfsfsfd.Dfdfdfyyyy true "创建Dfdfdfyyyy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dfdfdfyyy/createDfdfdfyyyy [post]
func (dfdfdfyyyApi *DfdfdfyyyyApi) CreateDfdfdfyyyy(c *gin.Context) {
	var dfdfdfyyy sdfsfsfd.Dfdfdfyyyy
	err := c.ShouldBindJSON(&dfdfdfyyy)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := dfdfdfyyyService.CreateDfdfdfyyyy(dfdfdfyyy); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteDfdfdfyyyy 删除Dfdfdfyyyy
// @Tags Dfdfdfyyyy
// @Summary 删除Dfdfdfyyyy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sdfsfsfd.Dfdfdfyyyy true "删除Dfdfdfyyyy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dfdfdfyyy/deleteDfdfdfyyyy [delete]
func (dfdfdfyyyApi *DfdfdfyyyyApi) DeleteDfdfdfyyyy(c *gin.Context) {
	var dfdfdfyyy sdfsfsfd.Dfdfdfyyyy
	err := c.ShouldBindJSON(&dfdfdfyyy)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := dfdfdfyyyService.DeleteDfdfdfyyyy(dfdfdfyyy); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDfdfdfyyyyByIds 批量删除Dfdfdfyyyy
// @Tags Dfdfdfyyyy
// @Summary 批量删除Dfdfdfyyyy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Dfdfdfyyyy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /dfdfdfyyy/deleteDfdfdfyyyyByIds [delete]
func (dfdfdfyyyApi *DfdfdfyyyyApi) DeleteDfdfdfyyyyByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := dfdfdfyyyService.DeleteDfdfdfyyyyByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDfdfdfyyyy 更新Dfdfdfyyyy
// @Tags Dfdfdfyyyy
// @Summary 更新Dfdfdfyyyy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sdfsfsfd.Dfdfdfyyyy true "更新Dfdfdfyyyy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dfdfdfyyy/updateDfdfdfyyyy [put]
func (dfdfdfyyyApi *DfdfdfyyyyApi) UpdateDfdfdfyyyy(c *gin.Context) {
	var dfdfdfyyy sdfsfsfd.Dfdfdfyyyy
	err := c.ShouldBindJSON(&dfdfdfyyy)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := dfdfdfyyyService.UpdateDfdfdfyyyy(dfdfdfyyy); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindDfdfdfyyyy 用id查询Dfdfdfyyyy
// @Tags Dfdfdfyyyy
// @Summary 用id查询Dfdfdfyyyy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query sdfsfsfd.Dfdfdfyyyy true "用id查询Dfdfdfyyyy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dfdfdfyyy/findDfdfdfyyyy [get]
func (dfdfdfyyyApi *DfdfdfyyyyApi) FindDfdfdfyyyy(c *gin.Context) {
	var dfdfdfyyy sdfsfsfd.Dfdfdfyyyy
	err := c.ShouldBindQuery(&dfdfdfyyy)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if redfdfdfyyy, err := dfdfdfyyyService.GetDfdfdfyyyy(dfdfdfyyy.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redfdfdfyyy": redfdfdfyyy}, c)
	}
}

// GetDfdfdfyyyyList 分页获取Dfdfdfyyyy列表
// @Tags Dfdfdfyyyy
// @Summary 分页获取Dfdfdfyyyy列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query sdfsfsfdReq.DfdfdfyyyySearch true "分页获取Dfdfdfyyyy列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dfdfdfyyy/getDfdfdfyyyyList [get]
func (dfdfdfyyyApi *DfdfdfyyyyApi) GetDfdfdfyyyyList(c *gin.Context) {
	var pageInfo sdfsfsfdReq.DfdfdfyyyySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := dfdfdfyyyService.GetDfdfdfyyyyInfoList(pageInfo); err != nil {
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
