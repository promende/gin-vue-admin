import service from '@/utils/request'

// @Tags Pact
// @Summary 创建Pact
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Pact true "创建Pact"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pt/createPact [post]
export const createPact = (data) => {
  return service({
    url: '/pt/createPact',
    method: 'post',
    data
  })
}

// @Tags Pact
// @Summary 删除Pact
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Pact true "删除Pact"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pt/deletePact [delete]
export const deletePact = (data) => {
  return service({
    url: '/pt/deletePact',
    method: 'delete',
    data
  })
}

// @Tags Pact
// @Summary 删除Pact
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Pact"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pt/deletePact [delete]
export const deletePactByIds = (data) => {
  return service({
    url: '/pt/deletePactByIds',
    method: 'delete',
    data
  })
}

// @Tags Pact
// @Summary 更新Pact
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Pact true "更新Pact"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /pt/updatePact [put]
export const updatePact = (data) => {
  return service({
    url: '/pt/updatePact',
    method: 'put',
    data
  })
}

// @Tags Pact
// @Summary 用id查询Pact
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Pact true "用id查询Pact"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /pt/findPact [get]
export const findPact = (params) => {
  return service({
    url: '/pt/findPact',
    method: 'get',
    params
  })
}

// @Tags Pact
// @Summary 分页获取Pact列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Pact列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pt/getPactList [get]
export const getPactList = (params) => {
  return service({
    url: '/pt/getPactList',
    method: 'get',
    params
  })
}
