import service from '@/utils/request'

// @Tags Ceshi4
// @Summary 创建Ceshi4
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Ceshi4 true "创建Ceshi4"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ceshi4/createCeshi4 [post]
export const createCeshi4 = (data) => {
  return service({
    url: '/ceshi4/createCeshi4',
    method: 'post',
    data
  })
}

// @Tags Ceshi4
// @Summary 删除Ceshi4
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Ceshi4 true "删除Ceshi4"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ceshi4/deleteCeshi4 [delete]
export const deleteCeshi4 = (data) => {
  return service({
    url: '/ceshi4/deleteCeshi4',
    method: 'delete',
    data
  })
}

// @Tags Ceshi4
// @Summary 删除Ceshi4
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Ceshi4"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ceshi4/deleteCeshi4 [delete]
export const deleteCeshi4ByIds = (data) => {
  return service({
    url: '/ceshi4/deleteCeshi4ByIds',
    method: 'delete',
    data
  })
}

// @Tags Ceshi4
// @Summary 更新Ceshi4
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Ceshi4 true "更新Ceshi4"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ceshi4/updateCeshi4 [put]
export const updateCeshi4 = (data) => {
  return service({
    url: '/ceshi4/updateCeshi4',
    method: 'put',
    data
  })
}

// @Tags Ceshi4
// @Summary 用id查询Ceshi4
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Ceshi4 true "用id查询Ceshi4"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ceshi4/findCeshi4 [get]
export const findCeshi4 = (params) => {
  return service({
    url: '/ceshi4/findCeshi4',
    method: 'get',
    params
  })
}

// @Tags Ceshi4
// @Summary 分页获取Ceshi4列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Ceshi4列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ceshi4/getCeshi4List [get]
export const getCeshi4List = (params) => {
  return service({
    url: '/ceshi4/getCeshi4List',
    method: 'get',
    params
  })
}
