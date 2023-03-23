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

type Struct2Api struct {
}

var struct2Service = service.ServiceGroupApp.Package1ServiceGroup.Struct2Service


// CreateStruct2 创建Struct2
// @Tags Struct2
// @Summary 创建Struct2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body package1.Struct2 true "创建Struct2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /struct2/createStruct2 [post]
func (struct2Api *Struct2Api) CreateStruct2(c *gin.Context) {
	var struct2 package1.Struct2
	err := c.ShouldBindJSON(&struct2)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := struct2Service.CreateStruct2(struct2); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteStruct2 删除Struct2
// @Tags Struct2
// @Summary 删除Struct2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body package1.Struct2 true "删除Struct2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /struct2/deleteStruct2 [delete]
func (struct2Api *Struct2Api) DeleteStruct2(c *gin.Context) {
	var struct2 package1.Struct2
	err := c.ShouldBindJSON(&struct2)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := struct2Service.DeleteStruct2(struct2); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteStruct2ByIds 批量删除Struct2
// @Tags Struct2
// @Summary 批量删除Struct2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Struct2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /struct2/deleteStruct2ByIds [delete]
func (struct2Api *Struct2Api) DeleteStruct2ByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := struct2Service.DeleteStruct2ByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateStruct2 更新Struct2
// @Tags Struct2
// @Summary 更新Struct2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body package1.Struct2 true "更新Struct2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /struct2/updateStruct2 [put]
func (struct2Api *Struct2Api) UpdateStruct2(c *gin.Context) {
	var struct2 package1.Struct2
	err := c.ShouldBindJSON(&struct2)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := struct2Service.UpdateStruct2(struct2); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindStruct2 用id查询Struct2
// @Tags Struct2
// @Summary 用id查询Struct2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query package1.Struct2 true "用id查询Struct2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /struct2/findStruct2 [get]
func (struct2Api *Struct2Api) FindStruct2(c *gin.Context) {
	var struct2 package1.Struct2
	err := c.ShouldBindQuery(&struct2)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if restruct2, err := struct2Service.GetStruct2(struct2.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"restruct2": restruct2}, c)
	}
}

// GetStruct2List 分页获取Struct2列表
// @Tags Struct2
// @Summary 分页获取Struct2列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query package1Req.Struct2Search true "分页获取Struct2列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /struct2/getStruct2List [get]
func (struct2Api *Struct2Api) GetStruct2List(c *gin.Context) {
	var pageInfo package1Req.Struct2Search
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := struct2Service.GetStruct2InfoList(pageInfo); err != nil {
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
