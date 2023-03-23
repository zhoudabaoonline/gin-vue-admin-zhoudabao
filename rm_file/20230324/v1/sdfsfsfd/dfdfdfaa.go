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

type DfdfdfaaaApi struct {
}

var dfdfdfaaService = service.ServiceGroupApp.SdfsfsfdServiceGroup.DfdfdfaaaService


// CreateDfdfdfaaa 创建Dfdfdfaaa
// @Tags Dfdfdfaaa
// @Summary 创建Dfdfdfaaa
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sdfsfsfd.Dfdfdfaaa true "创建Dfdfdfaaa"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dfdfdfaa/createDfdfdfaaa [post]
func (dfdfdfaaApi *DfdfdfaaaApi) CreateDfdfdfaaa(c *gin.Context) {
	var dfdfdfaa sdfsfsfd.Dfdfdfaaa
	err := c.ShouldBindJSON(&dfdfdfaa)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := dfdfdfaaService.CreateDfdfdfaaa(dfdfdfaa); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteDfdfdfaaa 删除Dfdfdfaaa
// @Tags Dfdfdfaaa
// @Summary 删除Dfdfdfaaa
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sdfsfsfd.Dfdfdfaaa true "删除Dfdfdfaaa"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dfdfdfaa/deleteDfdfdfaaa [delete]
func (dfdfdfaaApi *DfdfdfaaaApi) DeleteDfdfdfaaa(c *gin.Context) {
	var dfdfdfaa sdfsfsfd.Dfdfdfaaa
	err := c.ShouldBindJSON(&dfdfdfaa)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := dfdfdfaaService.DeleteDfdfdfaaa(dfdfdfaa); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDfdfdfaaaByIds 批量删除Dfdfdfaaa
// @Tags Dfdfdfaaa
// @Summary 批量删除Dfdfdfaaa
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Dfdfdfaaa"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /dfdfdfaa/deleteDfdfdfaaaByIds [delete]
func (dfdfdfaaApi *DfdfdfaaaApi) DeleteDfdfdfaaaByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := dfdfdfaaService.DeleteDfdfdfaaaByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDfdfdfaaa 更新Dfdfdfaaa
// @Tags Dfdfdfaaa
// @Summary 更新Dfdfdfaaa
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sdfsfsfd.Dfdfdfaaa true "更新Dfdfdfaaa"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dfdfdfaa/updateDfdfdfaaa [put]
func (dfdfdfaaApi *DfdfdfaaaApi) UpdateDfdfdfaaa(c *gin.Context) {
	var dfdfdfaa sdfsfsfd.Dfdfdfaaa
	err := c.ShouldBindJSON(&dfdfdfaa)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := dfdfdfaaService.UpdateDfdfdfaaa(dfdfdfaa); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindDfdfdfaaa 用id查询Dfdfdfaaa
// @Tags Dfdfdfaaa
// @Summary 用id查询Dfdfdfaaa
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query sdfsfsfd.Dfdfdfaaa true "用id查询Dfdfdfaaa"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dfdfdfaa/findDfdfdfaaa [get]
func (dfdfdfaaApi *DfdfdfaaaApi) FindDfdfdfaaa(c *gin.Context) {
	var dfdfdfaa sdfsfsfd.Dfdfdfaaa
	err := c.ShouldBindQuery(&dfdfdfaa)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if redfdfdfaa, err := dfdfdfaaService.GetDfdfdfaaa(dfdfdfaa.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redfdfdfaa": redfdfdfaa}, c)
	}
}

// GetDfdfdfaaaList 分页获取Dfdfdfaaa列表
// @Tags Dfdfdfaaa
// @Summary 分页获取Dfdfdfaaa列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query sdfsfsfdReq.DfdfdfaaaSearch true "分页获取Dfdfdfaaa列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dfdfdfaa/getDfdfdfaaaList [get]
func (dfdfdfaaApi *DfdfdfaaaApi) GetDfdfdfaaaList(c *gin.Context) {
	var pageInfo sdfsfsfdReq.DfdfdfaaaSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := dfdfdfaaService.GetDfdfdfaaaInfoList(pageInfo); err != nil {
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
