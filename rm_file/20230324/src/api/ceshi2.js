import service from '@/utils/request'

// @Tags Ceshi2
// @Summary 创建Ceshi2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Ceshi2 true "创建Ceshi2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ceshi2/createCeshi2 [post]
export const createCeshi2 = (data) => {
  return service({
    url: '/ceshi2/createCeshi2',
    method: 'post',
    data
  })
}

// @Tags Ceshi2
// @Summary 删除Ceshi2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Ceshi2 true "删除Ceshi2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ceshi2/deleteCeshi2 [delete]
export const deleteCeshi2 = (data) => {
  return service({
    url: '/ceshi2/deleteCeshi2',
    method: 'delete',
    data
  })
}

// @Tags Ceshi2
// @Summary 删除Ceshi2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Ceshi2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ceshi2/deleteCeshi2 [delete]
export const deleteCeshi2ByIds = (data) => {
  return service({
    url: '/ceshi2/deleteCeshi2ByIds',
    method: 'delete',
    data
  })
}

// @Tags Ceshi2
// @Summary 更新Ceshi2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Ceshi2 true "更新Ceshi2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ceshi2/updateCeshi2 [put]
export const updateCeshi2 = (data) => {
  return service({
    url: '/ceshi2/updateCeshi2',
    method: 'put',
    data
  })
}

// @Tags Ceshi2
// @Summary 用id查询Ceshi2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Ceshi2 true "用id查询Ceshi2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ceshi2/findCeshi2 [get]
export const findCeshi2 = (params) => {
  return service({
    url: '/ceshi2/findCeshi2',
    method: 'get',
    params
  })
}

// @Tags Ceshi2
// @Summary 分页获取Ceshi2列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Ceshi2列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ceshi2/getCeshi2List [get]
export const getCeshi2List = (params) => {
  return service({
    url: '/ceshi2/getCeshi2List',
    method: 'get',
    params
  })
}
