package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autocodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type FloorInformationApi struct {
}

var floorInformationService = service.ServiceGroupApp.AutoCodeServiceGroup.FloorInformationService


// CreateFloorInformation 创建FloorInformation
// @Tags FloorInformation
// @Summary 创建FloorInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.FloorInformation true "创建FloorInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /floorInformation/createFloorInformation [post]
func (floorInformationApi *FloorInformationApi) CreateFloorInformation(c *gin.Context) {
	var floorInformation autocode.FloorInformation
	_ = c.ShouldBindJSON(&floorInformation)
	if err := floorInformationService.CreateFloorInformation(floorInformation); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteFloorInformation 删除FloorInformation
// @Tags FloorInformation
// @Summary 删除FloorInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.FloorInformation true "删除FloorInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /floorInformation/deleteFloorInformation [delete]
func (floorInformationApi *FloorInformationApi) DeleteFloorInformation(c *gin.Context) {
	var floorInformation autocode.FloorInformation
	_ = c.ShouldBindJSON(&floorInformation)
	if err := floorInformationService.DeleteFloorInformation(floorInformation); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteFloorInformationByIds 批量删除FloorInformation
// @Tags FloorInformation
// @Summary 批量删除FloorInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除FloorInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /floorInformation/deleteFloorInformationByIds [delete]
func (floorInformationApi *FloorInformationApi) DeleteFloorInformationByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := floorInformationService.DeleteFloorInformationByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateFloorInformation 更新FloorInformation
// @Tags FloorInformation
// @Summary 更新FloorInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.FloorInformation true "更新FloorInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /floorInformation/updateFloorInformation [put]
func (floorInformationApi *FloorInformationApi) UpdateFloorInformation(c *gin.Context) {
	var floorInformation autocode.FloorInformation
	_ = c.ShouldBindJSON(&floorInformation)
	if err := floorInformationService.UpdateFloorInformation(floorInformation); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindFloorInformation 用id查询FloorInformation
// @Tags FloorInformation
// @Summary 用id查询FloorInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.FloorInformation true "用id查询FloorInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /floorInformation/findFloorInformation [get]
func (floorInformationApi *FloorInformationApi) FindFloorInformation(c *gin.Context) {
	var floorInformation autocode.FloorInformation
	_ = c.ShouldBindQuery(&floorInformation)
	if err, refloorInformation := floorInformationService.GetFloorInformation(floorInformation.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"refloorInformation": refloorInformation}, c)
	}
}

// GetFloorInformationList 分页获取FloorInformation列表
// @Tags FloorInformation
// @Summary 分页获取FloorInformation列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.FloorInformationSearch true "分页获取FloorInformation列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /floorInformation/getFloorInformationList [get]
func (floorInformationApi *FloorInformationApi) GetFloorInformationList(c *gin.Context) {
	var pageInfo autocodeReq.FloorInformationSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := floorInformationService.GetFloorInformationInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
