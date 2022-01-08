// 自动生成模板ProjectInformation
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// ProjectInformation 结构体
// 如果含有time.Time 请自行import time包
type ProjectInformation struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:项目名称;"`
      Abbreviation  string `json:"abbreviation" form:"abbreviation" gorm:"column:abbreviation;comment:项目简称;"`
      Address  string `json:"address" form:"address" gorm:"column:address;comment:项目地址;"`
      OperatingState  *int `json:"operatingState" form:"operatingState" gorm:"column:operating_state;comment:营运状态;"`
      ManagementType  *int `json:"managementType" form:"managementType" gorm:"column:management_type;comment:经营类型;"`
      PropertyManagementType  *int `json:"propertyManagementType" form:"propertyManagementType" gorm:"column:property_management_type;comment:物业管理类型;"`
      CoveredArea  *float64 `json:"coveredArea" form:"coveredArea" gorm:"column:covered_area;comment:建筑面积（㎡);"`
      OperatingArea  *float64 `json:"operatingArea" form:"operatingArea" gorm:"column:operating_area;comment:经营面积（㎡）;"`
      Principal  string `json:"principal" form:"principal" gorm:"column:principal;comment:负责人;"`
}


// TableName ProjectInformation 表名
func (ProjectInformation) TableName() string {
  return "projectInformation"
}

