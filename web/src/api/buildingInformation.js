import service from '@/utils/request'

// @Tags BuildingInformation
// @Summary 创建BuildingInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BuildingInformation true "创建BuildingInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /buildingInformation/createBuildingInformation [post]
export const createBuildingInformation = (data) => {
  return service({
    url: '/buildingInformation/createBuildingInformation',
    method: 'post',
    data
  })
}

// @Tags BuildingInformation
// @Summary 删除BuildingInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BuildingInformation true "删除BuildingInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /buildingInformation/deleteBuildingInformation [delete]
export const deleteBuildingInformation = (data) => {
  return service({
    url: '/buildingInformation/deleteBuildingInformation',
    method: 'delete',
    data
  })
}

// @Tags BuildingInformation
// @Summary 删除BuildingInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BuildingInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /buildingInformation/deleteBuildingInformation [delete]
export const deleteBuildingInformationByIds = (data) => {
  return service({
    url: '/buildingInformation/deleteBuildingInformationByIds',
    method: 'delete',
    data
  })
}

// @Tags BuildingInformation
// @Summary 更新BuildingInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BuildingInformation true "更新BuildingInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /buildingInformation/updateBuildingInformation [put]
export const updateBuildingInformation = (data) => {
  return service({
    url: '/buildingInformation/updateBuildingInformation',
    method: 'put',
    data
  })
}

// @Tags BuildingInformation
// @Summary 用id查询BuildingInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BuildingInformation true "用id查询BuildingInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /buildingInformation/findBuildingInformation [get]
export const findBuildingInformation = (params) => {
  return service({
    url: '/buildingInformation/findBuildingInformation',
    method: 'get',
    params
  })
}

// @Tags BuildingInformation
// @Summary 分页获取BuildingInformation列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取BuildingInformation列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /buildingInformation/getBuildingInformationList [get]
export const getBuildingInformationList = (params) => {
  return service({
    url: '/buildingInformation/getBuildingInformationList',
    method: 'get',
    params
  })
}
