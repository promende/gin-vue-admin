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

type IntermediaryApi struct {
}

var intermediaryService = service.ServiceGroupApp.AutoCodeServiceGroup.IntermediaryService


// CreateIntermediary 创建Intermediary
// @Tags Intermediary
// @Summary 创建Intermediary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Intermediary true "创建Intermediary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /intermediary/createIntermediary [post]
func (intermediaryApi *IntermediaryApi) CreateIntermediary(c *gin.Context) {
	var intermediary autocode.Intermediary
	_ = c.ShouldBindJSON(&intermediary)
	if err := intermediaryService.CreateIntermediary(intermediary); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteIntermediary 删除Intermediary
// @Tags Intermediary
// @Summary 删除Intermediary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Intermediary true "删除Intermediary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /intermediary/deleteIntermediary [delete]
func (intermediaryApi *IntermediaryApi) DeleteIntermediary(c *gin.Context) {
	var intermediary autocode.Intermediary
	_ = c.ShouldBindJSON(&intermediary)
	if err := intermediaryService.DeleteIntermediary(intermediary); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteIntermediaryByIds 批量删除Intermediary
// @Tags Intermediary
// @Summary 批量删除Intermediary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Intermediary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /intermediary/deleteIntermediaryByIds [delete]
func (intermediaryApi *IntermediaryApi) DeleteIntermediaryByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := intermediaryService.DeleteIntermediaryByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateIntermediary 更新Intermediary
// @Tags Intermediary
// @Summary 更新Intermediary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Intermediary true "更新Intermediary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /intermediary/updateIntermediary [put]
func (intermediaryApi *IntermediaryApi) UpdateIntermediary(c *gin.Context) {
	var intermediary autocode.Intermediary
	_ = c.ShouldBindJSON(&intermediary)
	if err := intermediaryService.UpdateIntermediary(intermediary); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindIntermediary 用id查询Intermediary
// @Tags Intermediary
// @Summary 用id查询Intermediary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.Intermediary true "用id查询Intermediary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /intermediary/findIntermediary [get]
func (intermediaryApi *IntermediaryApi) FindIntermediary(c *gin.Context) {
	var intermediary autocode.Intermediary
	_ = c.ShouldBindQuery(&intermediary)
	if err, reintermediary := intermediaryService.GetIntermediary(intermediary.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reintermediary": reintermediary}, c)
	}
}

// GetIntermediaryList 分页获取Intermediary列表
// @Tags Intermediary
// @Summary 分页获取Intermediary列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.IntermediarySearch true "分页获取所有公司名"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /intermediary/getIntermediaryList [get]
func (intermediaryApi *IntermediaryApi) GetIntermediaryList(c *gin.Context) {
	var pageInfo autocodeReq.IntermediarySearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := intermediaryService.GetIntermediaryInfoList(pageInfo); err != nil {
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
