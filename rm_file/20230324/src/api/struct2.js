import service from '@/utils/request'

// @Tags Struct2
// @Summary 创建Struct2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Struct2 true "创建Struct2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /struct2/createStruct2 [post]
export const createStruct2 = (data) => {
  return service({
    url: '/struct2/createStruct2',
    method: 'post',
    data
  })
}

// @Tags Struct2
// @Summary 删除Struct2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Struct2 true "删除Struct2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /struct2/deleteStruct2 [delete]
export const deleteStruct2 = (data) => {
  return service({
    url: '/struct2/deleteStruct2',
    method: 'delete',
    data
  })
}

// @Tags Struct2
// @Summary 删除Struct2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Struct2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /struct2/deleteStruct2 [delete]
export const deleteStruct2ByIds = (data) => {
  return service({
    url: '/struct2/deleteStruct2ByIds',
    method: 'delete',
    data
  })
}

// @Tags Struct2
// @Summary 更新Struct2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Struct2 true "更新Struct2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /struct2/updateStruct2 [put]
export const updateStruct2 = (data) => {
  return service({
    url: '/struct2/updateStruct2',
    method: 'put',
    data
  })
}

// @Tags Struct2
// @Summary 用id查询Struct2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Struct2 true "用id查询Struct2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /struct2/findStruct2 [get]
export const findStruct2 = (params) => {
  return service({
    url: '/struct2/findStruct2',
    method: 'get',
    params
  })
}

// @Tags Struct2
// @Summary 分页获取Struct2列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Struct2列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /struct2/getStruct2List [get]
export const getStruct2List = (params) => {
  return service({
    url: '/struct2/getStruct2List',
    method: 'get',
    params
  })
}
