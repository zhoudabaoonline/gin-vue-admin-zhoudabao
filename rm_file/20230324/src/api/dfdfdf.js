import service from '@/utils/request'

// @Tags Dfdfdf
// @Summary 创建Dfdfdf
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Dfdfdf true "创建Dfdfdf"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dfdfdf/createDfdfdf [post]
export const createDfdfdf = (data) => {
  return service({
    url: '/dfdfdf/createDfdfdf',
    method: 'post',
    data
  })
}

// @Tags Dfdfdf
// @Summary 删除Dfdfdf
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Dfdfdf true "删除Dfdfdf"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dfdfdf/deleteDfdfdf [delete]
export const deleteDfdfdf = (data) => {
  return service({
    url: '/dfdfdf/deleteDfdfdf',
    method: 'delete',
    data
  })
}

// @Tags Dfdfdf
// @Summary 删除Dfdfdf
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Dfdfdf"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dfdfdf/deleteDfdfdf [delete]
export const deleteDfdfdfByIds = (data) => {
  return service({
    url: '/dfdfdf/deleteDfdfdfByIds',
    method: 'delete',
    data
  })
}

// @Tags Dfdfdf
// @Summary 更新Dfdfdf
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Dfdfdf true "更新Dfdfdf"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dfdfdf/updateDfdfdf [put]
export const updateDfdfdf = (data) => {
  return service({
    url: '/dfdfdf/updateDfdfdf',
    method: 'put',
    data
  })
}

// @Tags Dfdfdf
// @Summary 用id查询Dfdfdf
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Dfdfdf true "用id查询Dfdfdf"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dfdfdf/findDfdfdf [get]
export const findDfdfdf = (params) => {
  return service({
    url: '/dfdfdf/findDfdfdf',
    method: 'get',
    params
  })
}

// @Tags Dfdfdf
// @Summary 分页获取Dfdfdf列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Dfdfdf列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dfdfdf/getDfdfdfList [get]
export const getDfdfdfList = (params) => {
  return service({
    url: '/dfdfdf/getDfdfdfList',
    method: 'get',
    params
  })
}
