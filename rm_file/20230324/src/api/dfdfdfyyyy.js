import service from '@/utils/request'

// @Tags Dfdfdfyyyy
// @Summary 创建Dfdfdfyyyy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Dfdfdfyyyy true "创建Dfdfdfyyyy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dfdfdfyyy/createDfdfdfyyyy [post]
export const createDfdfdfyyyy = (data) => {
  return service({
    url: '/dfdfdfyyy/createDfdfdfyyyy',
    method: 'post',
    data
  })
}

// @Tags Dfdfdfyyyy
// @Summary 删除Dfdfdfyyyy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Dfdfdfyyyy true "删除Dfdfdfyyyy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dfdfdfyyy/deleteDfdfdfyyyy [delete]
export const deleteDfdfdfyyyy = (data) => {
  return service({
    url: '/dfdfdfyyy/deleteDfdfdfyyyy',
    method: 'delete',
    data
  })
}

// @Tags Dfdfdfyyyy
// @Summary 删除Dfdfdfyyyy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Dfdfdfyyyy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dfdfdfyyy/deleteDfdfdfyyyy [delete]
export const deleteDfdfdfyyyyByIds = (data) => {
  return service({
    url: '/dfdfdfyyy/deleteDfdfdfyyyyByIds',
    method: 'delete',
    data
  })
}

// @Tags Dfdfdfyyyy
// @Summary 更新Dfdfdfyyyy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Dfdfdfyyyy true "更新Dfdfdfyyyy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dfdfdfyyy/updateDfdfdfyyyy [put]
export const updateDfdfdfyyyy = (data) => {
  return service({
    url: '/dfdfdfyyy/updateDfdfdfyyyy',
    method: 'put',
    data
  })
}

// @Tags Dfdfdfyyyy
// @Summary 用id查询Dfdfdfyyyy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Dfdfdfyyyy true "用id查询Dfdfdfyyyy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dfdfdfyyy/findDfdfdfyyyy [get]
export const findDfdfdfyyyy = (params) => {
  return service({
    url: '/dfdfdfyyy/findDfdfdfyyyy',
    method: 'get',
    params
  })
}

// @Tags Dfdfdfyyyy
// @Summary 分页获取Dfdfdfyyyy列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Dfdfdfyyyy列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dfdfdfyyy/getDfdfdfyyyyList [get]
export const getDfdfdfyyyyList = (params) => {
  return service({
    url: '/dfdfdfyyy/getDfdfdfyyyyList',
    method: 'get',
    params
  })
}
