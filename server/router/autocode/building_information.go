package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BuildingInformationRouter struct {
}

// InitBuildingInformationRouter 初始化 BuildingInformation 路由信息
func (s *BuildingInformationRouter) InitBuildingInformationRouter(Router *gin.RouterGroup) {
	buildingInformationRouter := Router.Group("buildingInformation").Use(middleware.OperationRecord())
	buildingInformationRouterWithoutRecord := Router.Group("buildingInformation")
	var buildingInformationApi = v1.ApiGroupApp.AutoCodeApiGroup.BuildingInformationApi
	{
		buildingInformationRouter.POST("createBuildingInformation", buildingInformationApi.CreateBuildingInformation)   // 新建BuildingInformation
		buildingInformationRouter.DELETE("deleteBuildingInformation", buildingInformationApi.DeleteBuildingInformation) // 删除BuildingInformation
		buildingInformationRouter.DELETE("deleteBuildingInformationByIds", buildingInformationApi.DeleteBuildingInformationByIds) // 批量删除BuildingInformation
		buildingInformationRouter.PUT("updateBuildingInformation", buildingInformationApi.UpdateBuildingInformation)    // 更新BuildingInformation
	}
	{
		buildingInformationRouterWithoutRecord.GET("findBuildingInformation", buildingInformationApi.FindBuildingInformation)        // 根据ID获取BuildingInformation
		buildingInformationRouterWithoutRecord.GET("getBuildingInformationList", buildingInformationApi.GetBuildingInformationList)  // 获取BuildingInformation列表
	}
}
