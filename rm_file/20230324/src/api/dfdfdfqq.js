import service from '@/utils/request'

// @Tags Dfdfdfqqq
// @Summary 创建Dfdfdfqqq
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Dfdfdfqqq true "创建Dfdfdfqqq"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dfdfdfqqq/createDfdfdfqqq [post]
export const createDfdfdfqqq = (data) => {
  return service({
    url: '/dfdfdfqqq/createDfdfdfqqq',
    method: 'post',
    data
  })
}

// @Tags Dfdfdfqqq
// @Summary 删除Dfdfdfqqq
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Dfdfdfqqq true "删除Dfdfdfqqq"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dfdfdfqqq/deleteDfdfdfqqq [delete]
export const deleteDfdfdfqqq = (data) => {
  return service({
    url: '/dfdfdfqqq/deleteDfdfdfqqq',
    method: 'delete',
    data
  })
}

// @Tags Dfdfdfqqq
// @Summary 删除Dfdfdfqqq
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Dfdfdfqqq"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dfdfdfqqq/deleteDfdfdfqqq [delete]
export const deleteDfdfdfqqqByIds = (data) => {
  return service({
    url: '/dfdfdfqqq/deleteDfdfdfqqqByIds',
    method: 'delete',
    data
  })
}

// @Tags Dfdfdfqqq
// @Summary 更新Dfdfdfqqq
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Dfdfdfqqq true "更新Dfdfdfqqq"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dfdfdfqqq/updateDfdfdfqqq [put]
export const updateDfdfdfqqq = (data) => {
  return service({
    url: '/dfdfdfqqq/updateDfdfdfqqq',
    method: 'put',
    data
  })
}

// @Tags Dfdfdfqqq
// @Summary 用id查询Dfdfdfqqq
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Dfdfdfqqq true "用id查询Dfdfdfqqq"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dfdfdfqqq/findDfdfdfqqq [get]
export const findDfdfdfqqq = (params) => {
  return service({
    url: '/dfdfdfqqq/findDfdfdfqqq',
    method: 'get',
    params
  })
}

// @Tags Dfdfdfqqq
// @Summary 分页获取Dfdfdfqqq列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Dfdfdfqqq列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dfdfdfqqq/getDfdfdfqqqList [get]
export const getDfdfdfqqqList = (params) => {
  return service({
    url: '/dfdfdfqqq/getDfdfdfqqqList',
    method: 'get',
    params
  })
}
