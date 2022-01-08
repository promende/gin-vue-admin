package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type IntermediaryContactRouter struct {
}

// InitIntermediaryContactRouter 初始化 IntermediaryContact 路由信息
func (s *IntermediaryContactRouter) InitIntermediaryContactRouter(Router *gin.RouterGroup) {
	intermediarycontactRouter := Router.Group("intermediarycontact").Use(middleware.OperationRecord())
	intermediarycontactRouterWithoutRecord := Router.Group("intermediarycontact")
	var intermediarycontactApi = v1.ApiGroupApp.AutoCodeApiGroup.IntermediaryContactApi
	{
		intermediarycontactRouter.POST("createIntermediaryContact", intermediarycontactApi.CreateIntermediaryContact)   // 新建IntermediaryContact
		intermediarycontactRouter.DELETE("deleteIntermediaryContact", intermediarycontactApi.DeleteIntermediaryContact) // 删除IntermediaryContact
		intermediarycontactRouter.DELETE("deleteIntermediaryContactByIds", intermediarycontactApi.DeleteIntermediaryContactByIds) // 批量删除IntermediaryContact
		intermediarycontactRouter.PUT("updateIntermediaryContact", intermediarycontactApi.UpdateIntermediaryContact)    // 更新IntermediaryContact
	}
	{
		intermediarycontactRouterWithoutRecord.GET("findIntermediaryContact", intermediarycontactApi.FindIntermediaryContact)        // 根据ID获取IntermediaryContact
		intermediarycontactRouterWithoutRecord.GET("getIntermediaryContactList", intermediarycontactApi.GetIntermediaryContactList)  // 获取IntermediaryContact列表
	}
}
