import service from '@/utils/request'

// @Tags Dfdfdfssss
// @Summary 创建Dfdfdfssss
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Dfdfdfssss true "创建Dfdfdfssss"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dfdfdfsss/createDfdfdfssss [post]
export const createDfdfdfssss = (data) => {
  return service({
    url: '/dfdfdfsss/createDfdfdfssss',
    method: 'post',
    data
  })
}

// @Tags Dfdfdfssss
// @Summary 删除Dfdfdfssss
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Dfdfdfssss true "删除Dfdfdfssss"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dfdfdfsss/deleteDfdfdfssss [delete]
export const deleteDfdfdfssss = (data) => {
  return service({
    url: '/dfdfdfsss/deleteDfdfdfssss',
    method: 'delete',
    data
  })
}

// @Tags Dfdfdfssss
// @Summary 删除Dfdfdfssss
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Dfdfdfssss"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dfdfdfsss/deleteDfdfdfssss [delete]
export const deleteDfdfdfssssByIds = (data) => {
  return service({
    url: '/dfdfdfsss/deleteDfdfdfssssByIds',
    method: 'delete',
    data
  })
}

// @Tags Dfdfdfssss
// @Summary 更新Dfdfdfssss
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Dfdfdfssss true "更新Dfdfdfssss"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dfdfdfsss/updateDfdfdfssss [put]
export const updateDfdfdfssss = (data) => {
  return service({
    url: '/dfdfdfsss/updateDfdfdfssss',
    method: 'put',
    data
  })
}

// @Tags Dfdfdfssss
// @Summary 用id查询Dfdfdfssss
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Dfdfdfssss true "用id查询Dfdfdfssss"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dfdfdfsss/findDfdfdfssss [get]
export const findDfdfdfssss = (params) => {
  return service({
    url: '/dfdfdfsss/findDfdfdfssss',
    method: 'get',
    params
  })
}

// @Tags Dfdfdfssss
// @Summary 分页获取Dfdfdfssss列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Dfdfdfssss列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dfdfdfsss/getDfdfdfssssList [get]
export const getDfdfdfssssList = (params) => {
  return service({
    url: '/dfdfdfsss/getDfdfdfssssList',
    method: 'get',
    params
  })
}
