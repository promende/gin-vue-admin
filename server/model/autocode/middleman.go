// 自动生成模板Middleman
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Middleman 结构体
// 如果含有time.Time 请自行import time包
type Middleman struct {
      global.GVA_MODEL
      Company  string `json:"company" form:"company" gorm:"column:company;comment:所属公司;"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:联系人;"`
      TelephoneNumber  string `json:"telephoneNumber" form:"telephoneNumber" gorm:"column:telephone_number;comment:联系电话;"`
      Remark  string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;"`
      AuditType  *int `json:"auditType" form:"auditType" gorm:"column:audit_type;comment:审核状态;"`
      Creator  string `json:"creator" form:"creator" gorm:"column:creator;comment:创建人;"`
}


// TableName Middleman 表名
func (Middleman) TableName() string {
  return "middleman"
}

