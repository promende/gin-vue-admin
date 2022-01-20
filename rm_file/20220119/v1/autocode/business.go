package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autocodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BusinessApi struct {
}

var businessService = service.ServiceGroupApp.AutoCodeServiceGroup.BusinessService


// CreateBusiness 创建Business
// @Tags Business
// @Summary 创建Business
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Business true "创建Business"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /business/createBusiness [post]
func (businessApi *BusinessApi) CreateBusiness(c *gin.Context) {
	var business autocode.Business
	_ = c.ShouldBindJSON(&business)
	if err := businessService.CreateBusiness(business); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBusiness 删除Business
// @Tags Business
// @Summary 删除Business
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Business true "删除Business"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /business/deleteBusiness [delete]
func (businessApi *BusinessApi) DeleteBusiness(c *gin.Context) {
	var business autocode.Business
	_ = c.ShouldBindJSON(&business)
	if err := businessService.DeleteBusiness(business); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBusinessByIds 批量删除Business
// @Tags Business
// @Summary 批量删除Business
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Business"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /business/deleteBusinessByIds [delete]
func (businessApi *BusinessApi) DeleteBusinessByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := businessService.DeleteBusinessByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBusiness 更新Business
// @Tags Business
// @Summary 更新Business
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Business true "更新Business"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /business/updateBusiness [put]
func (businessApi *BusinessApi) UpdateBusiness(c *gin.Context) {
	var business autocode.Business
	_ = c.ShouldBindJSON(&business)
	if err := businessService.UpdateBusiness(business); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBusiness 用id查询Business
// @Tags Business
// @Summary 用id查询Business
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.Business true "用id查询Business"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /business/findBusiness [get]
func (businessApi *BusinessApi) FindBusiness(c *gin.Context) {
	var business autocode.Business
	_ = c.ShouldBindQuery(&business)
	if err, rebusiness := businessService.GetBusiness(business.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rebusiness": rebusiness}, c)
	}
}

// GetBusinessList 分页获取Business列表
// @Tags Business
// @Summary 分页获取Business列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.BusinessSearch true "分页获取Business列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /business/getBusinessList [get]
func (businessApi *BusinessApi) GetBusinessList(c *gin.Context) {
	var pageInfo autocodeReq.BusinessSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := businessService.GetBusinessInfoList(pageInfo); err != nil {
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
