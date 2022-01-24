package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type PactService struct {
}

// CreatePact 创建Pact记录
// Author [piexlmax](https://github.com/piexlmax)
func (ptService *PactService) CreatePact(pt autocode.Pact) (err error) {
	err = global.GVA_DB.Create(&pt).Error
	return err
}

// DeletePact 删除Pact记录
// Author [piexlmax](https://github.com/piexlmax)
func (ptService *PactService)DeletePact(pt autocode.Pact) (err error) {
	err = global.GVA_DB.Delete(&pt).Error
	return err
}

// DeletePactByIds 批量删除Pact记录
// Author [piexlmax](https://github.com/piexlmax)
func (ptService *PactService)DeletePactByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.Pact{},"id in ?",ids.Ids).Error
	return err
}

// UpdatePact 更新Pact记录
// Author [piexlmax](https://github.com/piexlmax)
func (ptService *PactService)UpdatePact(pt autocode.Pact) (err error) {
	err = global.GVA_DB.Save(&pt).Error
	return err
}

// GetPact 根据id获取Pact记录
// Author [piexlmax](https://github.com/piexlmax)
func (ptService *PactService)GetPact(id uint) (err error, pt autocode.Pact) {
	err = global.GVA_DB.Where("id = ?", id).First(&pt).Error
	return
}

// GetPactInfoList 分页获取Pact记录
// Author [piexlmax](https://github.com/piexlmax)
func (ptService *PactService)GetPactInfoList(info autoCodeReq.PactSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.Pact{})
    var pts []autocode.Pact
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.ContractNumber != "" {
        db = db.Where("contract_number LIKE ?","%"+ info.ContractNumber+"%")
    }
    if info.Project != "" {
        db = db.Where("project LIKE ?","%"+ info.Project+"%")
    }
    if info.Building != "" {
        db = db.Where("building LIKE ?","%"+ info.Building+"%")
    }
    if info.Floor != "" {
        db = db.Where("floor LIKE ?","%"+ info.Floor+"%")
    }
    if info.Housing != "" {
        db = db.Where("housing LIKE ?","%"+ info.Housing+"%")
    }
    if info.Merchant != "" {
        db = db.Where("merchant LIKE ?","%"+ info.Merchant+"%")
    }
    if info.ContractType != nil {
        db = db.Where("contract_type = ?",info.ContractType)
    }
    if info.ContractSigning != nil {
        db = db.Where("contract_signing = ?",info.ContractSigning)
    }
    if info.Renew != nil {
        db = db.Where("renew = ?",info.Renew)
    }
    if info.AssociatedContractNumber != "" {
        db = db.Where("associated_contract_number LIKE ?","%"+ info.AssociatedContractNumber+"%")
    }
    if info.Intermediary != nil {
        db = db.Where("intermediary = ?",info.Intermediary)
    }
    if info.Agency != "" {
        db = db.Where("agency LIKE ?","%"+ info.Agency+"%")
    }
    if info.IntermediaryContact != "" {
        db = db.Where("intermediary_contact LIKE ?","%"+ info.IntermediaryContact+"%")
    }
    if info.Principal != "" {
        db = db.Where("principal LIKE ?","%"+ info.Principal+"%")
    }
    if info.DeliveryDate != nil {
         db = db.Where("delivery_date = ?",info.DeliveryDate)
    }
    if info.StartTime != nil {
         db = db.Where("start_time = ?",info.StartTime)
    }
    if info.EndTime != nil {
         db = db.Where("end_time = ?",info.EndTime)
    }
    if info.PaymentCycle != nil {
        db = db.Where("payment_cycle = ?",info.PaymentCycle)
    }
    if info.Univalence != nil {
        db = db.Where("univalence = ?",info.Univalence)
    }
    if info.Rent != nil {
        db = db.Where("rent = ?",info.Rent)
    }
    if info.ServiceCharge != nil {
        db = db.Where("service_charge = ?",info.ServiceCharge)
    }
    if info.PropertyManagementFee != nil {
        db = db.Where("property_management_fee = ?",info.PropertyManagementFee)
    }
    if info.ContractGrandTotal != nil {
        db = db.Where("contract_grand_total = ?",info.ContractGrandTotal)
    }
    if info.SetUpFee != nil {
        db = db.Where("set_up_fee = ?",info.SetUpFee)
    }
    if info.EarnestMoney != nil {
        db = db.Where("earnest_money = ?",info.EarnestMoney)
    }
    if info.Remark != "" {
        db = db.Where("remark LIKE ?","%"+ info.Remark+"%")
    }
    if info.AuditType != nil {
        db = db.Where("audit_type = ?",info.AuditType)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&pts).Error
	return err, pts, total
}
