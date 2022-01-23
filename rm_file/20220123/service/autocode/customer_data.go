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
func (customerDaService *CustomerService) CreateCustomer(customerDa autocode.Customer) (err error) {
	err = global.GVA_DB.Create(&customerDa).Error
	return err
}

// DeleteCustomer 删除Customer记录
// Author [piexlmax](https://github.com/piexlmax)
func (customerDaService *CustomerService)DeleteCustomer(customerDa autocode.Customer) (err error) {
	err = global.GVA_DB.Delete(&customerDa).Error
	return err
}

// DeleteCustomerByIds 批量删除Customer记录
// Author [piexlmax](https://github.com/piexlmax)
func (customerDaService *CustomerService)DeleteCustomerByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.Customer{},"id in ?",ids.Ids).Error
	return err
}

// UpdateCustomer 更新Customer记录
// Author [piexlmax](https://github.com/piexlmax)
func (customerDaService *CustomerService)UpdateCustomer(customerDa autocode.Customer) (err error) {
	err = global.GVA_DB.Save(&customerDa).Error
	return err
}

// GetCustomer 根据id获取Customer记录
// Author [piexlmax](https://github.com/piexlmax)
func (customerDaService *CustomerService)GetCustomer(id uint) (err error, customerDa autocode.Customer) {
	err = global.GVA_DB.Where("id = ?", id).First(&customerDa).Error
	return
}

// GetCustomerInfoList 分页获取Customer记录
// Author [piexlmax](https://github.com/piexlmax)
func (customerDaService *CustomerService)GetCustomerInfoList(info autoCodeReq.CustomerSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.Customer{})
    var customerDas []autocode.Customer
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.TData != "" {
        db = db.Where("t_data = ?",info.TData)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&customerDas).Error
	return err, customerDas, total
}
