import service from '@/utils/request'

// @Tags Struct4
// @Summary 创建Struct4
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Struct4 true "创建Struct4"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /struct4/createStruct4 [post]
export const createStruct4 = (data) => {
  return service({
    url: '/struct4/createStruct4',
    method: 'post',
    data
  })
}

// @Tags Struct4
// @Summary 删除Struct4
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Struct4 true "删除Struct4"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /struct4/deleteStruct4 [delete]
export const deleteStruct4 = (data) => {
  return service({
    url: '/struct4/deleteStruct4',
    method: 'delete',
    data
  })
}

// @Tags Struct4
// @Summary 删除Struct4
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Struct4"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /struct4/deleteStruct4 [delete]
export const deleteStruct4ByIds = (data) => {
  return service({
    url: '/struct4/deleteStruct4ByIds',
    method: 'delete',
    data
  })
}

// @Tags Struct4
// @Summary 更新Struct4
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Struct4 true "更新Struct4"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /struct4/updateStruct4 [put]
export const updateStruct4 = (data) => {
  return service({
    url: '/struct4/updateStruct4',
    method: 'put',
    data
  })
}

// @Tags Struct4
// @Summary 用id查询Struct4
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Struct4 true "用id查询Struct4"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /struct4/findStruct4 [get]
export const findStruct4 = (params) => {
  return service({
    url: '/struct4/findStruct4',
    method: 'get',
    params
  })
}

// @Tags Struct4
// @Summary 分页获取Struct4列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Struct4列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /struct4/getStruct4List [get]
export const getStruct4List = (params) => {
  return service({
    url: '/struct4/getStruct4List',
    method: 'get',
    params
  })
}
