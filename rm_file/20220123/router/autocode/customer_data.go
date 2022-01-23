package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CustomerRouter struct {
}

// InitCustomerRouter 初始化 Customer 路由信息
func (s *CustomerRouter) InitCustomerRouter(Router *gin.RouterGroup) {
	customerDaRouter := Router.Group("customerDa").Use(middleware.OperationRecord())
	customerDaRouterWithoutRecord := Router.Group("customerDa")
	var customerDaApi = v1.ApiGroupApp.AutoCodeApiGroup.CustomerApi
	{
		customerDaRouter.POST("createCustomer", customerDaApi.CreateCustomer)   // 新建Customer
		customerDaRouter.DELETE("deleteCustomer", customerDaApi.DeleteCustomer) // 删除Customer
		customerDaRouter.DELETE("deleteCustomerByIds", customerDaApi.DeleteCustomerByIds) // 批量删除Customer
		customerDaRouter.PUT("updateCustomer", customerDaApi.UpdateCustomer)    // 更新Customer
	}
	{
		customerDaRouterWithoutRecord.GET("findCustomer", customerDaApi.FindCustomer)        // 根据ID获取Customer
		customerDaRouterWithoutRecord.GET("getCustomerList", customerDaApi.GetCustomerList)  // 获取Customer列表
	}
}
