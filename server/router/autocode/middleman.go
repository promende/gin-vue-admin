package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MiddlemanRouter struct {
}

// InitMiddlemanRouter 初始化 Middleman 路由信息
func (s *MiddlemanRouter) InitMiddlemanRouter(Router *gin.RouterGroup) {
	middlemanRouter := Router.Group("middleman").Use(middleware.OperationRecord())
	middlemanRouterWithoutRecord := Router.Group("middleman")
	var middlemanApi = v1.ApiGroupApp.AutoCodeApiGroup.MiddlemanApi
	{
		middlemanRouter.POST("createMiddleman", middlemanApi.CreateMiddleman)   // 新建Middleman
		middlemanRouter.DELETE("deleteMiddleman", middlemanApi.DeleteMiddleman) // 删除Middleman
		middlemanRouter.DELETE("deleteMiddlemanByIds", middlemanApi.DeleteMiddlemanByIds) // 批量删除Middleman
		middlemanRouter.PUT("updateMiddleman", middlemanApi.UpdateMiddleman)    // 更新Middleman
	}
	{
		middlemanRouterWithoutRecord.GET("findMiddleman", middlemanApi.FindMiddleman)        // 根据ID获取Middleman
		middlemanRouterWithoutRecord.GET("getMiddlemanList", middlemanApi.GetMiddlemanList)  // 获取Middleman列表
	}
}
