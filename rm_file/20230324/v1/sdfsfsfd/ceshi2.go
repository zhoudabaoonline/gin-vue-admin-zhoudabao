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

type Ceshi2Api struct {
}

var ceshi2Service = service.ServiceGroupApp.SdfsfsfdServiceGroup.Ceshi2Service


// CreateCeshi2 创建Ceshi2
// @Tags Ceshi2
// @Summary 创建Ceshi2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sdfsfsfd.Ceshi2 true "创建Ceshi2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ceshi2/createCeshi2 [post]
func (ceshi2Api *Ceshi2Api) CreateCeshi2(c *gin.Context) {
	var ceshi2 sdfsfsfd.Ceshi2
	err := c.ShouldBindJSON(&ceshi2)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ceshi2Service.CreateCeshi2(ceshi2); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCeshi2 删除Ceshi2
// @Tags Ceshi2
// @Summary 删除Ceshi2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sdfsfsfd.Ceshi2 true "删除Ceshi2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ceshi2/deleteCeshi2 [delete]
func (ceshi2Api *Ceshi2Api) DeleteCeshi2(c *gin.Context) {
	var ceshi2 sdfsfsfd.Ceshi2
	err := c.ShouldBindJSON(&ceshi2)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ceshi2Service.DeleteCeshi2(ceshi2); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCeshi2ByIds 批量删除Ceshi2
// @Tags Ceshi2
// @Summary 批量删除Ceshi2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Ceshi2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /ceshi2/deleteCeshi2ByIds [delete]
func (ceshi2Api *Ceshi2Api) DeleteCeshi2ByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ceshi2Service.DeleteCeshi2ByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCeshi2 更新Ceshi2
// @Tags Ceshi2
// @Summary 更新Ceshi2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sdfsfsfd.Ceshi2 true "更新Ceshi2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ceshi2/updateCeshi2 [put]
func (ceshi2Api *Ceshi2Api) UpdateCeshi2(c *gin.Context) {
	var ceshi2 sdfsfsfd.Ceshi2
	err := c.ShouldBindJSON(&ceshi2)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ceshi2Service.UpdateCeshi2(ceshi2); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCeshi2 用id查询Ceshi2
// @Tags Ceshi2
// @Summary 用id查询Ceshi2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query sdfsfsfd.Ceshi2 true "用id查询Ceshi2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ceshi2/findCeshi2 [get]
func (ceshi2Api *Ceshi2Api) FindCeshi2(c *gin.Context) {
	var ceshi2 sdfsfsfd.Ceshi2
	err := c.ShouldBindQuery(&ceshi2)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if receshi2, err := ceshi2Service.GetCeshi2(ceshi2.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"receshi2": receshi2}, c)
	}
}

// GetCeshi2List 分页获取Ceshi2列表
// @Tags Ceshi2
// @Summary 分页获取Ceshi2列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query sdfsfsfdReq.Ceshi2Search true "分页获取Ceshi2列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ceshi2/getCeshi2List [get]
func (ceshi2Api *Ceshi2Api) GetCeshi2List(c *gin.Context) {
	var pageInfo sdfsfsfdReq.Ceshi2Search
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := ceshi2Service.GetCeshi2InfoList(pageInfo); err != nil {
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
