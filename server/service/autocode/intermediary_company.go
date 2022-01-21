package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type IntermediaryCompanyService struct {
}

// CreateIntermediaryCompany 创建IntermediaryCompany记录
// Author [piexlmax](https://github.com/piexlmax)
func (intermediaryCompanyService *IntermediaryCompanyService) CreateIntermediaryCompany(intermediaryCompany autocode.IntermediaryCompany) (err error) {
	err = global.GVA_DB.Create(&intermediaryCompany).Error
	return err
}

// DeleteIntermediaryCompany 删除IntermediaryCompany记录
// Author [piexlmax](https://github.com/piexlmax)
func (intermediaryCompanyService *IntermediaryCompanyService)DeleteIntermediaryCompany(intermediaryCompany autocode.IntermediaryCompany) (err error) {
	err = global.GVA_DB.Delete(&intermediaryCompany).Error
	return err
}

// DeleteIntermediaryCompanyByIds 批量删除IntermediaryCompany记录
// Author [piexlmax](https://github.com/piexlmax)
func (intermediaryCompanyService *IntermediaryCompanyService)DeleteIntermediaryCompanyByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.IntermediaryCompany{},"id in ?",ids.Ids).Error
	return err
}

// UpdateIntermediaryCompany 更新IntermediaryCompany记录
// Author [piexlmax](https://github.com/piexlmax)
func (intermediaryCompanyService *IntermediaryCompanyService)UpdateIntermediaryCompany(intermediaryCompany autocode.IntermediaryCompany) (err error) {
	err = global.GVA_DB.Save(&intermediaryCompany).Error
	return err
}

// GetIntermediaryCompany 根据id获取IntermediaryCompany记录
// Author [piexlmax](https://github.com/piexlmax)
func (intermediaryCompanyService *IntermediaryCompanyService)GetIntermediaryCompany(id uint) (err error, intermediaryCompany autocode.IntermediaryCompany) {
	err = global.GVA_DB.Where("id = ?", id).First(&intermediaryCompany).Error
	return
}

// GetIntermediaryCompanyInfoList 分页获取IntermediaryCompany记录
// Author [piexlmax](https://github.com/piexlmax)
func (intermediaryCompanyService *IntermediaryCompanyService)GetIntermediaryCompanyInfoList(info autoCodeReq.IntermediaryCompanySearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.IntermediaryCompany{})
    var intermediaryCompanys []autocode.IntermediaryCompany
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
    if info.TelephoneNumber != "" {
        db = db.Where("telephone_number LIKE ?","%"+ info.TelephoneNumber+"%")
    }
    if info.Code != "" {
        db = db.Where("code LIKE ?","%"+ info.Code+"%")
    }
    if info.Creator != "" {
        db = db.Where("creator LIKE ?","%"+ info.Creator+"%")
    }
    if info.AuditType != nil {
        db = db.Where("audit_type = ?",info.AuditType)
    }
    if info.Remark != "" {
        db = db.Where("remark LIKE ?","%"+ info.Remark+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&intermediaryCompanys).Error
	return err, intermediaryCompanys, total
}
