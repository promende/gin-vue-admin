package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ProjectInformationRouter struct {
}

// InitProjectInformationRouter 初始化 ProjectInformation 路由信息
func (s *ProjectInformationRouter) InitProjectInformationRouter(Router *gin.RouterGroup) {
	projectRouter := Router.Group("project").Use(middleware.OperationRecord())
	projectRouterWithoutRecord := Router.Group("project")
	var projectApi = v1.ApiGroupApp.AutoCodeApiGroup.ProjectInformationApi
	{
		projectRouter.POST("createProjectInformation", projectApi.CreateProjectInformation)   // 新建ProjectInformation
		projectRouter.DELETE("deleteProjectInformation", projectApi.DeleteProjectInformation) // 删除ProjectInformation
		projectRouter.DELETE("deleteProjectInformationByIds", projectApi.DeleteProjectInformationByIds) // 批量删除ProjectInformation
		projectRouter.PUT("updateProjectInformation", projectApi.UpdateProjectInformation)    // 更新ProjectInformation
	}
	{
		projectRouterWithoutRecord.GET("findProjectInformation", projectApi.FindProjectInformation)        // 根据ID获取ProjectInformation
		projectRouterWithoutRecord.GET("getProjectInformationList", projectApi.GetProjectInformationList)  // 获取ProjectInformation列表
	}
}
