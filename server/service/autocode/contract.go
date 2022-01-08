package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type ContractService struct {
}

// CreateContract 创建Contract记录
// Author [piexlmax](https://github.com/piexlmax)
func (pactService *ContractService) CreateContract(pact autocode.Contract) (err error) {
	err = global.GVA_DB.Create(&pact).Error
	return err
}

// DeleteContract 删除Contract记录
// Author [piexlmax](https://github.com/piexlmax)
func (pactService *ContractService)DeleteContract(pact autocode.Contract) (err error) {
	err = global.GVA_DB.Delete(&pact).Error
	return err
}

// DeleteContractByIds 批量删除Contract记录
// Author [piexlmax](https://github.com/piexlmax)
func (pactService *ContractService)DeleteContractByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.Contract{},"id in ?",ids.Ids).Error
	return err
}

// UpdateContract 更新Contract记录
// Author [piexlmax](https://github.com/piexlmax)
func (pactService *ContractService)UpdateContract(pact autocode.Contract) (err error) {
	err = global.GVA_DB.Save(&pact).Error
	return err
}

// GetContract 根据id获取Contract记录
// Author [piexlmax](https://github.com/piexlmax)
func (pactService *ContractService)GetContract(id uint) (err error, pact autocode.Contract) {
	err = global.GVA_DB.Where("id = ?", id).First(&pact).Error
	return
}

// GetContractInfoList 分页获取Contract记录
// Author [piexlmax](https://github.com/piexlmax)
func (pactService *ContractService)GetContractInfoList(info autoCodeReq.ContractSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.Contract{})
    var pacts []autocode.Contract
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
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&pacts).Error
	return err, pacts, total
}
