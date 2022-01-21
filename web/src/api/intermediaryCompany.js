import service from '@/utils/request'

// @Tags IntermediaryCompany
// @Summary 创建IntermediaryCompany
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IntermediaryCompany true "创建IntermediaryCompany"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /intermediaryCompany/createIntermediaryCompany [post]
export const createIntermediaryCompany = (data) => {
  return service({
    url: '/intermediaryCompany/createIntermediaryCompany',
    method: 'post',
    data
  })
}

// @Tags IntermediaryCompany
// @Summary 删除IntermediaryCompany
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IntermediaryCompany true "删除IntermediaryCompany"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /intermediaryCompany/deleteIntermediaryCompany [delete]
export const deleteIntermediaryCompany = (data) => {
  return service({
    url: '/intermediaryCompany/deleteIntermediaryCompany',
    method: 'delete',
    data
  })
}

// @Tags IntermediaryCompany
// @Summary 删除IntermediaryCompany
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除IntermediaryCompany"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /intermediaryCompany/deleteIntermediaryCompany [delete]
export const deleteIntermediaryCompanyByIds = (data) => {
  return service({
    url: '/intermediaryCompany/deleteIntermediaryCompanyByIds',
    method: 'delete',
    data
  })
}

// @Tags IntermediaryCompany
// @Summary 更新IntermediaryCompany
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IntermediaryCompany true "更新IntermediaryCompany"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /intermediaryCompany/updateIntermediaryCompany [put]
export const updateIntermediaryCompany = (data) => {
  return service({
    url: '/intermediaryCompany/updateIntermediaryCompany',
    method: 'put',
    data
  })
}

// @Tags IntermediaryCompany
// @Summary 用id查询IntermediaryCompany
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.IntermediaryCompany true "用id查询IntermediaryCompany"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /intermediaryCompany/findIntermediaryCompany [get]
export const findIntermediaryCompany = (params) => {
  return service({
    url: '/intermediaryCompany/findIntermediaryCompany',
    method: 'get',
    params
  })
}

// @Tags IntermediaryCompany
// @Summary 分页获取IntermediaryCompany列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取IntermediaryCompany列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /intermediaryCompany/getIntermediaryCompanyList [get]
export const getIntermediaryCompanyList = (params) => {
  return service({
    url: '/intermediaryCompany/getIntermediaryCompanyList',
    method: 'get',
    params
  })
}
