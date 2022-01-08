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

type IntermediaryContactApi struct {
}

var intermediarycontactService = service.ServiceGroupApp.AutoCodeServiceGroup.IntermediaryContactService


// CreateIntermediaryContact 创建IntermediaryContact
// @Tags IntermediaryContact
// @Summary 创建IntermediaryContact
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.IntermediaryContact true "创建IntermediaryContact"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /intermediarycontact/createIntermediaryContact [post]
func (intermediarycontactApi *IntermediaryContactApi) CreateIntermediaryContact(c *gin.Context) {
	var intermediarycontact autocode.IntermediaryContact
	_ = c.ShouldBindJSON(&intermediarycontact)
	if err := intermediarycontactService.CreateIntermediaryContact(intermediarycontact); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteIntermediaryContact 删除IntermediaryContact
// @Tags IntermediaryContact
// @Summary 删除IntermediaryContact
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.IntermediaryContact true "删除IntermediaryContact"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /intermediarycontact/deleteIntermediaryContact [delete]
func (intermediarycontactApi *IntermediaryContactApi) DeleteIntermediaryContact(c *gin.Context) {
	var intermediarycontact autocode.IntermediaryContact
	_ = c.ShouldBindJSON(&intermediarycontact)
	if err := intermediarycontactService.DeleteIntermediaryContact(intermediarycontact); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteIntermediaryContactByIds 批量删除IntermediaryContact
// @Tags IntermediaryContact
// @Summary 批量删除IntermediaryContact
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除IntermediaryContact"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /intermediarycontact/deleteIntermediaryContactByIds [delete]
func (intermediarycontactApi *IntermediaryContactApi) DeleteIntermediaryContactByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := intermediarycontactService.DeleteIntermediaryContactByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateIntermediaryContact 更新IntermediaryContact
// @Tags IntermediaryContact
// @Summary 更新IntermediaryContact
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.IntermediaryContact true "更新IntermediaryContact"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /intermediarycontact/updateIntermediaryContact [put]
func (intermediarycontactApi *IntermediaryContactApi) UpdateIntermediaryContact(c *gin.Context) {
	var intermediarycontact autocode.IntermediaryContact
	_ = c.ShouldBindJSON(&intermediarycontact)
	if err := intermediarycontactService.UpdateIntermediaryContact(intermediarycontact); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindIntermediaryContact 用id查询IntermediaryContact
// @Tags IntermediaryContact
// @Summary 用id查询IntermediaryContact
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.IntermediaryContact true "用id查询IntermediaryContact"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /intermediarycontact/findIntermediaryContact [get]
func (intermediarycontactApi *IntermediaryContactApi) FindIntermediaryContact(c *gin.Context) {
	var intermediarycontact autocode.IntermediaryContact
	_ = c.ShouldBindQuery(&intermediarycontact)
	if err, reintermediarycontact := intermediarycontactService.GetIntermediaryContact(intermediarycontact.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reintermediarycontact": reintermediarycontact}, c)
	}
}

// GetIntermediaryContactList 分页获取IntermediaryContact列表
// @Tags IntermediaryContact
// @Summary 分页获取IntermediaryContact列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.IntermediaryContactSearch true "分页获取IntermediaryContact列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /intermediarycontact/getIntermediaryContactList [get]
func (intermediarycontactApi *IntermediaryContactApi) GetIntermediaryContactList(c *gin.Context) {
	var pageInfo autocodeReq.IntermediaryContactSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := intermediarycontactService.GetIntermediaryContactInfoList(pageInfo); err != nil {
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
