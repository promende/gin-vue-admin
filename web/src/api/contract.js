import service from '@/utils/request'

// @Tags Contract
// @Summary 创建Contract
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Contract true "创建Contract"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pact/createContract [post]
export const createContract = (data) => {
  return service({
    url: '/pact/createContract',
    method: 'post',
    data
  })
}

// @Tags Contract
// @Summary 删除Contract
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Contract true "删除Contract"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pact/deleteContract [delete]
export const deleteContract = (data) => {
  return service({
    url: '/pact/deleteContract',
    method: 'delete',
    data
  })
}

// @Tags Contract
// @Summary 删除Contract
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Contract"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pact/deleteContract [delete]
export const deleteContractByIds = (data) => {
  return service({
    url: '/pact/deleteContractByIds',
    method: 'delete',
    data
  })
}

// @Tags Contract
// @Summary 更新Contract
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Contract true "更新Contract"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /pact/updateContract [put]
export const updateContract = (data) => {
  return service({
    url: '/pact/updateContract',
    method: 'put',
    data
  })
}

// @Tags Contract
// @Summary 用id查询Contract
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Contract true "用id查询Contract"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /pact/findContract [get]
export const findContract = (params) => {
  return service({
    url: '/pact/findContract',
    method: 'get',
    params
  })
}

// @Tags Contract
// @Summary 分页获取Contract列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Contract列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pact/getContractList [get]
export const getContractList = (params) => {
  return service({
    url: '/pact/getContractList',
    method: 'get',
    params
  })
}
