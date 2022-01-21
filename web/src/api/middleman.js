import service from '@/utils/request'

// @Tags Middleman
// @Summary 创建Middleman
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Middleman true "创建Middleman"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /middleman/createMiddleman [post]
export const createMiddleman = (data) => {
  return service({
    url: '/middleman/createMiddleman',
    method: 'post',
    data
  })
}

// @Tags Middleman
// @Summary 删除Middleman
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Middleman true "删除Middleman"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /middleman/deleteMiddleman [delete]
export const deleteMiddleman = (data) => {
  return service({
    url: '/middleman/deleteMiddleman',
    method: 'delete',
    data
  })
}

// @Tags Middleman
// @Summary 删除Middleman
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Middleman"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /middleman/deleteMiddleman [delete]
export const deleteMiddlemanByIds = (data) => {
  return service({
    url: '/middleman/deleteMiddlemanByIds',
    method: 'delete',
    data
  })
}

// @Tags Middleman
// @Summary 更新Middleman
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Middleman true "更新Middleman"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /middleman/updateMiddleman [put]
export const updateMiddleman = (data) => {
  return service({
    url: '/middleman/updateMiddleman',
    method: 'put',
    data
  })
}

// @Tags Middleman
// @Summary 用id查询Middleman
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Middleman true "用id查询Middleman"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /middleman/findMiddleman [get]
export const findMiddleman = (params) => {
  return service({
    url: '/middleman/findMiddleman',
    method: 'get',
    params
  })
}

// @Tags Middleman
// @Summary 分页获取Middleman列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Middleman列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /middleman/getMiddlemanList [get]
export const getMiddlemanList = (params) => {
  return service({
    url: '/middleman/getMiddlemanList',
    method: 'get',
    params
  })
}
