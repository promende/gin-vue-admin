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

type ProjectInformationApi struct {
}

var projectService = service.ServiceGroupApp.AutoCodeServiceGroup.ProjectInformationService


// CreateProjectInformation 创建ProjectInformation
// @Tags ProjectInformation
// @Summary 创建ProjectInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.ProjectInformation true "创建ProjectInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /project/createProjectInformation [post]
func (projectApi *ProjectInformationApi) CreateProjectInformation(c *gin.Context) {
	var project autocode.ProjectInformation
	_ = c.ShouldBindJSON(&project)
	if err := projectService.CreateProjectInformation(project); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteProjectInformation 删除ProjectInformation
// @Tags ProjectInformation
// @Summary 删除ProjectInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.ProjectInformation true "删除ProjectInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /project/deleteProjectInformation [delete]
func (projectApi *ProjectInformationApi) DeleteProjectInformation(c *gin.Context) {
	var project autocode.ProjectInformation
	_ = c.ShouldBindJSON(&project)
	if err := projectService.DeleteProjectInformation(project); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteProjectInformationByIds 批量删除ProjectInformation
// @Tags ProjectInformation
// @Summary 批量删除ProjectInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ProjectInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /project/deleteProjectInformationByIds [delete]
func (projectApi *ProjectInformationApi) DeleteProjectInformationByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := projectService.DeleteProjectInformationByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateProjectInformation 更新ProjectInformation
// @Tags ProjectInformation
// @Summary 更新ProjectInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.ProjectInformation true "更新ProjectInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /project/updateProjectInformation [put]
func (projectApi *ProjectInformationApi) UpdateProjectInformation(c *gin.Context) {
	var project autocode.ProjectInformation
	_ = c.ShouldBindJSON(&project)
	if err := projectService.UpdateProjectInformation(project); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindProjectInformation 用id查询ProjectInformation
// @Tags ProjectInformation
// @Summary 用id查询ProjectInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.ProjectInformation true "用id查询ProjectInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /project/findProjectInformation [get]
func (projectApi *ProjectInformationApi) FindProjectInformation(c *gin.Context) {
	var project autocode.ProjectInformation
	_ = c.ShouldBindQuery(&project)
	if err, reproject := projectService.GetProjectInformation(project.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reproject": reproject}, c)
	}
}

// GetProjectInformationList 分页获取ProjectInformation列表
// @Tags ProjectInformation
// @Summary 分页获取ProjectInformation列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.ProjectInformationSearch true "分页获取ProjectInformation列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /project/getProjectInformationList [get]
func (projectApi *ProjectInformationApi) GetProjectInformationList(c *gin.Context) {
	var pageInfo autocodeReq.ProjectInformationSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := projectService.GetProjectInformationInfoList(pageInfo); err != nil {
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
