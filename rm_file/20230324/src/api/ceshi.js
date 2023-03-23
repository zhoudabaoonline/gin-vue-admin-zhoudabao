import service from '@/utils/request'

// @Tags Ceshi
// @Summary 创建Ceshi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Ceshi true "创建Ceshi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ceshi/createCeshi [post]
export const createCeshi = (data) => {
  return service({
    url: '/ceshi/createCeshi',
    method: 'post',
    data
  })
}

// @Tags Ceshi
// @Summary 删除Ceshi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Ceshi true "删除Ceshi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ceshi/deleteCeshi [delete]
export const deleteCeshi = (data) => {
  return service({
    url: '/ceshi/deleteCeshi',
    method: 'delete',
    data
  })
}

// @Tags Ceshi
// @Summary 删除Ceshi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Ceshi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ceshi/deleteCeshi [delete]
export const deleteCeshiByIds = (data) => {
  return service({
    url: '/ceshi/deleteCeshiByIds',
    method: 'delete',
    data
  })
}

// @Tags Ceshi
// @Summary 更新Ceshi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Ceshi true "更新Ceshi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ceshi/updateCeshi [put]
export const updateCeshi = (data) => {
  return service({
    url: '/ceshi/updateCeshi',
    method: 'put',
    data
  })
}

// @Tags Ceshi
// @Summary 用id查询Ceshi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Ceshi true "用id查询Ceshi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ceshi/findCeshi [get]
export const findCeshi = (params) => {
  return service({
    url: '/ceshi/findCeshi',
    method: 'get',
    params
  })
}

// @Tags Ceshi
// @Summary 分页获取Ceshi列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Ceshi列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ceshi/getCeshiList [get]
export const getCeshiList = (params) => {
  return service({
    url: '/ceshi/getCeshiList',
    method: 'get',
    params
  })
}
