// 自动生成模板BuildingInformation
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// BuildingInformation 结构体
// 如果含有time.Time 请自行import time包
type BuildingInformation struct {
      global.GVA_MODEL
      Project  string `json:"project" form:"project" gorm:"column:project;comment:所属项目;"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:楼栋名称;"`
      BuildState  *int `json:"buildState" form:"buildState" gorm:"column:build_state;comment:楼栋状态;"`
      CoveredArea  *float64 `json:"coveredArea" form:"coveredArea" gorm:"column:covered_area;comment:建筑面积（㎡);"`
      BusinessArea  *float64 `json:"businessArea" form:"businessArea" gorm:"column:business_area;comment:经营面积（㎡）;"`
      Upstairs  *int `json:"upstairs" form:"upstairs" gorm:"column:upstairs;comment:楼上楼层数;"`
      Downstair  *int `json:"downstair" form:"downstair" gorm:"column:downstair;comment:楼下楼层数;"`
}


// TableName BuildingInformation 表名
func (BuildingInformation) TableName() string {
  return "BuildingInformation"
}

