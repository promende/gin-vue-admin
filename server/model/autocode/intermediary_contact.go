// 自动生成模板IntermediaryContact
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// IntermediaryContact 结构体
// 如果含有time.Time 请自行import time包
type IntermediaryContact struct {
      global.GVA_MODEL
      Company  string `json:"company" form:"company" gorm:"column:company;comment:所属公司;"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:联系人;"`
      TelephoneNumber  string `json:"telephoneNumber" form:"telephoneNumber" gorm:"column:telephone_number;comment:联系电话;"`
      Principal  string `json:"principal" form:"principal" gorm:"column:principal;comment:负责人;"`
      Remark  string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;"`
}


// TableName IntermediaryContact 表名
func (IntermediaryContact) TableName() string {
  return "IntermediaryContact"
}

