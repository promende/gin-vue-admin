import service from '@/utils/request'

// @Tags Customer
// @Summary 创建Customer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Customer true "创建Customer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customerDa/createCustomer [post]
export const createCustomer = (data) => {
  return service({
    url: '/customerDa/createCustomer',
    method: 'post',
    data
  })
}

// @Tags Customer
// @Summary 删除Customer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Customer true "删除Customer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /customerDa/deleteCustomer [delete]
export const deleteCustomer = (data) => {
  return service({
    url: '/customerDa/deleteCustomer',
    method: 'delete',
    data
  })
}

// @Tags Customer
// @Summary 删除Customer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Customer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /customerDa/deleteCustomer [delete]
export const deleteCustomerByIds = (data) => {
  return service({
    url: '/customerDa/deleteCustomerByIds',
    method: 'delete',
    data
  })
}

// @Tags Customer
// @Summary 更新Customer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Customer true "更新Customer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /customerDa/updateCustomer [put]
export const updateCustomer = (data) => {
  return service({
    url: '/customerDa/updateCustomer',
    method: 'put',
    data
  })
}

// @Tags Customer
// @Summary 用id查询Customer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Customer true "用id查询Customer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /customerDa/findCustomer [get]
export const findCustomer = (params) => {
  return service({
    url: '/customerDa/findCustomer',
    method: 'get',
    params
  })
}

// @Tags Customer
// @Summary 分页获取Customer列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Customer列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customerDa/getCustomerList [get]
export const getCustomerList = (params) => {
  return service({
    url: '/customerDa/getCustomerList',
    method: 'get',
    params
  })
}
