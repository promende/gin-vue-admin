package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type FloorInformationService struct {
}

// CreateFloorInformation 创建FloorInformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (floorInformationService *FloorInformationService) CreateFloorInformation(floorInformation autocode.FloorInformation) (err error) {
	err = global.GVA_DB.Create(&floorInformation).Error
	return err
}

// DeleteFloorInformation 删除FloorInformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (floorInformationService *FloorInformationService)DeleteFloorInformation(floorInformation autocode.FloorInformation) (err error) {
	err = global.GVA_DB.Delete(&floorInformation).Error
	return err
}

// DeleteFloorInformationByIds 批量删除FloorInformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (floorInformationService *FloorInformationService)DeleteFloorInformationByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.FloorInformation{},"id in ?",ids.Ids).Error
	return err
}

// UpdateFloorInformation 更新FloorInformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (floorInformationService *FloorInformationService)UpdateFloorInformation(floorInformation autocode.FloorInformation) (err error) {
	err = global.GVA_DB.Save(&floorInformation).Error
	return err
}

// GetFloorInformation 根据id获取FloorInformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (floorInformationService *FloorInformationService)GetFloorInformation(id uint) (err error, floorInformation autocode.FloorInformation) {
	err = global.GVA_DB.Where("id = ?", id).First(&floorInformation).Error
	return
}

// GetFloorInformationInfoList 分页获取FloorInformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (floorInformationService *FloorInformationService)GetFloorInformationInfoList(info autoCodeReq.FloorInformationSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.FloorInformation{})
    var floorInformations []autocode.FloorInformation
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Project != "" {
        db = db.Where("project LIKE ?","%"+ info.Project+"%")
    }
    if info.Build != "" {
        db = db.Where("build LIKE ?","%"+ info.Build+"%")
    }
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
    if info.FloorState != nil {
        db = db.Where("floor_state = ?",info.FloorState)
    }
    if info.CoveredArea != nil {
        db = db.Where("covered_area = ?",info.CoveredArea)
    }
    if info.OperatingArea != nil {
        db = db.Where("operating_area = ?",info.OperatingArea)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&floorInformations).Error
	return err, floorInformations, total
}
