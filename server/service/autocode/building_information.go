package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type BuildingInformationService struct {
}

// CreateBuildingInformation 创建BuildingInformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (buildingInformationService *BuildingInformationService) CreateBuildingInformation(buildingInformation autocode.BuildingInformation) (err error) {
	err = global.GVA_DB.Create(&buildingInformation).Error
	return err
}

// DeleteBuildingInformation 删除BuildingInformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (buildingInformationService *BuildingInformationService)DeleteBuildingInformation(buildingInformation autocode.BuildingInformation) (err error) {
	err = global.GVA_DB.Delete(&buildingInformation).Error
	return err
}

// DeleteBuildingInformationByIds 批量删除BuildingInformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (buildingInformationService *BuildingInformationService)DeleteBuildingInformationByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.BuildingInformation{},"id in ?",ids.Ids).Error
	return err
}

// UpdateBuildingInformation 更新BuildingInformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (buildingInformationService *BuildingInformationService)UpdateBuildingInformation(buildingInformation autocode.BuildingInformation) (err error) {
	err = global.GVA_DB.Save(&buildingInformation).Error
	return err
}

// GetBuildingInformation 根据id获取BuildingInformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (buildingInformationService *BuildingInformationService)GetBuildingInformation(id uint) (err error, buildingInformation autocode.BuildingInformation) {
	err = global.GVA_DB.Where("id = ?", id).First(&buildingInformation).Error
	return
}

// GetBuildingInformationInfoList 分页获取BuildingInformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (buildingInformationService *BuildingInformationService)GetBuildingInformationInfoList(info autoCodeReq.BuildingInformationSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.BuildingInformation{})
    var buildingInformations []autocode.BuildingInformation
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Project != "" {
        db = db.Where("project LIKE ?","%"+ info.Project+"%")
    }
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
    if info.BuildState != nil {
        db = db.Where("build_state = ?",info.BuildState)
    }
    if info.CoveredArea != nil {
        db = db.Where("covered_area = ?",info.CoveredArea)
    }
    if info.BusinessArea != nil {
        db = db.Where("business_area = ?",info.BusinessArea)
    }
    if info.Upstairs != nil {
        db = db.Where("upstairs = ?",info.Upstairs)
    }
    if info.Downstair != nil {
        db = db.Where("downstair = ?",info.Downstair)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&buildingInformations).Error
	return err, buildingInformations, total
}
