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

type ContractApi struct {
}

var pactService = service.ServiceGroupApp.AutoCodeServiceGroup.ContractService


// CreateContract 创建Contract
// @Tags Contract
// @Summary 创建Contract
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Contract true "创建Contract"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pact/createContract [post]
func (pactApi *ContractApi) CreateContract(c *gin.Context) {
	var pact autocode.Contract
	_ = c.ShouldBindJSON(&pact)
	if err := pactService.CreateContract(pact); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteContract 删除Contract
// @Tags Contract
// @Summary 删除Contract
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Contract true "删除Contract"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pact/deleteContract [delete]
func (pactApi *ContractApi) DeleteContract(c *gin.Context) {
	var pact autocode.Contract
	_ = c.ShouldBindJSON(&pact)
	if err := pactService.DeleteContract(pact); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteContractByIds 批量删除Contract
// @Tags Contract
// @Summary 批量删除Contract
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Contract"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /pact/deleteContractByIds [delete]
func (pactApi *ContractApi) DeleteContractByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := pactService.DeleteContractByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateContract 更新Contract
// @Tags Contract
// @Summary 更新Contract
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Contract true "更新Contract"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /pact/updateContract [put]
func (pactApi *ContractApi) UpdateContract(c *gin.Context) {
	var pact autocode.Contract
	_ = c.ShouldBindJSON(&pact)
	if err := pactService.UpdateContract(pact); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindContract 用id查询Contract
// @Tags Contract
// @Summary 用id查询Contract
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.Contract true "用id查询Contract"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /pact/findContract [get]
func (pactApi *ContractApi) FindContract(c *gin.Context) {
	var pact autocode.Contract
	_ = c.ShouldBindQuery(&pact)
	if err, repact := pactService.GetContract(pact.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"repact": repact}, c)
	}
}

// GetContractList 分页获取Contract列表
// @Tags Contract
// @Summary 分页获取Contract列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.ContractSearch true "分页获取Contract列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pact/getContractList [get]
func (pactApi *ContractApi) GetContractList(c *gin.Context) {
	var pageInfo autocodeReq.ContractSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := pactService.GetContractInfoList(pageInfo); err != nil {
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
