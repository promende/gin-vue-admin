package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type FloorInformationRouter struct {
}

// InitFloorInformationRouter 初始化 FloorInformation 路由信息
func (s *FloorInformationRouter) InitFloorInformationRouter(Router *gin.RouterGroup) {
	floorInformationRouter := Router.Group("floorInformation").Use(middleware.OperationRecord())
	floorInformationRouterWithoutRecord := Router.Group("floorInformation")
	var floorInformationApi = v1.ApiGroupApp.AutoCodeApiGroup.FloorInformationApi
	{
		floorInformationRouter.POST("createFloorInformation", floorInformationApi.CreateFloorInformation)   // 新建FloorInformation
		floorInformationRouter.DELETE("deleteFloorInformation", floorInformationApi.DeleteFloorInformation) // 删除FloorInformation
		floorInformationRouter.DELETE("deleteFloorInformationByIds", floorInformationApi.DeleteFloorInformationByIds) // 批量删除FloorInformation
		floorInformationRouter.PUT("updateFloorInformation", floorInformationApi.UpdateFloorInformation)    // 更新FloorInformation
	}
	{
		floorInformationRouterWithoutRecord.GET("findFloorInformation", floorInformationApi.FindFloorInformation)        // 根据ID获取FloorInformation
		floorInformationRouterWithoutRecord.GET("getFloorInformationList", floorInformationApi.GetFloorInformationList)  // 获取FloorInformation列表
	}
}
