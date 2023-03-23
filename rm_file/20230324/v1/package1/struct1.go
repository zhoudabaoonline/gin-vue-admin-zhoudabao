package package1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/package1"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    package1Req "github.com/flipped-aurora/gin-vue-admin/server/model/package1/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type Struct1Api struct {
}

var struct1Service = service.ServiceGroupApp.Package1ServiceGroup.Struct1Service


// CreateStruct1 创建Struct1
// @Tags Struct1
// @Summary 创建Struct1
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body package1.Struct1 true "创建Struct1"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /struct1/createStruct1 [post]
func (struct1Api *Struct1Api) CreateStruct1(c *gin.Context) {
	var struct1 package1.Struct1
	err := c.ShouldBindJSON(&struct1)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := struct1Service.CreateStruct1(struct1); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteStruct1 删除Struct1
// @Tags Struct1
// @Summary 删除Struct1
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body package1.Struct1 true "删除Struct1"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /struct1/deleteStruct1 [delete]
func (struct1Api *Struct1Api) DeleteStruct1(c *gin.Context) {
	var struct1 package1.Struct1
	err := c.ShouldBindJSON(&struct1)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := struct1Service.DeleteStruct1(struct1); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteStruct1ByIds 批量删除Struct1
// @Tags Struct1
// @Summary 批量删除Struct1
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Struct1"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /struct1/deleteStruct1ByIds [delete]
func (struct1Api *Struct1Api) DeleteStruct1ByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := struct1Service.DeleteStruct1ByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateStruct1 更新Struct1
// @Tags Struct1
// @Summary 更新Struct1
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body package1.Struct1 true "更新Struct1"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /struct1/updateStruct1 [put]
func (struct1Api *Struct1Api) UpdateStruct1(c *gin.Context) {
	var struct1 package1.Struct1
	err := c.ShouldBindJSON(&struct1)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := struct1Service.UpdateStruct1(struct1); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindStruct1 用id查询Struct1
// @Tags Struct1
// @Summary 用id查询Struct1
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query package1.Struct1 true "用id查询Struct1"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /struct1/findStruct1 [get]
func (struct1Api *Struct1Api) FindStruct1(c *gin.Context) {
	var struct1 package1.Struct1
	err := c.ShouldBindQuery(&struct1)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if restruct1, err := struct1Service.GetStruct1(struct1.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"restruct1": restruct1}, c)
	}
}

// GetStruct1List 分页获取Struct1列表
// @Tags Struct1
// @Summary 分页获取Struct1列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query package1Req.Struct1Search true "分页获取Struct1列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /struct1/getStruct1List [get]
func (struct1Api *Struct1Api) GetStruct1List(c *gin.Context) {
	var pageInfo package1Req.Struct1Search
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := struct1Service.GetStruct1InfoList(pageInfo); err != nil {
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
