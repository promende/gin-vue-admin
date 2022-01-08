import service from '@/utils/request'

// @Tags HousingMaintenance
// @Summary 创建HousingMaintenance
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HousingMaintenance true "创建HousingMaintenance"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /housingMaintenance/createHousingMaintenance [post]
export const createHousingMaintenance = (data) => {
  return service({
    url: '/housingMaintenance/createHousingMaintenance',
    method: 'post',
    data
  })
}

// @Tags HousingMaintenance
// @Summary 删除HousingMaintenance
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HousingMaintenance true "删除HousingMaintenance"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /housingMaintenance/deleteHousingMaintenance [delete]
export const deleteHousingMaintenance = (data) => {
  return service({
    url: '/housingMaintenance/deleteHousingMaintenance',
    method: 'delete',
    data
  })
}

// @Tags HousingMaintenance
// @Summary 删除HousingMaintenance
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HousingMaintenance"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /housingMaintenance/deleteHousingMaintenance [delete]
export const deleteHousingMaintenanceByIds = (data) => {
  return service({
    url: '/housingMaintenance/deleteHousingMaintenanceByIds',
    method: 'delete',
    data
  })
}

// @Tags HousingMaintenance
// @Summary 更新HousingMaintenance
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HousingMaintenance true "更新HousingMaintenance"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /housingMaintenance/updateHousingMaintenance [put]
export const updateHousingMaintenance = (data) => {
  return service({
    url: '/housingMaintenance/updateHousingMaintenance',
    method: 'put',
    data
  })
}

// @Tags HousingMaintenance
// @Summary 用id查询HousingMaintenance
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HousingMaintenance true "用id查询HousingMaintenance"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /housingMaintenance/findHousingMaintenance [get]
export const findHousingMaintenance = (params) => {
  return service({
    url: '/housingMaintenance/findHousingMaintenance',
    method: 'get',
    params
  })
}

// @Tags HousingMaintenance
// @Summary 分页获取HousingMaintenance列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HousingMaintenance列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /housingMaintenance/getHousingMaintenanceList [get]
export const getHousingMaintenanceList = (params) => {
  return service({
    url: '/housingMaintenance/getHousingMaintenanceList',
    method: 'get',
    params
  })
}
