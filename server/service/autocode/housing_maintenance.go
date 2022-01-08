package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type HousingMaintenanceService struct {
}

// CreateHousingMaintenance 创建HousingMaintenance记录
// Author [piexlmax](https://github.com/piexlmax)
func (housingMaintenanceService *HousingMaintenanceService) CreateHousingMaintenance(housingMaintenance autocode.HousingMaintenance) (err error) {
	err = global.GVA_DB.Create(&housingMaintenance).Error
	return err
}

// DeleteHousingMaintenance 删除HousingMaintenance记录
// Author [piexlmax](https://github.com/piexlmax)
func (housingMaintenanceService *HousingMaintenanceService)DeleteHousingMaintenance(housingMaintenance autocode.HousingMaintenance) (err error) {
	err = global.GVA_DB.Delete(&housingMaintenance).Error
	return err
}

// DeleteHousingMaintenanceByIds 批量删除HousingMaintenance记录
// Author [piexlmax](https://github.com/piexlmax)
func (housingMaintenanceService *HousingMaintenanceService)DeleteHousingMaintenanceByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.HousingMaintenance{},"id in ?",ids.Ids).Error
	return err
}

// UpdateHousingMaintenance 更新HousingMaintenance记录
// Author [piexlmax](https://github.com/piexlmax)
func (housingMaintenanceService *HousingMaintenanceService)UpdateHousingMaintenance(housingMaintenance autocode.HousingMaintenance) (err error) {
	err = global.GVA_DB.Save(&housingMaintenance).Error
	return err
}

// GetHousingMaintenance 根据id获取HousingMaintenance记录
// Author [piexlmax](https://github.com/piexlmax)
func (housingMaintenanceService *HousingMaintenanceService)GetHousingMaintenance(id uint) (err error, housingMaintenance autocode.HousingMaintenance) {
	err = global.GVA_DB.Where("id = ?", id).First(&housingMaintenance).Error
	return
}

// GetHousingMaintenanceInfoList 分页获取HousingMaintenance记录
// Author [piexlmax](https://github.com/piexlmax)
func (housingMaintenanceService *HousingMaintenanceService)GetHousingMaintenanceInfoList(info autoCodeReq.HousingMaintenanceSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.HousingMaintenance{})
    var housingMaintenances []autocode.HousingMaintenance
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Project != "" {
        db = db.Where("project LIKE ?","%"+ info.Project+"%")
    }
    if info.Build != "" {
        db = db.Where("build LIKE ?","%"+ info.Build+"%")
    }
    if info.Floor != "" {
        db = db.Where("floor LIKE ?","%"+ info.Floor+"%")
    }
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
    if info.State != nil {
        db = db.Where("state = ?",info.State)
    }
    if info.Type != nil {
        db = db.Where("type = ?",info.Type)
    }
    if info.MeterRentArea != nil {
        db = db.Where("meter_rent_area = ?",info.MeterRentArea)
    }
    if info.UsableArea != nil {
        db = db.Where("usable_area = ?",info.UsableArea)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&housingMaintenances).Error
	return err, housingMaintenances, total
}
