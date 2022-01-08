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

type BuildingInformationApi struct {
}

var buildingInformationService = service.ServiceGroupApp.AutoCodeServiceGroup.BuildingInformationService


// CreateBuildingInformation 创建BuildingInformation
// @Tags BuildingInformation
// @Summary 创建BuildingInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.BuildingInformation true "创建BuildingInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /buildingInformation/createBuildingInformation [post]
func (buildingInformationApi *BuildingInformationApi) CreateBuildingInformation(c *gin.Context) {
	var buildingInformation autocode.BuildingInformation
	_ = c.ShouldBindJSON(&buildingInformation)
	if err := buildingInformationService.CreateBuildingInformation(buildingInformation); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBuildingInformation 删除BuildingInformation
// @Tags BuildingInformation
// @Summary 删除BuildingInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.BuildingInformation true "删除BuildingInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /buildingInformation/deleteBuildingInformation [delete]
func (buildingInformationApi *BuildingInformationApi) DeleteBuildingInformation(c *gin.Context) {
	var buildingInformation autocode.BuildingInformation
	_ = c.ShouldBindJSON(&buildingInformation)
	if err := buildingInformationService.DeleteBuildingInformation(buildingInformation); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBuildingInformationByIds 批量删除BuildingInformation
// @Tags BuildingInformation
// @Summary 批量删除BuildingInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BuildingInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /buildingInformation/deleteBuildingInformationByIds [delete]
func (buildingInformationApi *BuildingInformationApi) DeleteBuildingInformationByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := buildingInformationService.DeleteBuildingInformationByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBuildingInformation 更新BuildingInformation
// @Tags BuildingInformation
// @Summary 更新BuildingInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.BuildingInformation true "更新BuildingInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /buildingInformation/updateBuildingInformation [put]
func (buildingInformationApi *BuildingInformationApi) UpdateBuildingInformation(c *gin.Context) {
	var buildingInformation autocode.BuildingInformation
	_ = c.ShouldBindJSON(&buildingInformation)
	if err := buildingInformationService.UpdateBuildingInformation(buildingInformation); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBuildingInformation 用id查询BuildingInformation
// @Tags BuildingInformation
// @Summary 用id查询BuildingInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.BuildingInformation true "用id查询BuildingInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /buildingInformation/findBuildingInformation [get]
func (buildingInformationApi *BuildingInformationApi) FindBuildingInformation(c *gin.Context) {
	var buildingInformation autocode.BuildingInformation
	_ = c.ShouldBindQuery(&buildingInformation)
	if err, rebuildingInformation := buildingInformationService.GetBuildingInformation(buildingInformation.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rebuildingInformation": rebuildingInformation}, c)
	}
}

// GetBuildingInformationList 分页获取BuildingInformation列表
// @Tags BuildingInformation
// @Summary 分页获取BuildingInformation列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.BuildingInformationSearch true "分页获取BuildingInformation列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /buildingInformation/getBuildingInformationList [get]
func (buildingInformationApi *BuildingInformationApi) GetBuildingInformationList(c *gin.Context) {
	var pageInfo autocodeReq.BuildingInformationSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := buildingInformationService.GetBuildingInformationInfoList(pageInfo); err != nil {
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
