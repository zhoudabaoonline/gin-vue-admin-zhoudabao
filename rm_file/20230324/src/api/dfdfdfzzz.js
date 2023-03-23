import service from '@/utils/request'

// @Tags Dfdfdfzz
// @Summary 创建Dfdfdfzz
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Dfdfdfzz true "创建Dfdfdfzz"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dfdfdfzzzz/createDfdfdfzz [post]
export const createDfdfdfzz = (data) => {
  return service({
    url: '/dfdfdfzzzz/createDfdfdfzz',
    method: 'post',
    data
  })
}

// @Tags Dfdfdfzz
// @Summary 删除Dfdfdfzz
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Dfdfdfzz true "删除Dfdfdfzz"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dfdfdfzzzz/deleteDfdfdfzz [delete]
export const deleteDfdfdfzz = (data) => {
  return service({
    url: '/dfdfdfzzzz/deleteDfdfdfzz',
    method: 'delete',
    data
  })
}

// @Tags Dfdfdfzz
// @Summary 删除Dfdfdfzz
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Dfdfdfzz"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dfdfdfzzzz/deleteDfdfdfzz [delete]
export const deleteDfdfdfzzByIds = (data) => {
  return service({
    url: '/dfdfdfzzzz/deleteDfdfdfzzByIds',
    method: 'delete',
    data
  })
}

// @Tags Dfdfdfzz
// @Summary 更新Dfdfdfzz
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Dfdfdfzz true "更新Dfdfdfzz"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dfdfdfzzzz/updateDfdfdfzz [put]
export const updateDfdfdfzz = (data) => {
  return service({
    url: '/dfdfdfzzzz/updateDfdfdfzz',
    method: 'put',
    data
  })
}

// @Tags Dfdfdfzz
// @Summary 用id查询Dfdfdfzz
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Dfdfdfzz true "用id查询Dfdfdfzz"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dfdfdfzzzz/findDfdfdfzz [get]
export const findDfdfdfzz = (params) => {
  return service({
    url: '/dfdfdfzzzz/findDfdfdfzz',
    method: 'get',
    params
  })
}

// @Tags Dfdfdfzz
// @Summary 分页获取Dfdfdfzz列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Dfdfdfzz列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dfdfdfzzzz/getDfdfdfzzList [get]
export const getDfdfdfzzList = (params) => {
  return service({
    url: '/dfdfdfzzzz/getDfdfdfzzList',
    method: 'get',
    params
  })
}
