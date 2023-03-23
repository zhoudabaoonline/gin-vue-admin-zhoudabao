import service from '@/utils/request'

// @Tags Dfdfdfvvvv
// @Summary 创建Dfdfdfvvvv
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Dfdfdfvvvv true "创建Dfdfdfvvvv"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dfdfdfvv/createDfdfdfvvvv [post]
export const createDfdfdfvvvv = (data) => {
  return service({
    url: '/dfdfdfvv/createDfdfdfvvvv',
    method: 'post',
    data
  })
}

// @Tags Dfdfdfvvvv
// @Summary 删除Dfdfdfvvvv
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Dfdfdfvvvv true "删除Dfdfdfvvvv"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dfdfdfvv/deleteDfdfdfvvvv [delete]
export const deleteDfdfdfvvvv = (data) => {
  return service({
    url: '/dfdfdfvv/deleteDfdfdfvvvv',
    method: 'delete',
    data
  })
}

// @Tags Dfdfdfvvvv
// @Summary 删除Dfdfdfvvvv
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Dfdfdfvvvv"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dfdfdfvv/deleteDfdfdfvvvv [delete]
export const deleteDfdfdfvvvvByIds = (data) => {
  return service({
    url: '/dfdfdfvv/deleteDfdfdfvvvvByIds',
    method: 'delete',
    data
  })
}

// @Tags Dfdfdfvvvv
// @Summary 更新Dfdfdfvvvv
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Dfdfdfvvvv true "更新Dfdfdfvvvv"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dfdfdfvv/updateDfdfdfvvvv [put]
export const updateDfdfdfvvvv = (data) => {
  return service({
    url: '/dfdfdfvv/updateDfdfdfvvvv',
    method: 'put',
    data
  })
}

// @Tags Dfdfdfvvvv
// @Summary 用id查询Dfdfdfvvvv
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Dfdfdfvvvv true "用id查询Dfdfdfvvvv"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dfdfdfvv/findDfdfdfvvvv [get]
export const findDfdfdfvvvv = (params) => {
  return service({
    url: '/dfdfdfvv/findDfdfdfvvvv',
    method: 'get',
    params
  })
}

// @Tags Dfdfdfvvvv
// @Summary 分页获取Dfdfdfvvvv列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Dfdfdfvvvv列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dfdfdfvv/getDfdfdfvvvvList [get]
export const getDfdfdfvvvvList = (params) => {
  return service({
    url: '/dfdfdfvv/getDfdfdfvvvvList',
    method: 'get',
    params
  })
}
