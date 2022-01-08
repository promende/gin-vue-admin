import service from '@/utils/request'

// @Tags FloorInformation
// @Summary 创建FloorInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FloorInformation true "创建FloorInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /floorInformation/createFloorInformation [post]
export const createFloorInformation = (data) => {
  return service({
    url: '/floorInformation/createFloorInformation',
    method: 'post',
    data
  })
}

// @Tags FloorInformation
// @Summary 删除FloorInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FloorInformation true "删除FloorInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /floorInformation/deleteFloorInformation [delete]
export const deleteFloorInformation = (data) => {
  return service({
    url: '/floorInformation/deleteFloorInformation',
    method: 'delete',
    data
  })
}

// @Tags FloorInformation
// @Summary 删除FloorInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除FloorInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /floorInformation/deleteFloorInformation [delete]
export const deleteFloorInformationByIds = (data) => {
  return service({
    url: '/floorInformation/deleteFloorInformationByIds',
    method: 'delete',
    data
  })
}

// @Tags FloorInformation
// @Summary 更新FloorInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FloorInformation true "更新FloorInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /floorInformation/updateFloorInformation [put]
export const updateFloorInformation = (data) => {
  return service({
    url: '/floorInformation/updateFloorInformation',
    method: 'put',
    data
  })
}

// @Tags FloorInformation
// @Summary 用id查询FloorInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.FloorInformation true "用id查询FloorInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /floorInformation/findFloorInformation [get]
export const findFloorInformation = (params) => {
  return service({
    url: '/floorInformation/findFloorInformation',
    method: 'get',
    params
  })
}

// @Tags FloorInformation
// @Summary 分页获取FloorInformation列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取FloorInformation列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /floorInformation/getFloorInformationList [get]
export const getFloorInformationList = (params) => {
  return service({
    url: '/floorInformation/getFloorInformationList',
    method: 'get',
    params
  })
}
