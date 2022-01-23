// 自动生成模板Customer
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Customer 结构体
// 如果含有time.Time 请自行import time包
type Customer struct {
      global.GVA_MODEL
      TData  string `json:"tData" form:"tData" gorm:"column:t_data;comment:;"`
}


// TableName Customer 表名
func (Customer) TableName() string {
  return "customer"
}

