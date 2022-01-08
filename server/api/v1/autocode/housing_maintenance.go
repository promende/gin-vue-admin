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

type HousingMaintenanceApi struct {
}

var housingMaintenanceService = service.ServiceGroupApp.AutoCodeServiceGroup.HousingMaintenanceService


// CreateHousingMaintenance 创建HousingMaintenance
// @Tags HousingMaintenance
// @Summary 创建HousingMaintenance
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.HousingMaintenance true "创建HousingMaintenance"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /housingMaintenance/createHousingMaintenance [post]
func (housingMaintenanceApi *HousingMaintenanceApi) CreateHousingMaintenance(c *gin.Context) {
	var housingMaintenance autocode.HousingMaintenance
	_ = c.ShouldBindJSON(&housingMaintenance)
	if err := housingMaintenanceService.CreateHousingMaintenance(housingMaintenance); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHousingMaintenance 删除HousingMaintenance
// @Tags HousingMaintenance
// @Summary 删除HousingMaintenance
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.HousingMaintenance true "删除HousingMaintenance"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /housingMaintenance/deleteHousingMaintenance [delete]
func (housingMaintenanceApi *HousingMaintenanceApi) DeleteHousingMaintenance(c *gin.Context) {
	var housingMaintenance autocode.HousingMaintenance
	_ = c.ShouldBindJSON(&housingMaintenance)
	if err := housingMaintenanceService.DeleteHousingMaintenance(housingMaintenance); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHousingMaintenanceByIds 批量删除HousingMaintenance
// @Tags HousingMaintenance
// @Summary 批量删除HousingMaintenance
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HousingMaintenance"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /housingMaintenance/deleteHousingMaintenanceByIds [delete]
func (housingMaintenanceApi *HousingMaintenanceApi) DeleteHousingMaintenanceByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := housingMaintenanceService.DeleteHousingMaintenanceByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHousingMaintenance 更新HousingMaintenance
// @Tags HousingMaintenance
// @Summary 更新HousingMaintenance
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.HousingMaintenance true "更新HousingMaintenance"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /housingMaintenance/updateHousingMaintenance [put]
func (housingMaintenanceApi *HousingMaintenanceApi) UpdateHousingMaintenance(c *gin.Context) {
	var housingMaintenance autocode.HousingMaintenance
	_ = c.ShouldBindJSON(&housingMaintenance)
	if err := housingMaintenanceService.UpdateHousingMaintenance(housingMaintenance); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHousingMaintenance 用id查询HousingMaintenance
// @Tags HousingMaintenance
// @Summary 用id查询HousingMaintenance
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.HousingMaintenance true "用id查询HousingMaintenance"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /housingMaintenance/findHousingMaintenance [get]
func (housingMaintenanceApi *HousingMaintenanceApi) FindHousingMaintenance(c *gin.Context) {
	var housingMaintenance autocode.HousingMaintenance
	_ = c.ShouldBindQuery(&housingMaintenance)
	if err, rehousingMaintenance := housingMaintenanceService.GetHousingMaintenance(housingMaintenance.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehousingMaintenance": rehousingMaintenance}, c)
	}
}

// GetHousingMaintenanceList 分页获取HousingMaintenance列表
// @Tags HousingMaintenance
// @Summary 分页获取HousingMaintenance列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.HousingMaintenanceSearch true "分页获取HousingMaintenance列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /housingMaintenance/getHousingMaintenanceList [get]
func (housingMaintenanceApi *HousingMaintenanceApi) GetHousingMaintenanceList(c *gin.Context) {
	var pageInfo autocodeReq.HousingMaintenanceSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := housingMaintenanceService.GetHousingMaintenanceInfoList(pageInfo); err != nil {
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
