package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type ProjectInformationService struct {
}

// CreateProjectInformation 创建ProjectInformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (projectService *ProjectInformationService) CreateProjectInformation(project autocode.ProjectInformation) (err error) {
	err = global.GVA_DB.Create(&project).Error
	return err
}

// DeleteProjectInformation 删除ProjectInformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (projectService *ProjectInformationService)DeleteProjectInformation(project autocode.ProjectInformation) (err error) {
	err = global.GVA_DB.Delete(&project).Error
	return err
}

// DeleteProjectInformationByIds 批量删除ProjectInformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (projectService *ProjectInformationService)DeleteProjectInformationByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.ProjectInformation{},"id in ?",ids.Ids).Error
	return err
}

// UpdateProjectInformation 更新ProjectInformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (projectService *ProjectInformationService)UpdateProjectInformation(project autocode.ProjectInformation) (err error) {
	err = global.GVA_DB.Save(&project).Error
	return err
}

// GetProjectInformation 根据id获取ProjectInformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (projectService *ProjectInformationService)GetProjectInformation(id uint) (err error, project autocode.ProjectInformation) {
	err = global.GVA_DB.Where("id = ?", id).First(&project).Error
	return
}

// GetProjectInformationInfoList 分页获取ProjectInformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (projectService *ProjectInformationService)GetProjectInformationInfoList(info autoCodeReq.ProjectInformationSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.ProjectInformation{})
    var projects []autocode.ProjectInformation
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
    if info.Abbreviation != "" {
        db = db.Where("abbreviation LIKE ?","%"+ info.Abbreviation+"%")
    }
    if info.Address != "" {
        db = db.Where("address LIKE ?","%"+ info.Address+"%")
    }
    if info.OperatingState != nil {
        db = db.Where("operating_state = ?",info.OperatingState)
    }
    if info.ManagementType != nil {
        db = db.Where("management_type = ?",info.ManagementType)
    }
    if info.PropertyManagementType != nil {
        db = db.Where("property_management_type = ?",info.PropertyManagementType)
    }
    if info.CoveredArea != nil {
        db = db.Where("covered_area = ?",info.CoveredArea)
    }
    if info.OperatingArea != nil {
        db = db.Where("operating_area = ?",info.OperatingArea)
    }
    if info.Principal != "" {
        db = db.Where("principal LIKE ?","%"+ info.Principal+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&projects).Error
	return err, projects, total
}
