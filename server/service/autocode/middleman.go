package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type MiddlemanService struct {
}

// CreateMiddleman 创建Middleman记录
// Author [piexlmax](https://github.com/piexlmax)
func (middlemanService *MiddlemanService) CreateMiddleman(middleman autocode.Middleman) (err error) {
	err = global.GVA_DB.Create(&middleman).Error
	return err
}

// DeleteMiddleman 删除Middleman记录
// Author [piexlmax](https://github.com/piexlmax)
func (middlemanService *MiddlemanService)DeleteMiddleman(middleman autocode.Middleman) (err error) {
	err = global.GVA_DB.Delete(&middleman).Error
	return err
}

// DeleteMiddlemanByIds 批量删除Middleman记录
// Author [piexlmax](https://github.com/piexlmax)
func (middlemanService *MiddlemanService)DeleteMiddlemanByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.Middleman{},"id in ?",ids.Ids).Error
	return err
}

// UpdateMiddleman 更新Middleman记录
// Author [piexlmax](https://github.com/piexlmax)
func (middlemanService *MiddlemanService)UpdateMiddleman(middleman autocode.Middleman) (err error) {
	err = global.GVA_DB.Save(&middleman).Error
	return err
}

// GetMiddleman 根据id获取Middleman记录
// Author [piexlmax](https://github.com/piexlmax)
func (middlemanService *MiddlemanService)GetMiddleman(id uint) (err error, middleman autocode.Middleman) {
	err = global.GVA_DB.Where("id = ?", id).First(&middleman).Error
	return
}

// GetMiddlemanInfoList 分页获取Middleman记录
// Author [piexlmax](https://github.com/piexlmax)
func (middlemanService *MiddlemanService)GetMiddlemanInfoList(info autoCodeReq.MiddlemanSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.Middleman{})
    var middlemans []autocode.Middleman
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
    if info.Remark != "" {
        db = db.Where("remark LIKE ?","%"+ info.Remark+"%")
    }
    if info.AuditType != nil {
        db = db.Where("audit_type = ?",info.AuditType)
    }
    if info.Creator != "" {
        db = db.Where("creator LIKE ?","%"+ info.Creator+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&middlemans).Error
	return err, middlemans, total
}
