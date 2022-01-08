package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type BusinessService struct {
}

// CreateBusiness 创建Business记录
// Author [piexlmax](https://github.com/piexlmax)
func (businessService *BusinessService) CreateBusiness(business autocode.Business) (err error) {
	err = global.GVA_DB.Create(&business).Error
	return err
}

// DeleteBusiness 删除Business记录
// Author [piexlmax](https://github.com/piexlmax)
func (businessService *BusinessService)DeleteBusiness(business autocode.Business) (err error) {
	err = global.GVA_DB.Delete(&business).Error
	return err
}

// DeleteBusinessByIds 批量删除Business记录
// Author [piexlmax](https://github.com/piexlmax)
func (businessService *BusinessService)DeleteBusinessByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.Business{},"id in ?",ids.Ids).Error
	return err
}

// UpdateBusiness 更新Business记录
// Author [piexlmax](https://github.com/piexlmax)
func (businessService *BusinessService)UpdateBusiness(business autocode.Business) (err error) {
	err = global.GVA_DB.Save(&business).Error
	return err
}

// GetBusiness 根据id获取Business记录
// Author [piexlmax](https://github.com/piexlmax)
func (businessService *BusinessService)GetBusiness(id uint) (err error, business autocode.Business) {
	err = global.GVA_DB.Where("id = ?", id).First(&business).Error
	return
}

// GetBusinessInfoList 分页获取Business记录
// Author [piexlmax](https://github.com/piexlmax)
func (businessService *BusinessService)GetBusinessInfoList(info autoCodeReq.BusinessSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.Business{})
    var businesss []autocode.Business
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
    if info.TelephoneNumber != "" {
        db = db.Where("telephone_number LIKE ?","%"+ info.TelephoneNumber+"%")
    }
    if info.Source != nil {
        db = db.Where("source = ?",info.Source)
    }
    if info.Project != nil {
        db = db.Where("project = ?",info.Project)
    }
    if info.Principal != nil {
        db = db.Where("principal = ?",info.Principal)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&businesss).Error
	return err, businesss, total
}
