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

type ClientApi struct {
}

var clientService = service.ServiceGroupApp.AutoCodeServiceGroup.ClientService


// CreateClient 创建Client
// @Tags Client
// @Summary 创建Client
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Client true "创建Client"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /client/createClient [post]
func (clientApi *ClientApi) CreateClient(c *gin.Context) {
	var client autocode.Client
	_ = c.ShouldBindJSON(&client)
	if err := clientService.CreateClient(client); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteClient 删除Client
// @Tags Client
// @Summary 删除Client
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Client true "删除Client"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /client/deleteClient [delete]
func (clientApi *ClientApi) DeleteClient(c *gin.Context) {
	var client autocode.Client
	_ = c.ShouldBindJSON(&client)
	if err := clientService.DeleteClient(client); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteClientByIds 批量删除Client
// @Tags Client
// @Summary 批量删除Client
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Client"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /client/deleteClientByIds [delete]
func (clientApi *ClientApi) DeleteClientByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := clientService.DeleteClientByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateClient 更新Client
// @Tags Client
// @Summary 更新Client
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Client true "更新Client"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /client/updateClient [put]
func (clientApi *ClientApi) UpdateClient(c *gin.Context) {
	var client autocode.Client
	_ = c.ShouldBindJSON(&client)
	if err := clientService.UpdateClient(client); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindClient 用id查询Client
// @Tags Client
// @Summary 用id查询Client
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.Client true "用id查询Client"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /client/findClient [get]
func (clientApi *ClientApi) FindClient(c *gin.Context) {
	var client autocode.Client
	_ = c.ShouldBindQuery(&client)
	if err, reclient := clientService.GetClient(client.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reclient": reclient}, c)
	}
}

// GetClientList 分页获取Client列表
// @Tags Client
// @Summary 分页获取Client列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.ClientSearch true "分页获取Client列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /client/getClientList [get]
func (clientApi *ClientApi) GetClientList(c *gin.Context) {
	var pageInfo autocodeReq.ClientSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := clientService.GetClientInfoList(pageInfo); err != nil {
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
