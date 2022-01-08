// 自动生成模板FloorInformation
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// FloorInformation 结构体
// 如果含有time.Time 请自行import time包
type FloorInformation struct {
      global.GVA_MODEL
      Project  string `json:"project" form:"project" gorm:"column:project;comment:所属项目;"`
      Build  string `json:"build" form:"build" gorm:"column:build;comment:所属楼栋;"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:楼层名称;"`
      FloorState  *int `json:"floorState" form:"floorState" gorm:"column:floor_state;comment:楼层状态;"`
      CoveredArea  *float64 `json:"coveredArea" form:"coveredArea" gorm:"column:covered_area;comment:建筑面积（㎡);"`
      OperatingArea  *float64 `json:"operatingArea" form:"operatingArea" gorm:"column:operating_area;comment:经营面积（㎡）;"`
}


// TableName FloorInformation 表名
func (FloorInformation) TableName() string {
  return "FloorInformation"
}

