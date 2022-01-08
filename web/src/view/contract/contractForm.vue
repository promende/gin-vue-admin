<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="合同编号:">
          <el-input v-model="formData.contractNumber" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="项目名称:">
          <el-input v-model="formData.project" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="楼栋名称:">
          <el-input v-model="formData.building" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="楼层名称:">
          <el-input v-model="formData.floor" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="房源名称:">
          <el-input v-model="formData.housing" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商家名称:">
          <el-input v-model="formData.merchant" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="合同类型:">
          <el-select v-model="formData.contractType" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in contractTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="合同签署:">
          <el-select v-model="formData.contractSigning" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in contractSigningOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="是否续签:">
          <el-select v-model="formData.renew" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in renewOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="关联合同编号:">
          <el-input v-model="formData.associatedContractNumber" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="是否中介介入:">
          <el-select v-model="formData.intermediary" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in IntermediaryOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="中介公司:">
          <el-input v-model="formData.agency" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="中介联系人:">
          <el-input v-model="formData.intermediaryContact" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="负责人:">
          <el-input v-model="formData.principal" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="交付日:">
          <el-date-picker v-model="formData.deliveryDate" type="date" placeholder="选择日期" clearable></el-date-picker>
        </el-form-item>
        <el-form-item label="合同开始时间:">
          <el-date-picker v-model="formData.startTime" type="date" placeholder="选择日期" clearable></el-date-picker>
        </el-form-item>
        <el-form-item label="合同结束时间:">
          <el-date-picker v-model="formData.endTime" type="date" placeholder="选择日期" clearable></el-date-picker>
        </el-form-item>
        <el-form-item label="支付周期:">
          <el-select v-model="formData.paymentCycle" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in paymentCycleOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="单价:">
          <el-select v-model="formData.univalence" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in univalenceOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="租金:">
          <el-input-number v-model="formData.rent" :precision="2" clearable></el-input-number>
        </el-form-item>
        <el-form-item label="服务费:">
          <el-input-number v-model="formData.serviceCharge" :precision="2" clearable></el-input-number>
        </el-form-item>
        <el-form-item label="物管费:">
          <el-input-number v-model="formData.propertyManagementFee" :precision="2" clearable></el-input-number>
        </el-form-item>
        <el-form-item label="合同总金额:">
          <el-input-number v-model="formData.contractGrandTotal" :precision="2" clearable></el-input-number>
        </el-form-item>
        <el-form-item label="设置费:">
          <el-input-number v-model="formData.setUpFee" :precision="2" clearable></el-input-number>
        </el-form-item>
        <el-form-item label="保证金:">
          <el-input-number v-model="formData.earnestMoney" :precision="2" clearable></el-input-number>
        </el-form-item>
        <el-form-item label="备注:">
          <el-input v-model="formData.remark" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" @click="save">保存</el-button>
          <el-button size="mini" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
import {
  createContract,
  updateContract,
  findContract
} from '@/api/contract' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'Contract',
  mixins: [infoList],
  data() {
    return {
      type: '',
      IntermediaryOptions: [],
      paymentCycleOptions: [],
      univalenceOptions: [],
      contractTypeOptions: [],
      contractSigningOptions: [],
      renewOptions: [],
      formData: {
        contractNumber: '',
        project: '',
        building: '',
        floor: '',
        housing: '',
        merchant: '',
        contractType: undefined,
        contractSigning: undefined,
        renew: undefined,
        associatedContractNumber: '',
        intermediary: undefined,
        agency: '',
        intermediaryContact: '',
        principal: '',
        deliveryDate: new Date(),
        startTime: new Date(),
        endTime: new Date(),
        paymentCycle: undefined,
        univalence: undefined,
        rent: 0,
        serviceCharge: 0,
        propertyManagementFee: 0,
        contractGrandTotal: 0,
        setUpFee: 0,
        earnestMoney: 0,
        remark: '',
      }
    }
  },
  async created() {
    // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (this.$route.query.id) {
      const res = await findContract({ ID: this.$route.query.id })
      if (res.code === 0) {
        this.formData = res.data.repact
        this.type = 'update'
      }
    } else {
      this.type = 'create'
    }
    await this.getDict('Intermediary')
    await this.getDict('paymentCycle')
    await this.getDict('univalence')
    await this.getDict('contractType')
    await this.getDict('contractSigning')
    await this.getDict('renew')
  },
  methods: {
    async save() {
      let res
      switch (this.type) {
        case 'create':
          res = await createContract(this.formData)
          break
        case 'update':
          res = await updateContract(this.formData)
          break
        default:
          res = await createContract(this.formData)
          break
      }
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '创建/更改成功'
        })
      }
    },
    back() {
      this.$router.go(-1)
    }
  }
}
</script>

<style>
</style>
