// 自动生成模板Business
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Business 结构体
// 如果含有time.Time 请自行import time包
type Business struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:姓名;"`
      TelephoneNumber  string `json:"telephoneNumber" form:"telephoneNumber" gorm:"column:telephone_number;comment:电话号码;"`
      Source  *int `json:"source" form:"source" gorm:"column:source;comment:商机来源;"`
      Project  *int `json:"project" form:"project" gorm:"column:project;comment:意向项目;"`
      Principal  *int `json:"principal" form:"principal" gorm:"column:principal;comment:负责人;"`
}


// TableName Business 表名
func (Business) TableName() string {
  return "business"
}

