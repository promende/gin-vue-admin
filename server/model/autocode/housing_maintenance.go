// 自动生成模板HousingMaintenance
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// HousingMaintenance 结构体
// 如果含有time.Time 请自行import time包
type HousingMaintenance struct {
      global.GVA_MODEL
      Project  string `json:"project" form:"project" gorm:"column:project;comment:所属项目;"`
      Build  string `json:"build" form:"build" gorm:"column:build;comment:所属楼栋;"`
      Floor  string `json:"floor" form:"floor" gorm:"column:floor;comment:所属楼层;"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:房源名称;"`
      State  *int `json:"state" form:"state" gorm:"column:state;comment:房源状态;"`
      Type  *int `json:"type" form:"type" gorm:"column:type;comment:房源类型;"`
      MeterRentArea  *float64 `json:"meterRentArea" form:"meterRentArea" gorm:"column:meter_rent_area;comment:计租面积;"`
      UsableArea  *float64 `json:"usableArea" form:"usableArea" gorm:"column:usable_area;comment:实用面积;"`
}


// TableName HousingMaintenance 表名
func (HousingMaintenance) TableName() string {
  return "HousingMaintenance"
}

