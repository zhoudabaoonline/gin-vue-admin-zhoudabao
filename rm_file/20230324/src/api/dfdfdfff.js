import service from '@/utils/request'

// @Tags Dfdfdffdf
// @Summary 创建Dfdfdffdf
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Dfdfdffdf true "创建Dfdfdffdf"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dfdfdfsss/createDfdfdffdf [post]
export const createDfdfdffdf = (data) => {
  return service({
    url: '/dfdfdfsss/createDfdfdffdf',
    method: 'post',
    data
  })
}

// @Tags Dfdfdffdf
// @Summary 删除Dfdfdffdf
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Dfdfdffdf true "删除Dfdfdffdf"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dfdfdfsss/deleteDfdfdffdf [delete]
export const deleteDfdfdffdf = (data) => {
  return service({
    url: '/dfdfdfsss/deleteDfdfdffdf',
    method: 'delete',
    data
  })
}

// @Tags Dfdfdffdf
// @Summary 删除Dfdfdffdf
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Dfdfdffdf"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dfdfdfsss/deleteDfdfdffdf [delete]
export const deleteDfdfdffdfByIds = (data) => {
  return service({
    url: '/dfdfdfsss/deleteDfdfdffdfByIds',
    method: 'delete',
    data
  })
}

// @Tags Dfdfdffdf
// @Summary 更新Dfdfdffdf
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Dfdfdffdf true "更新Dfdfdffdf"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dfdfdfsss/updateDfdfdffdf [put]
export const updateDfdfdffdf = (data) => {
  return service({
    url: '/dfdfdfsss/updateDfdfdffdf',
    method: 'put',
    data
  })
}

// @Tags Dfdfdffdf
// @Summary 用id查询Dfdfdffdf
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Dfdfdffdf true "用id查询Dfdfdffdf"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dfdfdfsss/findDfdfdffdf [get]
export const findDfdfdffdf = (params) => {
  return service({
    url: '/dfdfdfsss/findDfdfdffdf',
    method: 'get',
    params
  })
}

// @Tags Dfdfdffdf
// @Summary 分页获取Dfdfdffdf列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Dfdfdffdf列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dfdfdfsss/getDfdfdffdfList [get]
export const getDfdfdffdfList = (params) => {
  return service({
    url: '/dfdfdfsss/getDfdfdffdfList',
    method: 'get',
    params
  })
}
