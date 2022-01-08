package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BusinessRouter struct {
}

// InitBusinessRouter 初始化 Business 路由信息
func (s *BusinessRouter) InitBusinessRouter(Router *gin.RouterGroup) {
	businessRouter := Router.Group("business").Use(middleware.OperationRecord())
	businessRouterWithoutRecord := Router.Group("business")
	var businessApi = v1.ApiGroupApp.AutoCodeApiGroup.BusinessApi
	{
		businessRouter.POST("createBusiness", businessApi.CreateBusiness)   // 新建Business
		businessRouter.DELETE("deleteBusiness", businessApi.DeleteBusiness) // 删除Business
		businessRouter.DELETE("deleteBusinessByIds", businessApi.DeleteBusinessByIds) // 批量删除Business
		businessRouter.PUT("updateBusiness", businessApi.UpdateBusiness)    // 更新Business
	}
	{
		businessRouterWithoutRecord.GET("findBusiness", businessApi.FindBusiness)        // 根据ID获取Business
		businessRouterWithoutRecord.GET("getBusinessList", businessApi.GetBusinessList)  // 获取Business列表
	}
}
