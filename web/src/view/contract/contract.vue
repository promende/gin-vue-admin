<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="合同编号">
          <el-input v-model="searchInfo.contractNumber" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="项目名称">
          <el-input v-model="searchInfo.project" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="楼栋名称">
          <el-input v-model="searchInfo.building" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="楼层名称">
          <el-input v-model="searchInfo.floor" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="房源名称">
          <el-input v-model="searchInfo.housing" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="商家名称">
          <el-input v-model="searchInfo.merchant" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="合同类型">
          <el-input v-model="searchInfo.contractType" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="合同签署">
          <el-input v-model="searchInfo.contractSigning" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="是否续签">
          <el-input v-model="searchInfo.renew" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="关联合同编号">
          <el-input v-model="searchInfo.associatedContractNumber" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="是否中介介入">
          <el-input v-model="searchInfo.intermediary" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="中介公司">
          <el-input v-model="searchInfo.agency" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="中介联系人">
          <el-input v-model="searchInfo.intermediaryContact" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="负责人">
          <el-input v-model="searchInfo.principal" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="交付日">
          <el-input v-model="searchInfo.deliveryDate" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="合同开始时间">
          <el-input v-model="searchInfo.startTime" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="合同结束时间">
          <el-input v-model="searchInfo.endTime" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="支付周期">
          <el-input v-model="searchInfo.paymentCycle" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="单价">
          <el-input v-model="searchInfo.univalence" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="租金">
          <el-input v-model="searchInfo.rent" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="服务费">
          <el-input v-model="searchInfo.serviceCharge" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="物管费">
          <el-input v-model="searchInfo.propertyManagementFee" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="合同总金额">
          <el-input v-model="searchInfo.contractGrandTotal" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="设置费">
          <el-input v-model="searchInfo.setUpFee" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="保证金">
          <el-input v-model="searchInfo.earnestMoney" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="searchInfo.remark" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="mini" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button size="mini" type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-popover v-model:visible="deleteVisible" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin-top: 8px;">
                <el-button size="mini" type="text" @click="deleteVisible = false">取消</el-button>
                <el-button size="mini" type="primary" @click="onDelete">确定</el-button>
            </div>
            <template #reference>
                <el-button icon="delete" size="mini" style="margin-left: 10px;" :disabled="!multipleSelection.length">删除</el-button>
            </template>
            </el-popover>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="合同编号" prop="contractNumber" width="120" />
        <el-table-column align="left" label="项目名称" prop="project" width="120" />
        <el-table-column align="left" label="楼栋名称" prop="building" width="120" />
        <el-table-column align="left" label="楼层名称" prop="floor" width="120" />
        <el-table-column align="left" label="房源名称" prop="housing" width="120" />
        <el-table-column align="left" label="商家名称" prop="merchant" width="120" />
        <el-table-column align="left" label="合同类型" prop="contractType" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.contractType,"contractType") }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="合同签署" prop="contractSigning" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.contractSigning,"contractSigning") }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="是否续签" prop="renew" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.renew,"renew") }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="关联合同编号" prop="associatedContractNumber" width="120" />
        <el-table-column align="left" label="是否中介介入" prop="intermediary" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.intermediary,"Intermediary") }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="中介公司" prop="agency" width="120" />
        <el-table-column align="left" label="中介联系人" prop="intermediaryContact" width="120" />
        <el-table-column align="left" label="负责人" prop="principal" width="120" />
        <el-table-column align="left" label="交付日" prop="deliveryDate" width="120" />
        <el-table-column align="left" label="合同开始时间" prop="startTime" width="120" />
        <el-table-column align="left" label="合同结束时间" prop="endTime" width="120" />
        <el-table-column align="left" label="支付周期" prop="paymentCycle" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.paymentCycle,"paymentCycle") }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="单价" prop="univalence" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.univalence,"univalence") }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="租金" prop="rent" width="120" />
        <el-table-column align="left" label="服务费" prop="serviceCharge" width="120" />
        <el-table-column align="left" label="物管费" prop="propertyManagementFee" width="120" />
        <el-table-column align="left" label="合同总金额" prop="contractGrandTotal" width="120" />
        <el-table-column align="left" label="设置费" prop="setUpFee" width="120" />
        <el-table-column align="left" label="保证金" prop="earnestMoney" width="120" />
        <el-table-column align="left" label="备注" prop="remark" width="120" />
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="text" icon="edit" size="small" class="table-button" @click="updateContract(scope.row)">变更</el-button>
            <el-button type="text" icon="delete" size="mini" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
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
          <el-select v-model="formData.contractType" placeholder="请选择" style="width:100%" clearable>
            <el-option v-for="(item,key) in contractTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="合同签署:">
          <el-select v-model="formData.contractSigning" placeholder="请选择" style="width:100%" clearable>
            <el-option v-for="(item,key) in contractSigningOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="是否续签:">
          <el-select v-model="formData.renew" placeholder="请选择" style="width:100%" clearable>
            <el-option v-for="(item,key) in renewOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="关联合同编号:">
          <el-input v-model="formData.associatedContractNumber" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="是否中介介入:">
          <el-select v-model="formData.intermediary" placeholder="请选择" style="width:100%" clearable>
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
          <el-date-picker v-model="formData.deliveryDate" type="date" style="width:100%" placeholder="选择日期" clearable />
        </el-form-item>
        <el-form-item label="合同开始时间:">
          <el-date-picker v-model="formData.startTime" type="date" style="width:100%" placeholder="选择日期" clearable />
        </el-form-item>
        <el-form-item label="合同结束时间:">
          <el-date-picker v-model="formData.endTime" type="date" style="width:100%" placeholder="选择日期" clearable />
        </el-form-item>
        <el-form-item label="支付周期:">
          <el-select v-model="formData.paymentCycle" placeholder="请选择" style="width:100%" clearable>
            <el-option v-for="(item,key) in paymentCycleOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="单价:">
          <el-select v-model="formData.univalence" placeholder="请选择" style="width:100%" clearable>
            <el-option v-for="(item,key) in univalenceOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="租金:">
          <el-input-number v-model="formData.rent"  style="width:100%" :precision="2" clearable />
        </el-form-item>
        <el-form-item label="服务费:">
          <el-input-number v-model="formData.serviceCharge"  style="width:100%" :precision="2" clearable />
        </el-form-item>
        <el-form-item label="物管费:">
          <el-input-number v-model="formData.propertyManagementFee"  style="width:100%" :precision="2" clearable />
        </el-form-item>
        <el-form-item label="合同总金额:">
          <el-input-number v-model="formData.contractGrandTotal"  style="width:100%" :precision="2" clearable />
        </el-form-item>
        <el-form-item label="设置费:">
          <el-input-number v-model="formData.setUpFee"  style="width:100%" :precision="2" clearable />
        </el-form-item>
        <el-form-item label="保证金:">
          <el-input-number v-model="formData.earnestMoney"  style="width:100%" :precision="2" clearable />
        </el-form-item>
        <el-form-item label="备注:">
          <el-input v-model="formData.remark" clearable placeholder="请输入" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
import {
  createContract,
  deleteContract,
  deleteContractByIds,
  updateContract,
  findContract,
  getContractList
} from '@/api/contract' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'Contract',
  mixins: [infoList],
  data() {
    return {
      listApi: getContractList,
      dialogFormVisible: false,
      type: '',
      deleteVisible: false,
      multipleSelection: [],
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
    await this.getTableData()
    await this.getDict('Intermediary')
    await this.getDict('paymentCycle')
    await this.getDict('univalence')
    await this.getDict('contractType')
    await this.getDict('contractSigning')
    await this.getDict('renew')
  },
  methods: {
  onReset() {
    this.searchInfo = {}
  },
  // 条件搜索前端看此方法
    onSubmit() {
      this.page = 1
      this.pageSize = 10
      this.getTableData()
    },
    handleSelectionChange(val) {
      this.multipleSelection = val
    },
    deleteRow(row) {
      this.$confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.deleteContract(row)
      })
    },
    async onDelete() {
      const ids = []
      if (this.multipleSelection.length === 0) {
        this.$message({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      this.multipleSelection &&
        this.multipleSelection.map(item => {
          ids.push(item.ID)
        })
      const res = await deleteContractByIds({ ids })
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '删除成功'
        })
        if (this.tableData.length === ids.length && this.page > 1) {
          this.page--
        }
        this.deleteVisible = false
        this.getTableData()
      }
    },
    async updateContract(row) {
      const res = await findContract({ ID: row.ID })
      this.type = 'update'
      if (res.code === 0) {
        this.formData = res.data.repact
        this.dialogFormVisible = true
      }
    },
    closeDialog() {
      this.dialogFormVisible = false
      this.formData = {
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
    },
    async deleteContract(row) {
      const res = await deleteContract({ ID: row.ID })
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '删除成功'
        })
        if (this.tableData.length === 1 && this.page > 1) {
          this.page--
        }
        this.getTableData()
      }
    },
    async enterDialog() {
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
        this.closeDialog()
        this.getTableData()
      }
    },
    openDialog() {
      this.type = 'create'
      this.dialogFormVisible = true
    }
  },
}
</script>

<style>
</style>
