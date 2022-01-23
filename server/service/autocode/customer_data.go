package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type CustomerService struct {
}

// CreateCustomer 创建Customer记录
// Author [piexlmax](https://github.com/piexlmax)
func (customerDataService *CustomerService) CreateCustomer(customerData autocode.Customer) (err error) {
	err = global.GVA_DB.Create(&customerData).Error
	return err
}

// DeleteCustomer 删除Customer记录
// Author [piexlmax](https://github.com/piexlmax)
func (customerDataService *CustomerService)DeleteCustomer(customerData autocode.Customer) (err error) {
	err = global.GVA_DB.Delete(&customerData).Error
	return err
}

// DeleteCustomerByIds 批量删除Customer记录
// Author [piexlmax](https://github.com/piexlmax)
func (customerDataService *CustomerService)DeleteCustomerByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.Customer{},"id in ?",ids.Ids).Error
	return err
}

// UpdateCustomer 更新Customer记录
// Author [piexlmax](https://github.com/piexlmax)
func (customerDataService *CustomerService)UpdateCustomer(customerData autocode.Customer) (err error) {
	err = global.GVA_DB.Save(&customerData).Error
	return err
}

// GetCustomer 根据id获取Customer记录
// Author [piexlmax](https://github.com/piexlmax)
func (customerDataService *CustomerService)GetCustomer(id uint) (err error, customerData autocode.Customer) {
	err = global.GVA_DB.Where("id = ?", id).First(&customerData).Error
	return
}

// GetCustomerInfoList 分页获取Customer记录
// Author [piexlmax](https://github.com/piexlmax)
func (customerDataService *CustomerService)GetCustomerInfoList(info autoCodeReq.CustomerSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.Customer{})
    var customerDatas []autocode.Customer
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
    if info.Remark != "" {
        db = db.Where("remark LIKE ?","%"+ info.Remark+"%")
    }
    if info.Audit != nil {
        db = db.Where("audit = ?",info.Audit)
    }
    if info.Principal != "" {
        db = db.Where("principal LIKE ?","%"+ info.Principal+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&customerDatas).Error
	return err, customerDatas, total
}
