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

type DfdfdffgApi struct {
}

var dfdfdffgService = service.ServiceGroupApp.SdfsfsfdServiceGroup.DfdfdffgService


// CreateDfdfdffg 创建Dfdfdffg
// @Tags Dfdfdffg
// @Summary 创建Dfdfdffg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sdfsfsfd.Dfdfdffg true "创建Dfdfdffg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dfdfdffg/createDfdfdffg [post]
func (dfdfdffgApi *DfdfdffgApi) CreateDfdfdffg(c *gin.Context) {
	var dfdfdffg sdfsfsfd.Dfdfdffg
	err := c.ShouldBindJSON(&dfdfdffg)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := dfdfdffgService.CreateDfdfdffg(dfdfdffg); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteDfdfdffg 删除Dfdfdffg
// @Tags Dfdfdffg
// @Summary 删除Dfdfdffg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sdfsfsfd.Dfdfdffg true "删除Dfdfdffg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dfdfdffg/deleteDfdfdffg [delete]
func (dfdfdffgApi *DfdfdffgApi) DeleteDfdfdffg(c *gin.Context) {
	var dfdfdffg sdfsfsfd.Dfdfdffg
	err := c.ShouldBindJSON(&dfdfdffg)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := dfdfdffgService.DeleteDfdfdffg(dfdfdffg); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDfdfdffgByIds 批量删除Dfdfdffg
// @Tags Dfdfdffg
// @Summary 批量删除Dfdfdffg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Dfdfdffg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /dfdfdffg/deleteDfdfdffgByIds [delete]
func (dfdfdffgApi *DfdfdffgApi) DeleteDfdfdffgByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := dfdfdffgService.DeleteDfdfdffgByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDfdfdffg 更新Dfdfdffg
// @Tags Dfdfdffg
// @Summary 更新Dfdfdffg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sdfsfsfd.Dfdfdffg true "更新Dfdfdffg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dfdfdffg/updateDfdfdffg [put]
func (dfdfdffgApi *DfdfdffgApi) UpdateDfdfdffg(c *gin.Context) {
	var dfdfdffg sdfsfsfd.Dfdfdffg
	err := c.ShouldBindJSON(&dfdfdffg)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := dfdfdffgService.UpdateDfdfdffg(dfdfdffg); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindDfdfdffg 用id查询Dfdfdffg
// @Tags Dfdfdffg
// @Summary 用id查询Dfdfdffg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query sdfsfsfd.Dfdfdffg true "用id查询Dfdfdffg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dfdfdffg/findDfdfdffg [get]
func (dfdfdffgApi *DfdfdffgApi) FindDfdfdffg(c *gin.Context) {
	var dfdfdffg sdfsfsfd.Dfdfdffg
	err := c.ShouldBindQuery(&dfdfdffg)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if redfdfdffg, err := dfdfdffgService.GetDfdfdffg(dfdfdffg.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redfdfdffg": redfdfdffg}, c)
	}
}

// GetDfdfdffgList 分页获取Dfdfdffg列表
// @Tags Dfdfdffg
// @Summary 分页获取Dfdfdffg列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query sdfsfsfdReq.DfdfdffgSearch true "分页获取Dfdfdffg列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dfdfdffg/getDfdfdffgList [get]
func (dfdfdffgApi *DfdfdffgApi) GetDfdfdffgList(c *gin.Context) {
	var pageInfo sdfsfsfdReq.DfdfdffgSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := dfdfdffgService.GetDfdfdffgInfoList(pageInfo); err != nil {
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
