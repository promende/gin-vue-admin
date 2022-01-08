// 自动生成模板Intermediary
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Intermediary 结构体
// 如果含有time.Time 请自行import time包
type Intermediary struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:中介名称;"`
      TelephoneNumber  string `json:"telephoneNumber" form:"telephoneNumber" gorm:"column:telephone_number;comment:联系电话;"`
      Code  string `json:"code" form:"code" gorm:"column:code;comment:社会信用代码;"`
      Creator  string `json:"creator" form:"creator" gorm:"column:creator;comment:创建人;"`
      Remark  string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;"`
}


// TableName Intermediary 表名
func (Intermediary) TableName() string {
  return "intermediary"
}

