import service from '@/utils/request'

// @Tags Strcut3
// @Summary 创建Strcut3
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Strcut3 true "创建Strcut3"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /strcut3/createStrcut3 [post]
export const createStrcut3 = (data) => {
  return service({
    url: '/strcut3/createStrcut3',
    method: 'post',
    data
  })
}

// @Tags Strcut3
// @Summary 删除Strcut3
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Strcut3 true "删除Strcut3"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /strcut3/deleteStrcut3 [delete]
export const deleteStrcut3 = (data) => {
  return service({
    url: '/strcut3/deleteStrcut3',
    method: 'delete',
    data
  })
}

// @Tags Strcut3
// @Summary 删除Strcut3
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Strcut3"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /strcut3/deleteStrcut3 [delete]
export const deleteStrcut3ByIds = (data) => {
  return service({
    url: '/strcut3/deleteStrcut3ByIds',
    method: 'delete',
    data
  })
}

// @Tags Strcut3
// @Summary 更新Strcut3
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Strcut3 true "更新Strcut3"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /strcut3/updateStrcut3 [put]
export const updateStrcut3 = (data) => {
  return service({
    url: '/strcut3/updateStrcut3',
    method: 'put',
    data
  })
}

// @Tags Strcut3
// @Summary 用id查询Strcut3
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Strcut3 true "用id查询Strcut3"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /strcut3/findStrcut3 [get]
export const findStrcut3 = (params) => {
  return service({
    url: '/strcut3/findStrcut3',
    method: 'get',
    params
  })
}

// @Tags Strcut3
// @Summary 分页获取Strcut3列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Strcut3列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /strcut3/getStrcut3List [get]
export const getStrcut3List = (params) => {
  return service({
    url: '/strcut3/getStrcut3List',
    method: 'get',
    params
  })
}
