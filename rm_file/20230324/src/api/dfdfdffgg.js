import service from '@/utils/request'

// @Tags Dfdfdffg
// @Summary 创建Dfdfdffg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Dfdfdffg true "创建Dfdfdffg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dfdfdffg/createDfdfdffg [post]
export const createDfdfdffg = (data) => {
  return service({
    url: '/dfdfdffg/createDfdfdffg',
    method: 'post',
    data
  })
}

// @Tags Dfdfdffg
// @Summary 删除Dfdfdffg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Dfdfdffg true "删除Dfdfdffg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dfdfdffg/deleteDfdfdffg [delete]
export const deleteDfdfdffg = (data) => {
  return service({
    url: '/dfdfdffg/deleteDfdfdffg',
    method: 'delete',
    data
  })
}

// @Tags Dfdfdffg
// @Summary 删除Dfdfdffg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Dfdfdffg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dfdfdffg/deleteDfdfdffg [delete]
export const deleteDfdfdffgByIds = (data) => {
  return service({
    url: '/dfdfdffg/deleteDfdfdffgByIds',
    method: 'delete',
    data
  })
}

// @Tags Dfdfdffg
// @Summary 更新Dfdfdffg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Dfdfdffg true "更新Dfdfdffg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dfdfdffg/updateDfdfdffg [put]
export const updateDfdfdffg = (data) => {
  return service({
    url: '/dfdfdffg/updateDfdfdffg',
    method: 'put',
    data
  })
}

// @Tags Dfdfdffg
// @Summary 用id查询Dfdfdffg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Dfdfdffg true "用id查询Dfdfdffg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dfdfdffg/findDfdfdffg [get]
export const findDfdfdffg = (params) => {
  return service({
    url: '/dfdfdffg/findDfdfdffg',
    method: 'get',
    params
  })
}

// @Tags Dfdfdffg
// @Summary 分页获取Dfdfdffg列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Dfdfdffg列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dfdfdffg/getDfdfdffgList [get]
export const getDfdfdffgList = (params) => {
  return service({
    url: '/dfdfdffg/getDfdfdffgList',
    method: 'get',
    params
  })
}
