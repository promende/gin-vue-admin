// 自动生成模板Pact
package autocode

import (
      "github.com/flipped-aurora/gin-vue-admin/server/global"
      "time"
)

// Pact 结构体
// 如果含有time.Time 请自行import time包
type Pact struct {
      global.GVA_MODEL
      ContractNumber  string `json:"contractNumber" form:"contractNumber" gorm:"column:contract_number;comment:合同编号;"`
      Project  string `json:"project" form:"project" gorm:"column:project;comment:项目名称;"`
      Building  string `json:"building" form:"building" gorm:"column:building;comment:楼栋名称;"`
      Floor  string `json:"floor" form:"floor" gorm:"column:floor;comment:楼层名称;"`
      Housing  string `json:"housing" form:"housing" gorm:"column:housing;comment:房源名称;"`
      Merchant  string `json:"merchant" form:"merchant" gorm:"column:merchant;comment:商家名称;"`
      ContractType  *int `json:"contractType" form:"contractType" gorm:"column:contract_type;comment:合同类型;"`
      ContractSigning  *int `json:"contractSigning" form:"contractSigning" gorm:"column:contract_signing;comment:合同签署;"`
      Renew  *int `json:"renew" form:"renew" gorm:"column:renew;comment:是否续签;"`
      AssociatedContractNumber  string `json:"associatedContractNumber" form:"associatedContractNumber" gorm:"column:associated_contract_number;comment:关联合同编号;"`
      Intermediary  *int `json:"intermediary" form:"intermediary" gorm:"column:intermediary;comment:是否中介介入;"`
      Agency  string `json:"agency" form:"agency" gorm:"column:agency;comment:中介公司;"`
      IntermediaryContact  string `json:"intermediaryContact" form:"intermediaryContact" gorm:"column:intermediary_contact;comment:中介联系人;"`
      Principal  string `json:"principal" form:"principal" gorm:"column:principal;comment:负责人;"`
      DeliveryDate  *time.Time `json:"deliveryDate" form:"deliveryDate" gorm:"column:delivery_date;comment:交付日;"`
      StartTime  *time.Time `json:"startTime" form:"startTime" gorm:"column:start_time;comment:合同开始时间;"`
      EndTime  *time.Time `json:"endTime" form:"endTime" gorm:"column:end_time;comment:合同结束时间;"`
      PaymentCycle  *int `json:"paymentCycle" form:"paymentCycle" gorm:"column:payment_cycle;comment:支付周期;"`
      Univalence  *int `json:"univalence" form:"univalence" gorm:"column:univalence;comment:univalence;"`
      Rent  *float64 `json:"rent" form:"rent" gorm:"column:rent;comment:租金;"`
      ServiceCharge  *float64 `json:"serviceCharge" form:"serviceCharge" gorm:"column:service_charge;comment:服务费;"`
      PropertyManagementFee  *float64 `json:"propertyManagementFee" form:"propertyManagementFee" gorm:"column:property_management_fee;comment:物管费;"`
      ContractGrandTotal  *float64 `json:"contractGrandTotal" form:"contractGrandTotal" gorm:"column:contract_grand_total;comment:合同总金额;"`
      SetUpFee  *float64 `json:"setUpFee" form:"setUpFee" gorm:"column:set_up_fee;comment:设置费;"`
      EarnestMoney  *float64 `json:"earnestMoney" form:"earnestMoney" gorm:"column:earnest_money;comment:保证金;"`
      Remark  string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;"`
      AuditType  *int `json:"auditType" form:"auditType" gorm:"column:audit_type;comment:审核状态;"`
}


// TableName Pact 表名
func (Pact) TableName() string {
  return "pact"
}

