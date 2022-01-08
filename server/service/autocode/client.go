package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type ClientService struct {
}

// CreateClient 创建Client记录
// Author [piexlmax](https://github.com/piexlmax)
func (clientService *ClientService) CreateClient(client autocode.Client) (err error) {
	err = global.GVA_DB.Create(&client).Error
	return err
}

// DeleteClient 删除Client记录
// Author [piexlmax](https://github.com/piexlmax)
func (clientService *ClientService)DeleteClient(client autocode.Client) (err error) {
	err = global.GVA_DB.Delete(&client).Error
	return err
}

// DeleteClientByIds 批量删除Client记录
// Author [piexlmax](https://github.com/piexlmax)
func (clientService *ClientService)DeleteClientByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.Client{},"id in ?",ids.Ids).Error
	return err
}

// UpdateClient 更新Client记录
// Author [piexlmax](https://github.com/piexlmax)
func (clientService *ClientService)UpdateClient(client autocode.Client) (err error) {
	err = global.GVA_DB.Save(&client).Error
	return err
}

// GetClient 根据id获取Client记录
// Author [piexlmax](https://github.com/piexlmax)
func (clientService *ClientService)GetClient(id uint) (err error, client autocode.Client) {
	err = global.GVA_DB.Where("id = ?", id).First(&client).Error
	return
}

// GetClientInfoList 分页获取Client记录
// Author [piexlmax](https://github.com/piexlmax)
func (clientService *ClientService)GetClientInfoList(info autoCodeReq.ClientSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.Client{})
    var clients []autocode.Client
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
    if info.Type != nil {
        db = db.Where("type = ?",info.Type)
    }
    if info.Linkman != "" {
        db = db.Where("linkman LIKE ?","%"+ info.Linkman+"%")
    }
    if info.IDNumber != "" {
        db = db.Where("i_d_number LIKE ?","%"+ info.IDNumber+"%")
    }
    if info.Address != "" {
        db = db.Where("address LIKE ?","%"+ info.Address+"%")
    }
    if info.Telephone != "" {
        db = db.Where("telephone LIKE ?","%"+ info.Telephone+"%")
    }
    if info.Invoice != "" {
        db = db.Where("invoice LIKE ?","%"+ info.Invoice+"%")
    }
    if info.Bank != "" {
        db = db.Where("bank LIKE ?","%"+ info.Bank+"%")
    }
    if info.Account != "" {
        db = db.Where("account LIKE ?","%"+ info.Account+"%")
    }
    if info.Audit != nil {
        db = db.Where("audit = ?",info.Audit)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&clients).Error
	return err, clients, total
}
