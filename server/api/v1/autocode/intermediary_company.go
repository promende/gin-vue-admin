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

type IntermediaryCompanyApi struct {
}

var intermediaryCompanyService = service.ServiceGroupApp.AutoCodeServiceGroup.IntermediaryCompanyService


// CreateIntermediaryCompany 创建IntermediaryCompany
// @Tags IntermediaryCompany
// @Summary 创建IntermediaryCompany
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.IntermediaryCompany true "创建IntermediaryCompany"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /intermediaryCompany/createIntermediaryCompany [post]
func (intermediaryCompanyApi *IntermediaryCompanyApi) CreateIntermediaryCompany(c *gin.Context) {
	var intermediaryCompany autocode.IntermediaryCompany
	_ = c.ShouldBindJSON(&intermediaryCompany)
	if err := intermediaryCompanyService.CreateIntermediaryCompany(intermediaryCompany); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteIntermediaryCompany 删除IntermediaryCompany
// @Tags IntermediaryCompany
// @Summary 删除IntermediaryCompany
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.IntermediaryCompany true "删除IntermediaryCompany"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /intermediaryCompany/deleteIntermediaryCompany [delete]
func (intermediaryCompanyApi *IntermediaryCompanyApi) DeleteIntermediaryCompany(c *gin.Context) {
	var intermediaryCompany autocode.IntermediaryCompany
	_ = c.ShouldBindJSON(&intermediaryCompany)
	if err := intermediaryCompanyService.DeleteIntermediaryCompany(intermediaryCompany); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteIntermediaryCompanyByIds 批量删除IntermediaryCompany
// @Tags IntermediaryCompany
// @Summary 批量删除IntermediaryCompany
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除IntermediaryCompany"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /intermediaryCompany/deleteIntermediaryCompanyByIds [delete]
func (intermediaryCompanyApi *IntermediaryCompanyApi) DeleteIntermediaryCompanyByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := intermediaryCompanyService.DeleteIntermediaryCompanyByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateIntermediaryCompany 更新IntermediaryCompany
// @Tags IntermediaryCompany
// @Summary 更新IntermediaryCompany
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.IntermediaryCompany true "更新IntermediaryCompany"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /intermediaryCompany/updateIntermediaryCompany [put]
func (intermediaryCompanyApi *IntermediaryCompanyApi) UpdateIntermediaryCompany(c *gin.Context) {
	var intermediaryCompany autocode.IntermediaryCompany
	_ = c.ShouldBindJSON(&intermediaryCompany)
	if err := intermediaryCompanyService.UpdateIntermediaryCompany(intermediaryCompany); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindIntermediaryCompany 用id查询IntermediaryCompany
// @Tags IntermediaryCompany
// @Summary 用id查询IntermediaryCompany
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.IntermediaryCompany true "用id查询IntermediaryCompany"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /intermediaryCompany/findIntermediaryCompany [get]
func (intermediaryCompanyApi *IntermediaryCompanyApi) FindIntermediaryCompany(c *gin.Context) {
	var intermediaryCompany autocode.IntermediaryCompany
	_ = c.ShouldBindQuery(&intermediaryCompany)
	if err, reintermediaryCompany := intermediaryCompanyService.GetIntermediaryCompany(intermediaryCompany.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reintermediaryCompany": reintermediaryCompany}, c)
	}
}

// GetIntermediaryCompanyList 分页获取IntermediaryCompany列表
// @Tags IntermediaryCompany
// @Summary 分页获取IntermediaryCompany列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.IntermediaryCompanySearch true "分页获取IntermediaryCompany列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /intermediaryCompany/getIntermediaryCompanyList [get]
func (intermediaryCompanyApi *IntermediaryCompanyApi) GetIntermediaryCompanyList(c *gin.Context) {
	var pageInfo autocodeReq.IntermediaryCompanySearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := intermediaryCompanyService.GetIntermediaryCompanyInfoList(pageInfo); err != nil {
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
