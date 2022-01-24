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

type PactApi struct {
}

var ptService = service.ServiceGroupApp.AutoCodeServiceGroup.PactService


// CreatePact 创建Pact
// @Tags Pact
// @Summary 创建Pact
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Pact true "创建Pact"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pt/createPact [post]
func (ptApi *PactApi) CreatePact(c *gin.Context) {
	var pt autocode.Pact
	_ = c.ShouldBindJSON(&pt)
	if err := ptService.CreatePact(pt); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeletePact 删除Pact
// @Tags Pact
// @Summary 删除Pact
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Pact true "删除Pact"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pt/deletePact [delete]
func (ptApi *PactApi) DeletePact(c *gin.Context) {
	var pt autocode.Pact
	_ = c.ShouldBindJSON(&pt)
	if err := ptService.DeletePact(pt); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeletePactByIds 批量删除Pact
// @Tags Pact
// @Summary 批量删除Pact
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Pact"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /pt/deletePactByIds [delete]
func (ptApi *PactApi) DeletePactByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := ptService.DeletePactByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdatePact 更新Pact
// @Tags Pact
// @Summary 更新Pact
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Pact true "更新Pact"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /pt/updatePact [put]
func (ptApi *PactApi) UpdatePact(c *gin.Context) {
	var pt autocode.Pact
	_ = c.ShouldBindJSON(&pt)
	if err := ptService.UpdatePact(pt); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindPact 用id查询Pact
// @Tags Pact
// @Summary 用id查询Pact
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.Pact true "用id查询Pact"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /pt/findPact [get]
func (ptApi *PactApi) FindPact(c *gin.Context) {
	var pt autocode.Pact
	_ = c.ShouldBindQuery(&pt)
	if err, rept := ptService.GetPact(pt.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rept": rept}, c)
	}
}

// GetPactList 分页获取Pact列表
// @Tags Pact
// @Summary 分页获取Pact列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.PactSearch true "分页获取Pact列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pt/getPactList [get]
func (ptApi *PactApi) GetPactList(c *gin.Context) {
	var pageInfo autocodeReq.PactSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := ptService.GetPactInfoList(pageInfo); err != nil {
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
