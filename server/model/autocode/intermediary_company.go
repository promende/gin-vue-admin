// 自动生成模板IntermediaryCompany
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// IntermediaryCompany 结构体
// 如果含有time.Time 请自行import time包
type IntermediaryCompany struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:中介名称;"`
      TelephoneNumber  string `json:"telephoneNumber" form:"telephoneNumber" gorm:"column:telephone_number;comment:联系电话;"`
      Code  string `json:"code" form:"code" gorm:"column:code;comment:社会信用代码;"`
      Creator  string `json:"creator" form:"creator" gorm:"column:creator;comment:创建人;"`
      AuditType  *int `json:"auditType" form:"auditType" gorm:"column:audit_type;comment:审核状态;"`
      Remark  string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;"`
}


// TableName IntermediaryCompany 表名
func (IntermediaryCompany) TableName() string {
  return "intermediaryCompany"
}

