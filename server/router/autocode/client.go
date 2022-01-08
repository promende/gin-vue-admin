package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ClientRouter struct {
}

// InitClientRouter 初始化 Client 路由信息
func (s *ClientRouter) InitClientRouter(Router *gin.RouterGroup) {
	clientRouter := Router.Group("client").Use(middleware.OperationRecord())
	clientRouterWithoutRecord := Router.Group("client")
	var clientApi = v1.ApiGroupApp.AutoCodeApiGroup.ClientApi
	{
		clientRouter.POST("createClient", clientApi.CreateClient)   // 新建Client
		clientRouter.DELETE("deleteClient", clientApi.DeleteClient) // 删除Client
		clientRouter.DELETE("deleteClientByIds", clientApi.DeleteClientByIds) // 批量删除Client
		clientRouter.PUT("updateClient", clientApi.UpdateClient)    // 更新Client
	}
	{
		clientRouterWithoutRecord.GET("findClient", clientApi.FindClient)        // 根据ID获取Client
		clientRouterWithoutRecord.GET("getClientList", clientApi.GetClientList)  // 获取Client列表
	}
}
