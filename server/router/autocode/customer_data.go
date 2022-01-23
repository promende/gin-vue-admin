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
	customerDataRouter := Router.Group("customerData").Use(middleware.OperationRecord())
	customerDataRouterWithoutRecord := Router.Group("customerData")
	var customerDataApi = v1.ApiGroupApp.AutoCodeApiGroup.CustomerApi
	{
		customerDataRouter.POST("createCustomer", customerDataApi.CreateCustomer)   // 新建Customer
		customerDataRouter.DELETE("deleteCustomer", customerDataApi.DeleteCustomer) // 删除Customer
		customerDataRouter.DELETE("deleteCustomerByIds", customerDataApi.DeleteCustomerByIds) // 批量删除Customer
		customerDataRouter.PUT("updateCustomer", customerDataApi.UpdateCustomer)    // 更新Customer
	}
	{
		customerDataRouterWithoutRecord.GET("findCustomer", customerDataApi.FindCustomer)        // 根据ID获取Customer
		customerDataRouterWithoutRecord.GET("getCustomerList", customerDataApi.GetCustomerList)  // 获取Customer列表
	}
}
