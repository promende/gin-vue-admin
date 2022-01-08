package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HousingMaintenanceRouter struct {
}

// InitHousingMaintenanceRouter 初始化 HousingMaintenance 路由信息
func (s *HousingMaintenanceRouter) InitHousingMaintenanceRouter(Router *gin.RouterGroup) {
	housingMaintenanceRouter := Router.Group("housingMaintenance").Use(middleware.OperationRecord())
	housingMaintenanceRouterWithoutRecord := Router.Group("housingMaintenance")
	var housingMaintenanceApi = v1.ApiGroupApp.AutoCodeApiGroup.HousingMaintenanceApi
	{
		housingMaintenanceRouter.POST("createHousingMaintenance", housingMaintenanceApi.CreateHousingMaintenance)   // 新建HousingMaintenance
		housingMaintenanceRouter.DELETE("deleteHousingMaintenance", housingMaintenanceApi.DeleteHousingMaintenance) // 删除HousingMaintenance
		housingMaintenanceRouter.DELETE("deleteHousingMaintenanceByIds", housingMaintenanceApi.DeleteHousingMaintenanceByIds) // 批量删除HousingMaintenance
		housingMaintenanceRouter.PUT("updateHousingMaintenance", housingMaintenanceApi.UpdateHousingMaintenance)    // 更新HousingMaintenance
	}
	{
		housingMaintenanceRouterWithoutRecord.GET("findHousingMaintenance", housingMaintenanceApi.FindHousingMaintenance)        // 根据ID获取HousingMaintenance
		housingMaintenanceRouterWithoutRecord.GET("getHousingMaintenanceList", housingMaintenanceApi.GetHousingMaintenanceList)  // 获取HousingMaintenance列表
	}
}
