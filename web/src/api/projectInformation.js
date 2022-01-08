import service from '@/utils/request'

// @Tags ProjectInformation
// @Summary 创建ProjectInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ProjectInformation true "创建ProjectInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /project/createProjectInformation [post]
export const createProjectInformation = (data) => {
  return service({
    url: '/project/createProjectInformation',
    method: 'post',
    data
  })
}

// @Tags ProjectInformation
// @Summary 删除ProjectInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ProjectInformation true "删除ProjectInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /project/deleteProjectInformation [delete]
export const deleteProjectInformation = (data) => {
  return service({
    url: '/project/deleteProjectInformation',
    method: 'delete',
    data
  })
}

// @Tags ProjectInformation
// @Summary 删除ProjectInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ProjectInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /project/deleteProjectInformation [delete]
export const deleteProjectInformationByIds = (data) => {
  return service({
    url: '/project/deleteProjectInformationByIds',
    method: 'delete',
    data
  })
}

// @Tags ProjectInformation
// @Summary 更新ProjectInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ProjectInformation true "更新ProjectInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /project/updateProjectInformation [put]
export const updateProjectInformation = (data) => {
  return service({
    url: '/project/updateProjectInformation',
    method: 'put',
    data
  })
}

// @Tags ProjectInformation
// @Summary 用id查询ProjectInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ProjectInformation true "用id查询ProjectInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /project/findProjectInformation [get]
export const findProjectInformation = (params) => {
  return service({
    url: '/project/findProjectInformation',
    method: 'get',
    params
  })
}

// @Tags ProjectInformation
// @Summary 分页获取ProjectInformation列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ProjectInformation列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /project/getProjectInformationList [get]
export const getProjectInformationList = (params) => {
  return service({
    url: '/project/getProjectInformationList',
    method: 'get',
    params
  })
}
