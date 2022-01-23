package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autocodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type CustomerApi struct {
}

var customerDaService = service.ServiceGroupApp.AutoCodeServiceGroup.CustomerService


// CreateCustomer 创建Customer
// @Tags Customer
// @Summary 创建Customer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Customer true "创建Customer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customerDa/createCustomer [post]
func (customerDaApi *CustomerApi) CreateCustomer(c *gin.Context) {
	var customerDa autocode.Customer
	_ = c.ShouldBindJSON(&customerDa)
	if err := customerDaService.CreateCustomer(customerDa); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCustomer 删除Customer
// @Tags Customer
// @Summary 删除Customer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Customer true "删除Customer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /customerDa/deleteCustomer [delete]
func (customerDaApi *CustomerApi) DeleteCustomer(c *gin.Context) {
	var customerDa autocode.Customer
	_ = c.ShouldBindJSON(&customerDa)
	if err := customerDaService.DeleteCustomer(customerDa); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCustomerByIds 批量删除Customer
// @Tags Customer
// @Summary 批量删除Customer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Customer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /customerDa/deleteCustomerByIds [delete]
func (customerDaApi *CustomerApi) DeleteCustomerByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := customerDaService.DeleteCustomerByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCustomer 更新Customer
// @Tags Customer
// @Summary 更新Customer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Customer true "更新Customer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /customerDa/updateCustomer [put]
func (customerDaApi *CustomerApi) UpdateCustomer(c *gin.Context) {
	var customerDa autocode.Customer
	_ = c.ShouldBindJSON(&customerDa)
	if err := customerDaService.UpdateCustomer(customerDa); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCustomer 用id查询Customer
// @Tags Customer
// @Summary 用id查询Customer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.Customer true "用id查询Customer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /customerDa/findCustomer [get]
func (customerDaApi *CustomerApi) FindCustomer(c *gin.Context) {
	var customerDa autocode.Customer
	_ = c.ShouldBindQuery(&customerDa)
	if err, recustomerDa := customerDaService.GetCustomer(customerDa.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recustomerDa": recustomerDa}, c)
	}
}

// GetCustomerList 分页获取Customer列表
// @Tags Customer
// @Summary 分页获取Customer列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.CustomerSearch true "分页获取Customer列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customerDa/getCustomerList [get]
func (customerDaApi *CustomerApi) GetCustomerList(c *gin.Context) {
	var pageInfo autocodeReq.CustomerSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := customerDaService.GetCustomerInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
