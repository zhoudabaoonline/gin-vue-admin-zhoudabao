import service from '@/utils/request'

// @Tags Dfdfdf22
// @Summary 创建Dfdfdf22
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Dfdfdf22 true "创建Dfdfdf22"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dfdfdf22/createDfdfdf22 [post]
export const createDfdfdf22 = (data) => {
  return service({
    url: '/dfdfdf22/createDfdfdf22',
    method: 'post',
    data
  })
}

// @Tags Dfdfdf22
// @Summary 删除Dfdfdf22
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Dfdfdf22 true "删除Dfdfdf22"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dfdfdf22/deleteDfdfdf22 [delete]
export const deleteDfdfdf22 = (data) => {
  return service({
    url: '/dfdfdf22/deleteDfdfdf22',
    method: 'delete',
    data
  })
}

// @Tags Dfdfdf22
// @Summary 删除Dfdfdf22
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Dfdfdf22"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dfdfdf22/deleteDfdfdf22 [delete]
export const deleteDfdfdf22ByIds = (data) => {
  return service({
    url: '/dfdfdf22/deleteDfdfdf22ByIds',
    method: 'delete',
    data
  })
}

// @Tags Dfdfdf22
// @Summary 更新Dfdfdf22
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Dfdfdf22 true "更新Dfdfdf22"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dfdfdf22/updateDfdfdf22 [put]
export const updateDfdfdf22 = (data) => {
  return service({
    url: '/dfdfdf22/updateDfdfdf22',
    method: 'put',
    data
  })
}

// @Tags Dfdfdf22
// @Summary 用id查询Dfdfdf22
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Dfdfdf22 true "用id查询Dfdfdf22"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dfdfdf22/findDfdfdf22 [get]
export const findDfdfdf22 = (params) => {
  return service({
    url: '/dfdfdf22/findDfdfdf22',
    method: 'get',
    params
  })
}

// @Tags Dfdfdf22
// @Summary 分页获取Dfdfdf22列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Dfdfdf22列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dfdfdf22/getDfdfdf22List [get]
export const getDfdfdf22List = (params) => {
  return service({
    url: '/dfdfdf22/getDfdfdf22List',
    method: 'get',
    params
  })
}
