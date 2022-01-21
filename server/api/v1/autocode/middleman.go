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

type MiddlemanApi struct {
}

var middlemanService = service.ServiceGroupApp.AutoCodeServiceGroup.MiddlemanService


// CreateMiddleman 创建Middleman
// @Tags Middleman
// @Summary 创建Middleman
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Middleman true "创建Middleman"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /middleman/createMiddleman [post]
func (middlemanApi *MiddlemanApi) CreateMiddleman(c *gin.Context) {
	var middleman autocode.Middleman
	_ = c.ShouldBindJSON(&middleman)
	if err := middlemanService.CreateMiddleman(middleman); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMiddleman 删除Middleman
// @Tags Middleman
// @Summary 删除Middleman
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Middleman true "删除Middleman"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /middleman/deleteMiddleman [delete]
func (middlemanApi *MiddlemanApi) DeleteMiddleman(c *gin.Context) {
	var middleman autocode.Middleman
	_ = c.ShouldBindJSON(&middleman)
	if err := middlemanService.DeleteMiddleman(middleman); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMiddlemanByIds 批量删除Middleman
// @Tags Middleman
// @Summary 批量删除Middleman
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Middleman"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /middleman/deleteMiddlemanByIds [delete]
func (middlemanApi *MiddlemanApi) DeleteMiddlemanByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := middlemanService.DeleteMiddlemanByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMiddleman 更新Middleman
// @Tags Middleman
// @Summary 更新Middleman
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Middleman true "更新Middleman"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /middleman/updateMiddleman [put]
func (middlemanApi *MiddlemanApi) UpdateMiddleman(c *gin.Context) {
	var middleman autocode.Middleman
	_ = c.ShouldBindJSON(&middleman)
	if err := middlemanService.UpdateMiddleman(middleman); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMiddleman 用id查询Middleman
// @Tags Middleman
// @Summary 用id查询Middleman
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.Middleman true "用id查询Middleman"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /middleman/findMiddleman [get]
func (middlemanApi *MiddlemanApi) FindMiddleman(c *gin.Context) {
	var middleman autocode.Middleman
	_ = c.ShouldBindQuery(&middleman)
	if err, remiddleman := middlemanService.GetMiddleman(middleman.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"remiddleman": remiddleman}, c)
	}
}

// GetMiddlemanList 分页获取Middleman列表
// @Tags Middleman
// @Summary 分页获取Middleman列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.MiddlemanSearch true "分页获取Middleman列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /middleman/getMiddlemanList [get]
func (middlemanApi *MiddlemanApi) GetMiddlemanList(c *gin.Context) {
	var pageInfo autocodeReq.MiddlemanSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := middlemanService.GetMiddlemanInfoList(pageInfo); err != nil {
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
