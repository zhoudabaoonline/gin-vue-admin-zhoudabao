import service from '@/utils/request'

// @Tags Dfdfdfcc
// @Summary 创建Dfdfdfcc
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Dfdfdfcc true "创建Dfdfdfcc"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dfdfdfcc/createDfdfdfcc [post]
export const createDfdfdfcc = (data) => {
  return service({
    url: '/dfdfdfcc/createDfdfdfcc',
    method: 'post',
    data
  })
}

// @Tags Dfdfdfcc
// @Summary 删除Dfdfdfcc
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Dfdfdfcc true "删除Dfdfdfcc"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dfdfdfcc/deleteDfdfdfcc [delete]
export const deleteDfdfdfcc = (data) => {
  return service({
    url: '/dfdfdfcc/deleteDfdfdfcc',
    method: 'delete',
    data
  })
}

// @Tags Dfdfdfcc
// @Summary 删除Dfdfdfcc
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Dfdfdfcc"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dfdfdfcc/deleteDfdfdfcc [delete]
export const deleteDfdfdfccByIds = (data) => {
  return service({
    url: '/dfdfdfcc/deleteDfdfdfccByIds',
    method: 'delete',
    data
  })
}

// @Tags Dfdfdfcc
// @Summary 更新Dfdfdfcc
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Dfdfdfcc true "更新Dfdfdfcc"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dfdfdfcc/updateDfdfdfcc [put]
export const updateDfdfdfcc = (data) => {
  return service({
    url: '/dfdfdfcc/updateDfdfdfcc',
    method: 'put',
    data
  })
}

// @Tags Dfdfdfcc
// @Summary 用id查询Dfdfdfcc
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Dfdfdfcc true "用id查询Dfdfdfcc"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dfdfdfcc/findDfdfdfcc [get]
export const findDfdfdfcc = (params) => {
  return service({
    url: '/dfdfdfcc/findDfdfdfcc',
    method: 'get',
    params
  })
}

// @Tags Dfdfdfcc
// @Summary 分页获取Dfdfdfcc列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Dfdfdfcc列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dfdfdfcc/getDfdfdfccList [get]
export const getDfdfdfccList = (params) => {
  return service({
    url: '/dfdfdfcc/getDfdfdfccList',
    method: 'get',
    params
  })
}
