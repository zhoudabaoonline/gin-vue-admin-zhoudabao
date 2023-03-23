import service from '@/utils/request'

// @Tags Dfdfdfvvv
// @Summary 创建Dfdfdfvvv
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Dfdfdfvvv true "创建Dfdfdfvvv"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dfdfdfvvvv/createDfdfdfvvv [post]
export const createDfdfdfvvv = (data) => {
  return service({
    url: '/dfdfdfvvvv/createDfdfdfvvv',
    method: 'post',
    data
  })
}

// @Tags Dfdfdfvvv
// @Summary 删除Dfdfdfvvv
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Dfdfdfvvv true "删除Dfdfdfvvv"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dfdfdfvvvv/deleteDfdfdfvvv [delete]
export const deleteDfdfdfvvv = (data) => {
  return service({
    url: '/dfdfdfvvvv/deleteDfdfdfvvv',
    method: 'delete',
    data
  })
}

// @Tags Dfdfdfvvv
// @Summary 删除Dfdfdfvvv
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Dfdfdfvvv"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dfdfdfvvvv/deleteDfdfdfvvv [delete]
export const deleteDfdfdfvvvByIds = (data) => {
  return service({
    url: '/dfdfdfvvvv/deleteDfdfdfvvvByIds',
    method: 'delete',
    data
  })
}

// @Tags Dfdfdfvvv
// @Summary 更新Dfdfdfvvv
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Dfdfdfvvv true "更新Dfdfdfvvv"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dfdfdfvvvv/updateDfdfdfvvv [put]
export const updateDfdfdfvvv = (data) => {
  return service({
    url: '/dfdfdfvvvv/updateDfdfdfvvv',
    method: 'put',
    data
  })
}

// @Tags Dfdfdfvvv
// @Summary 用id查询Dfdfdfvvv
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Dfdfdfvvv true "用id查询Dfdfdfvvv"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dfdfdfvvvv/findDfdfdfvvv [get]
export const findDfdfdfvvv = (params) => {
  return service({
    url: '/dfdfdfvvvv/findDfdfdfvvv',
    method: 'get',
    params
  })
}

// @Tags Dfdfdfvvv
// @Summary 分页获取Dfdfdfvvv列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Dfdfdfvvv列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dfdfdfvvvv/getDfdfdfvvvList [get]
export const getDfdfdfvvvList = (params) => {
  return service({
    url: '/dfdfdfvvvv/getDfdfdfvvvList',
    method: 'get',
    params
  })
}
