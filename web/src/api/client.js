import service from '@/utils/request'

// @Tags Client
// @Summary 创建Client
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Client true "创建Client"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /client/createClient [post]
export const createClient = (data) => {
  return service({
    url: '/client/createClient',
    method: 'post',
    data
  })
}

// @Tags Client
// @Summary 删除Client
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Client true "删除Client"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /client/deleteClient [delete]
export const deleteClient = (data) => {
  return service({
    url: '/client/deleteClient',
    method: 'delete',
    data
  })
}

// @Tags Client
// @Summary 删除Client
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Client"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /client/deleteClient [delete]
export const deleteClientByIds = (data) => {
  return service({
    url: '/client/deleteClientByIds',
    method: 'delete',
    data
  })
}

// @Tags Client
// @Summary 更新Client
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Client true "更新Client"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /client/updateClient [put]
export const updateClient = (data) => {
  return service({
    url: '/client/updateClient',
    method: 'put',
    data
  })
}

// @Tags Client
// @Summary 用id查询Client
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Client true "用id查询Client"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /client/findClient [get]
export const findClient = (params) => {
  return service({
    url: '/client/findClient',
    method: 'get',
    params
  })
}

// @Tags Client
// @Summary 分页获取Client列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Client列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /client/getClientList [get]
export const getClientList = (params) => {
  return service({
    url: '/client/getClientList',
    method: 'get',
    params
  })
}
