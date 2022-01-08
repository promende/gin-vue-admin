package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ContractRouter struct {
}

// InitContractRouter 初始化 Contract 路由信息
func (s *ContractRouter) InitContractRouter(Router *gin.RouterGroup) {
	pactRouter := Router.Group("pact").Use(middleware.OperationRecord())
	pactRouterWithoutRecord := Router.Group("pact")
	var pactApi = v1.ApiGroupApp.AutoCodeApiGroup.ContractApi
	{
		pactRouter.POST("createContract", pactApi.CreateContract)   // 新建Contract
		pactRouter.DELETE("deleteContract", pactApi.DeleteContract) // 删除Contract
		pactRouter.DELETE("deleteContractByIds", pactApi.DeleteContractByIds) // 批量删除Contract
		pactRouter.PUT("updateContract", pactApi.UpdateContract)    // 更新Contract
	}
	{
		pactRouterWithoutRecord.GET("findContract", pactApi.FindContract)        // 根据ID获取Contract
		pactRouterWithoutRecord.GET("getContractList", pactApi.GetContractList)  // 获取Contract列表
	}
}
