package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PactRouter struct {
}

// InitPactRouter 初始化 Pact 路由信息
func (s *PactRouter) InitPactRouter(Router *gin.RouterGroup) {
	ptRouter := Router.Group("pt").Use(middleware.OperationRecord())
	ptRouterWithoutRecord := Router.Group("pt")
	var ptApi = v1.ApiGroupApp.AutoCodeApiGroup.PactApi
	{
		ptRouter.POST("createPact", ptApi.CreatePact)   // 新建Pact
		ptRouter.DELETE("deletePact", ptApi.DeletePact) // 删除Pact
		ptRouter.DELETE("deletePactByIds", ptApi.DeletePactByIds) // 批量删除Pact
		ptRouter.PUT("updatePact", ptApi.UpdatePact)    // 更新Pact
	}
	{
		ptRouterWithoutRecord.GET("findPact", ptApi.FindPact)        // 根据ID获取Pact
		ptRouterWithoutRecord.GET("getPactList", ptApi.GetPactList)  // 获取Pact列表
	}
}
