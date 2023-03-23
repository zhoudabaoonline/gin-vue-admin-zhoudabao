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

type Struct4Api struct {
}

var struct4Service = service.ServiceGroupApp.Package1ServiceGroup.Struct4Service


// CreateStruct4 创建Struct4
// @Tags Struct4
// @Summary 创建Struct4
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body package1.Struct4 true "创建Struct4"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /struct4/createStruct4 [post]
func (struct4Api *Struct4Api) CreateStruct4(c *gin.Context) {
	var struct4 package1.Struct4
	err := c.ShouldBindJSON(&struct4)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := struct4Service.CreateStruct4(struct4); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteStruct4 删除Struct4
// @Tags Struct4
// @Summary 删除Struct4
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body package1.Struct4 true "删除Struct4"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /struct4/deleteStruct4 [delete]
func (struct4Api *Struct4Api) DeleteStruct4(c *gin.Context) {
	var struct4 package1.Struct4
	err := c.ShouldBindJSON(&struct4)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := struct4Service.DeleteStruct4(struct4); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteStruct4ByIds 批量删除Struct4
// @Tags Struct4
// @Summary 批量删除Struct4
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Struct4"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /struct4/deleteStruct4ByIds [delete]
func (struct4Api *Struct4Api) DeleteStruct4ByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := struct4Service.DeleteStruct4ByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateStruct4 更新Struct4
// @Tags Struct4
// @Summary 更新Struct4
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body package1.Struct4 true "更新Struct4"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /struct4/updateStruct4 [put]
func (struct4Api *Struct4Api) UpdateStruct4(c *gin.Context) {
	var struct4 package1.Struct4
	err := c.ShouldBindJSON(&struct4)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := struct4Service.UpdateStruct4(struct4); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindStruct4 用id查询Struct4
// @Tags Struct4
// @Summary 用id查询Struct4
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query package1.Struct4 true "用id查询Struct4"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /struct4/findStruct4 [get]
func (struct4Api *Struct4Api) FindStruct4(c *gin.Context) {
	var struct4 package1.Struct4
	err := c.ShouldBindQuery(&struct4)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if restruct4, err := struct4Service.GetStruct4(struct4.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"restruct4": restruct4}, c)
	}
}

// GetStruct4List 分页获取Struct4列表
// @Tags Struct4
// @Summary 分页获取Struct4列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query package1Req.Struct4Search true "分页获取Struct4列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /struct4/getStruct4List [get]
func (struct4Api *Struct4Api) GetStruct4List(c *gin.Context) {
	var pageInfo package1Req.Struct4Search
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := struct4Service.GetStruct4InfoList(pageInfo); err != nil {
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
