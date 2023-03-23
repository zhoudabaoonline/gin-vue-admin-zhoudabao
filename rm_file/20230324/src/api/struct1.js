import service from '@/utils/request'

// @Tags Struct1
// @Summary 创建Struct1
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Struct1 true "创建Struct1"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /struct1/createStruct1 [post]
export const createStruct1 = (data) => {
  return service({
    url: '/struct1/createStruct1',
    method: 'post',
    data
  })
}

// @Tags Struct1
// @Summary 删除Struct1
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Struct1 true "删除Struct1"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /struct1/deleteStruct1 [delete]
export const deleteStruct1 = (data) => {
  return service({
    url: '/struct1/deleteStruct1',
    method: 'delete',
    data
  })
}

// @Tags Struct1
// @Summary 删除Struct1
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Struct1"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /struct1/deleteStruct1 [delete]
export const deleteStruct1ByIds = (data) => {
  return service({
    url: '/struct1/deleteStruct1ByIds',
    method: 'delete',
    data
  })
}

// @Tags Struct1
// @Summary 更新Struct1
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Struct1 true "更新Struct1"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /struct1/updateStruct1 [put]
export const updateStruct1 = (data) => {
  return service({
    url: '/struct1/updateStruct1',
    method: 'put',
    data
  })
}

// @Tags Struct1
// @Summary 用id查询Struct1
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Struct1 true "用id查询Struct1"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /struct1/findStruct1 [get]
export const findStruct1 = (params) => {
  return service({
    url: '/struct1/findStruct1',
    method: 'get',
    params
  })
}

// @Tags Struct1
// @Summary 分页获取Struct1列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Struct1列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /struct1/getStruct1List [get]
export const getStruct1List = (params) => {
  return service({
    url: '/struct1/getStruct1List',
    method: 'get',
    params
  })
}
