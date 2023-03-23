import service from '@/utils/request'

// @Tags Dfdfdfaaa
// @Summary 创建Dfdfdfaaa
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Dfdfdfaaa true "创建Dfdfdfaaa"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dfdfdfaa/createDfdfdfaaa [post]
export const createDfdfdfaaa = (data) => {
  return service({
    url: '/dfdfdfaa/createDfdfdfaaa',
    method: 'post',
    data
  })
}

// @Tags Dfdfdfaaa
// @Summary 删除Dfdfdfaaa
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Dfdfdfaaa true "删除Dfdfdfaaa"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dfdfdfaa/deleteDfdfdfaaa [delete]
export const deleteDfdfdfaaa = (data) => {
  return service({
    url: '/dfdfdfaa/deleteDfdfdfaaa',
    method: 'delete',
    data
  })
}

// @Tags Dfdfdfaaa
// @Summary 删除Dfdfdfaaa
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Dfdfdfaaa"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dfdfdfaa/deleteDfdfdfaaa [delete]
export const deleteDfdfdfaaaByIds = (data) => {
  return service({
    url: '/dfdfdfaa/deleteDfdfdfaaaByIds',
    method: 'delete',
    data
  })
}

// @Tags Dfdfdfaaa
// @Summary 更新Dfdfdfaaa
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Dfdfdfaaa true "更新Dfdfdfaaa"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dfdfdfaa/updateDfdfdfaaa [put]
export const updateDfdfdfaaa = (data) => {
  return service({
    url: '/dfdfdfaa/updateDfdfdfaaa',
    method: 'put',
    data
  })
}

// @Tags Dfdfdfaaa
// @Summary 用id查询Dfdfdfaaa
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Dfdfdfaaa true "用id查询Dfdfdfaaa"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dfdfdfaa/findDfdfdfaaa [get]
export const findDfdfdfaaa = (params) => {
  return service({
    url: '/dfdfdfaa/findDfdfdfaaa',
    method: 'get',
    params
  })
}

// @Tags Dfdfdfaaa
// @Summary 分页获取Dfdfdfaaa列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Dfdfdfaaa列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dfdfdfaa/getDfdfdfaaaList [get]
export const getDfdfdfaaaList = (params) => {
  return service({
    url: '/dfdfdfaa/getDfdfdfaaaList',
    method: 'get',
    params
  })
}
