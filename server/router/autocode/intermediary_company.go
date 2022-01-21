package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type IntermediaryCompanyRouter struct {
}

// InitIntermediaryCompanyRouter 初始化 IntermediaryCompany 路由信息
func (s *IntermediaryCompanyRouter) InitIntermediaryCompanyRouter(Router *gin.RouterGroup) {
	intermediaryCompanyRouter := Router.Group("intermediaryCompany").Use(middleware.OperationRecord())
	intermediaryCompanyRouterWithoutRecord := Router.Group("intermediaryCompany")
	var intermediaryCompanyApi = v1.ApiGroupApp.AutoCodeApiGroup.IntermediaryCompanyApi
	{
		intermediaryCompanyRouter.POST("createIntermediaryCompany", intermediaryCompanyApi.CreateIntermediaryCompany)   // 新建IntermediaryCompany
		intermediaryCompanyRouter.DELETE("deleteIntermediaryCompany", intermediaryCompanyApi.DeleteIntermediaryCompany) // 删除IntermediaryCompany
		intermediaryCompanyRouter.DELETE("deleteIntermediaryCompanyByIds", intermediaryCompanyApi.DeleteIntermediaryCompanyByIds) // 批量删除IntermediaryCompany
		intermediaryCompanyRouter.PUT("updateIntermediaryCompany", intermediaryCompanyApi.UpdateIntermediaryCompany)    // 更新IntermediaryCompany
	}
	{
		intermediaryCompanyRouterWithoutRecord.GET("findIntermediaryCompany", intermediaryCompanyApi.FindIntermediaryCompany)        // 根据ID获取IntermediaryCompany
		intermediaryCompanyRouterWithoutRecord.GET("getIntermediaryCompanyList", intermediaryCompanyApi.GetIntermediaryCompanyList)  // 获取IntermediaryCompany列表
	}
}
