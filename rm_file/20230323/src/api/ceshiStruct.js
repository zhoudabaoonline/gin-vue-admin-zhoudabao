import service from '@/utils/request'

// @Tags CeshiStruct
// @Summary 创建CeshiStruct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CeshiStruct true "创建CeshiStruct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ceshiStruct/createCeshiStruct [post]
export const createCeshiStruct = (data) => {
  return service({
    url: '/ceshiStruct/createCeshiStruct',
    method: 'post',
    data
  })
}

// @Tags CeshiStruct
// @Summary 删除CeshiStruct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CeshiStruct true "删除CeshiStruct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ceshiStruct/deleteCeshiStruct [delete]
export const deleteCeshiStruct = (data) => {
  return service({
    url: '/ceshiStruct/deleteCeshiStruct',
    method: 'delete',
    data
  })
}

// @Tags CeshiStruct
// @Summary 删除CeshiStruct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CeshiStruct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ceshiStruct/deleteCeshiStruct [delete]
export const deleteCeshiStructByIds = (data) => {
  return service({
    url: '/ceshiStruct/deleteCeshiStructByIds',
    method: 'delete',
    data
  })
}

// @Tags CeshiStruct
// @Summary 更新CeshiStruct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CeshiStruct true "更新CeshiStruct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ceshiStruct/updateCeshiStruct [put]
export const updateCeshiStruct = (data) => {
  return service({
    url: '/ceshiStruct/updateCeshiStruct',
    method: 'put',
    data
  })
}

// @Tags CeshiStruct
// @Summary 用id查询CeshiStruct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CeshiStruct true "用id查询CeshiStruct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ceshiStruct/findCeshiStruct [get]
export const findCeshiStruct = (params) => {
  return service({
    url: '/ceshiStruct/findCeshiStruct',
    method: 'get',
    params
  })
}

// @Tags CeshiStruct
// @Summary 分页获取CeshiStruct列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取CeshiStruct列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ceshiStruct/getCeshiStructList [get]
export const getCeshiStructList = (params) => {
  return service({
    url: '/ceshiStruct/getCeshiStructList',
    method: 'get',
    params
  })
}
