// 自动生成模板Customer
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Customer 结构体
// 如果含有time.Time 请自行import time包
type Customer struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:商家名称;"`
      Type  *int `json:"type" form:"type" gorm:"column:type;comment:商家性质;"`
      Linkman  string `json:"linkman" form:"linkman" gorm:"column:linkman;comment:联系人;"`
      IDNumber  string `json:"iDNumber" form:"iDNumber" gorm:"column:i_d_number;comment:身份证号;"`
      Address  string `json:"address" form:"address" gorm:"column:address;comment:详细地址;"`
      Telephone  string `json:"telephone" form:"telephone" gorm:"column:telephone;comment:联系电话;"`
      Invoice  string `json:"invoice" form:"invoice" gorm:"column:invoice;comment:开票名称;"`
      Bank  string `json:"bank" form:"bank" gorm:"column:bank;comment:开户银行;"`
      Account  string `json:"account" form:"account" gorm:"column:account;comment:开户账号;"`
      Remark  string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;"`
      Audit  *int `json:"audit" form:"audit" gorm:"column:audit;comment:审核状态;"`
      Principal  string `json:"principal" form:"principal" gorm:"column:principal;comment:负责人;"`
}


// TableName Customer 表名
func (Customer) TableName() string {
  return "customer"
}

