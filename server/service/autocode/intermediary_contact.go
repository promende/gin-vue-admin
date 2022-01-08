package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type IntermediaryContactService struct {
}

// CreateIntermediaryContact 创建IntermediaryContact记录
// Author [piexlmax](https://github.com/piexlmax)
func (intermediarycontactService *IntermediaryContactService) CreateIntermediaryContact(intermediarycontact autocode.IntermediaryContact) (err error) {
	err = global.GVA_DB.Create(&intermediarycontact).Error
	return err
}

// DeleteIntermediaryContact 删除IntermediaryContact记录
// Author [piexlmax](https://github.com/piexlmax)
func (intermediarycontactService *IntermediaryContactService)DeleteIntermediaryContact(intermediarycontact autocode.IntermediaryContact) (err error) {
	err = global.GVA_DB.Delete(&intermediarycontact).Error
	return err
}

// DeleteIntermediaryContactByIds 批量删除IntermediaryContact记录
// Author [piexlmax](https://github.com/piexlmax)
func (intermediarycontactService *IntermediaryContactService)DeleteIntermediaryContactByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.IntermediaryContact{},"id in ?",ids.Ids).Error
	return err
}

// UpdateIntermediaryContact 更新IntermediaryContact记录
// Author [piexlmax](https://github.com/piexlmax)
func (intermediarycontactService *IntermediaryContactService)UpdateIntermediaryContact(intermediarycontact autocode.IntermediaryContact) (err error) {
	err = global.GVA_DB.Save(&intermediarycontact).Error
	return err
}

// GetIntermediaryContact 根据id获取IntermediaryContact记录
// Author [piexlmax](https://github.com/piexlmax)
func (intermediarycontactService *IntermediaryContactService)GetIntermediaryContact(id uint) (err error, intermediarycontact autocode.IntermediaryContact) {
	err = global.GVA_DB.Where("id = ?", id).First(&intermediarycontact).Error
	return
}

// GetIntermediaryContactInfoList 分页获取IntermediaryContact记录
// Author [piexlmax](https://github.com/piexlmax)
func (intermediarycontactService *IntermediaryContactService)GetIntermediaryContactInfoList(info autoCodeReq.IntermediaryContactSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.IntermediaryContact{})
    var intermediarycontacts []autocode.IntermediaryContact
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Company != "" {
        db = db.Where("company LIKE ?","%"+ info.Company+"%")
    }
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
    if info.TelephoneNumber != "" {
        db = db.Where("telephone_number LIKE ?","%"+ info.TelephoneNumber+"%")
    }
    if info.Principal != "" {
        db = db.Where("principal LIKE ?","%"+ info.Principal+"%")
    }
    if info.Remark != "" {
        db = db.Where("remark LIKE ?","%"+ info.Remark+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&intermediarycontacts).Error
	return err, intermediarycontacts, total
}
