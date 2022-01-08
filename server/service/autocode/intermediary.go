package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type IntermediaryService struct {
}

// CreateIntermediary 创建Intermediary记录
// Author [piexlmax](https://github.com/piexlmax)
func (intermediaryService *IntermediaryService) CreateIntermediary(intermediary autocode.Intermediary) (err error) {
	err = global.GVA_DB.Create(&intermediary).Error
	return err
}

// DeleteIntermediary 删除Intermediary记录
// Author [piexlmax](https://github.com/piexlmax)
func (intermediaryService *IntermediaryService)DeleteIntermediary(intermediary autocode.Intermediary) (err error) {
	err = global.GVA_DB.Delete(&intermediary).Error
	return err
}

// DeleteIntermediaryByIds 批量删除Intermediary记录
// Author [piexlmax](https://github.com/piexlmax)
func (intermediaryService *IntermediaryService)DeleteIntermediaryByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.Intermediary{},"id in ?",ids.Ids).Error
	return err
}

// UpdateIntermediary 更新Intermediary记录
// Author [piexlmax](https://github.com/piexlmax)
func (intermediaryService *IntermediaryService)UpdateIntermediary(intermediary autocode.Intermediary) (err error) {
	err = global.GVA_DB.Save(&intermediary).Error
	return err
}

// GetIntermediary 根据id获取Intermediary记录
// Author [piexlmax](https://github.com/piexlmax)
func (intermediaryService *IntermediaryService)GetIntermediary(id uint) (err error, intermediary autocode.Intermediary) {
	err = global.GVA_DB.Where("id = ?", id).First(&intermediary).Error
	return
}

// GetIntermediaryInfoList 分页获取Intermediary记录
// Author [piexlmax](https://github.com/piexlmax)
func (intermediaryService *IntermediaryService)GetIntermediaryInfoList(info autoCodeReq.IntermediarySearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.Intermediary{})
    var intermediarys []autocode.Intermediary
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
    if info.Remark != "" {
        db = db.Where("remark LIKE ?","%"+ info.Remark+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&intermediarys).Error
	return err, intermediarys, total
}
