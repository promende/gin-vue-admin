package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type IntermediaryRouter struct {
}

// InitIntermediaryRouter 初始化 Intermediary 路由信息
func (s *IntermediaryRouter) InitIntermediaryRouter(Router *gin.RouterGroup) {
	intermediaryRouter := Router.Group("intermediary").Use(middleware.OperationRecord())
	intermediaryRouterWithoutRecord := Router.Group("intermediary")
	var intermediaryApi = v1.ApiGroupApp.AutoCodeApiGroup.IntermediaryApi
	{
		intermediaryRouter.POST("createIntermediary", intermediaryApi.CreateIntermediary)   // 新建Intermediary
		intermediaryRouter.DELETE("deleteIntermediary", intermediaryApi.DeleteIntermediary) // 删除Intermediary
		intermediaryRouter.DELETE("deleteIntermediaryByIds", intermediaryApi.DeleteIntermediaryByIds) // 批量删除Intermediary
		intermediaryRouter.PUT("updateIntermediary", intermediaryApi.UpdateIntermediary)    // 更新Intermediary
	}
	{
		intermediaryRouterWithoutRecord.GET("findIntermediary", intermediaryApi.FindIntermediary)        // 根据ID获取Intermediary
		intermediaryRouterWithoutRecord.GET("getIntermediaryList", intermediaryApi.GetIntermediaryList)  // 获取Intermediary列表
	}
}
