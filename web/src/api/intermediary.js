import service from '@/utils/request'

// @Tags Intermediary
// @Summary 创建Intermediary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Intermediary true "创建Intermediary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /intermediary/createIntermediary [post]
export const createIntermediary = (data) => {
  return service({
    url: '/intermediary/createIntermediary',
    method: 'post',
    data
  })
}

// @Tags Intermediary
// @Summary 删除Intermediary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Intermediary true "删除Intermediary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /intermediary/deleteIntermediary [delete]
export const deleteIntermediary = (data) => {
  return service({
    url: '/intermediary/deleteIntermediary',
    method: 'delete',
    data
  })
}

// @Tags Intermediary
// @Summary 删除Intermediary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Intermediary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /intermediary/deleteIntermediary [delete]
export const deleteIntermediaryByIds = (data) => {
  return service({
    url: '/intermediary/deleteIntermediaryByIds',
    method: 'delete',
    data
  })
}

// @Tags Intermediary
// @Summary 更新Intermediary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Intermediary true "更新Intermediary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /intermediary/updateIntermediary [put]
export const updateIntermediary = (data) => {
  return service({
    url: '/intermediary/updateIntermediary',
    method: 'put',
    data
  })
}

// @Tags Intermediary
// @Summary 用id查询Intermediary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Intermediary true "用id查询Intermediary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /intermediary/findIntermediary [get]
export const findIntermediary = (params) => {
  return service({
    url: '/intermediary/findIntermediary',
    method: 'get',
    params
  })
}

// @Tags Intermediary
// @Summary 分页获取Intermediary列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Intermediary列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /intermediary/getIntermediaryList [get]
export const getIntermediaryList = (params) => {
  return service({
    url: '/intermediary/getIntermediaryList',
    method: 'get',
    params
  })
}
