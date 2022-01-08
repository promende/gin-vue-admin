import service from '@/utils/request'

// @Tags IntermediaryContact
// @Summary 创建IntermediaryContact
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IntermediaryContact true "创建IntermediaryContact"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /intermediarycontact/createIntermediaryContact [post]
export const createIntermediaryContact = (data) => {
  return service({
    url: '/intermediarycontact/createIntermediaryContact',
    method: 'post',
    data
  })
}

// @Tags IntermediaryContact
// @Summary 删除IntermediaryContact
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IntermediaryContact true "删除IntermediaryContact"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /intermediarycontact/deleteIntermediaryContact [delete]
export const deleteIntermediaryContact = (data) => {
  return service({
    url: '/intermediarycontact/deleteIntermediaryContact',
    method: 'delete',
    data
  })
}

// @Tags IntermediaryContact
// @Summary 删除IntermediaryContact
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除IntermediaryContact"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /intermediarycontact/deleteIntermediaryContact [delete]
export const deleteIntermediaryContactByIds = (data) => {
  return service({
    url: '/intermediarycontact/deleteIntermediaryContactByIds',
    method: 'delete',
    data
  })
}

// @Tags IntermediaryContact
// @Summary 更新IntermediaryContact
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IntermediaryContact true "更新IntermediaryContact"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /intermediarycontact/updateIntermediaryContact [put]
export const updateIntermediaryContact = (data) => {
  return service({
    url: '/intermediarycontact/updateIntermediaryContact',
    method: 'put',
    data
  })
}

// @Tags IntermediaryContact
// @Summary 用id查询IntermediaryContact
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.IntermediaryContact true "用id查询IntermediaryContact"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /intermediarycontact/findIntermediaryContact [get]
export const findIntermediaryContact = (params) => {
  return service({
    url: '/intermediarycontact/findIntermediaryContact',
    method: 'get',
    params
  })
}

// @Tags IntermediaryContact
// @Summary 分页获取IntermediaryContact列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取IntermediaryContact列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /intermediarycontact/getIntermediaryContactList [get]
export const getIntermediaryContactList = (params) => {
  return service({
    url: '/intermediarycontact/getIntermediaryContactList',
    method: 'get',
    params
  })
}

export const TestT = (params) => {
  return service({
    url: '/intermediarycontact/TestT',
    method: 'post',
    params
  })
}