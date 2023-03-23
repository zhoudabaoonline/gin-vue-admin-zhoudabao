import service from '@/utils/request'

// @Tags Dfdfdf4
// @Summary 创建Dfdfdf4
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Dfdfdf4 true "创建Dfdfdf4"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dfdfdf4/createDfdfdf4 [post]
export const createDfdfdf4 = (data) => {
  return service({
    url: '/dfdfdf4/createDfdfdf4',
    method: 'post',
    data
  })
}

// @Tags Dfdfdf4
// @Summary 删除Dfdfdf4
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Dfdfdf4 true "删除Dfdfdf4"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dfdfdf4/deleteDfdfdf4 [delete]
export const deleteDfdfdf4 = (data) => {
  return service({
    url: '/dfdfdf4/deleteDfdfdf4',
    method: 'delete',
    data
  })
}

// @Tags Dfdfdf4
// @Summary 删除Dfdfdf4
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Dfdfdf4"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dfdfdf4/deleteDfdfdf4 [delete]
export const deleteDfdfdf4ByIds = (data) => {
  return service({
    url: '/dfdfdf4/deleteDfdfdf4ByIds',
    method: 'delete',
    data
  })
}

// @Tags Dfdfdf4
// @Summary 更新Dfdfdf4
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Dfdfdf4 true "更新Dfdfdf4"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dfdfdf4/updateDfdfdf4 [put]
export const updateDfdfdf4 = (data) => {
  return service({
    url: '/dfdfdf4/updateDfdfdf4',
    method: 'put',
    data
  })
}

// @Tags Dfdfdf4
// @Summary 用id查询Dfdfdf4
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Dfdfdf4 true "用id查询Dfdfdf4"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dfdfdf4/findDfdfdf4 [get]
export const findDfdfdf4 = (params) => {
  return service({
    url: '/dfdfdf4/findDfdfdf4',
    method: 'get',
    params
  })
}

// @Tags Dfdfdf4
// @Summary 分页获取Dfdfdf4列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Dfdfdf4列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dfdfdf4/getDfdfdf4List [get]
export const getDfdfdf4List = (params) => {
  return service({
    url: '/dfdfdf4/getDfdfdf4List',
    method: 'get',
    params
  })
}
