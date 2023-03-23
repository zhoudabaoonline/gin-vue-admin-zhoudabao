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

type Strcut3Api struct {
}

var strcut3Service = service.ServiceGroupApp.Package1ServiceGroup.Strcut3Service


// CreateStrcut3 创建Strcut3
// @Tags Strcut3
// @Summary 创建Strcut3
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body package1.Strcut3 true "创建Strcut3"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /strcut3/createStrcut3 [post]
func (strcut3Api *Strcut3Api) CreateStrcut3(c *gin.Context) {
	var strcut3 package1.Strcut3
	err := c.ShouldBindJSON(&strcut3)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := strcut3Service.CreateStrcut3(strcut3); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteStrcut3 删除Strcut3
// @Tags Strcut3
// @Summary 删除Strcut3
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body package1.Strcut3 true "删除Strcut3"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /strcut3/deleteStrcut3 [delete]
func (strcut3Api *Strcut3Api) DeleteStrcut3(c *gin.Context) {
	var strcut3 package1.Strcut3
	err := c.ShouldBindJSON(&strcut3)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := strcut3Service.DeleteStrcut3(strcut3); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteStrcut3ByIds 批量删除Strcut3
// @Tags Strcut3
// @Summary 批量删除Strcut3
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Strcut3"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /strcut3/deleteStrcut3ByIds [delete]
func (strcut3Api *Strcut3Api) DeleteStrcut3ByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := strcut3Service.DeleteStrcut3ByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateStrcut3 更新Strcut3
// @Tags Strcut3
// @Summary 更新Strcut3
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body package1.Strcut3 true "更新Strcut3"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /strcut3/updateStrcut3 [put]
func (strcut3Api *Strcut3Api) UpdateStrcut3(c *gin.Context) {
	var strcut3 package1.Strcut3
	err := c.ShouldBindJSON(&strcut3)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := strcut3Service.UpdateStrcut3(strcut3); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindStrcut3 用id查询Strcut3
// @Tags Strcut3
// @Summary 用id查询Strcut3
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query package1.Strcut3 true "用id查询Strcut3"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /strcut3/findStrcut3 [get]
func (strcut3Api *Strcut3Api) FindStrcut3(c *gin.Context) {
	var strcut3 package1.Strcut3
	err := c.ShouldBindQuery(&strcut3)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if restrcut3, err := strcut3Service.GetStrcut3(strcut3.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"restrcut3": restrcut3}, c)
	}
}

// GetStrcut3List 分页获取Strcut3列表
// @Tags Strcut3
// @Summary 分页获取Strcut3列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query package1Req.Strcut3Search true "分页获取Strcut3列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /strcut3/getStrcut3List [get]
func (strcut3Api *Strcut3Api) GetStrcut3List(c *gin.Context) {
	var pageInfo package1Req.Strcut3Search
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := strcut3Service.GetStrcut3InfoList(pageInfo); err != nil {
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
