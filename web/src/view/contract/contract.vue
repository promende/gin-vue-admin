<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="合同编号">
          <el-select v-model="searchInfo.contractNumber" placeholder="请选择" style="width:100%" default-first-option clearable filterable @visible-change="setContractNumberOptions">
            <el-option v-for="(item,key) in contractNumberOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="商家名称">
          <el-select v-model="searchInfo.merchant" placeholder="请选择" style="width:100%" default-first-option clearable filterable @visible-change="setMerchantOptions">
            <el-option v-for="(item,key) in merchantOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="项目名称">
          <el-select v-model="searchInfo.project" placeholder="请选择" style="width:100%" default-first-option clearable filterable @visible-change="setProjectOptions">
            <el-option v-for="(item,key) in projectOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="负责人">
          <el-select v-model="searchInfo.principal" placeholder="请选择" style="width:100%" default-first-option clearable filterable @visible-change="setPrincipalOptions">
            <el-option v-for="(item,key) in principalOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
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
            <el-button style="margin-left:10px" size="mini" type="primary" icon="download" @click="exportExcel">导出</el-button>
        </div>
        <el-table
        id="rebateSetTable"
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        :default-sort = "{prop: 'date', order: 'descending'}"
        row-key="ID"
        @selection-change="handleSelectionChange"
        @row-click="handleRowClick"
        >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" type="expand" >
          <template #default="scope">
            <el-descriptions :column="3" border>
              <el-descriptions-item label="合同编号" label-align="center" align="left" width="10px" label-class-name="background: #E1F3D8;">
                {{ scope.row.contractNumber }}
              </el-descriptions-item>
              <el-descriptions-item label="项目名称" label-align="center" align="left" width="10px">
                {{ scope.row.project }}
              </el-descriptions-item>
              <el-descriptions-item label="楼栋名称" label-align="center" align="left" width="10px">
                {{ scope.row.building }}
              </el-descriptions-item>
              <el-descriptions-item label="楼层名称" label-align="center" align="left" width="80px">
                {{ scope.row.floor }}
              </el-descriptions-item>
              <el-descriptions-item label="房源名称" label-align="center" align="left" width="80px">
                {{ scope.row.housing }}
              </el-descriptions-item>
              <el-descriptions-item label="商家名称" label-align="center" align="left" width="80px">
                {{ scope.row.merchant }}
              </el-descriptions-item>
              <el-descriptions-item label="合同类型" label-align="center" align="left" width="80px">
                {{ filterDict(scope.row.contractType,"contractType") }}
              </el-descriptions-item>
              <el-descriptions-item label="合同签署" label-align="center" align="left" width="80px">
                {{ filterDict(scope.row.contractSigning,"contractSigning") }}
              </el-descriptions-item>
              <el-descriptions-item label="是否续签" label-align="center" align="left" width="80px">
                {{ filterDict(scope.row.renew,"renew") }}
              </el-descriptions-item>
              <el-descriptions-item label="关联合同编号" label-align="center" align="left" width="80px">
                {{ scope.row.associatedContractNumber }}
              </el-descriptions-item>
              <el-descriptions-item label="是否中介介入" label-align="center" align="left" width="80px">
                {{ filterDict(scope.row.intermediary,"Intermediary") }}
              </el-descriptions-item>
              <el-descriptions-item label="中介公司" label-align="center" align="left" width="10px" label-class-name="background: #E1F3D8;">
                {{ scope.row.agency }}
              </el-descriptions-item>
              <el-descriptions-item label="中介联系人" label-align="center" align="left" width="10px">
                {{ scope.row.intermediaryContact }}
              </el-descriptions-item>
              <el-descriptions-item label="负责人" label-align="center" align="left" width="10px">
                {{ scope.row.principal }}
              </el-descriptions-item>
              <el-descriptions-item label="交付日" label-align="center" align="left" width="80px">
                {{ formatDate(scope.row.deliveryDate) }}
              </el-descriptions-item>
              <el-descriptions-item label="合同开始时间" label-align="center" align="left" width="80px">
                {{ formatDate(scope.row.startTime) }}
              </el-descriptions-item>
              <el-descriptions-item label="合同结束时间" label-align="center" align="left" width="80px">
                {{ formatDate(scope.row.endTime) }}
              </el-descriptions-item>
              <el-descriptions-item label="支付周期" label-align="center" align="left" width="80px">
                {{ filterDict(scope.row.paymentCycle,"paymentCycle") }}
              </el-descriptions-item>
              <el-descriptions-item label="单价" label-align="center" align="left" width="10px" label-class-name="background: #E1F3D8;">
                {{ filterDict(scope.row.univalence,"univalence") }}
              </el-descriptions-item>
              <el-descriptions-item label="租金" label-align="center" align="left" width="10px">
                {{ scope.row.rent }}
              </el-descriptions-item>
              <el-descriptions-item label="服务费" label-align="center" align="left" width="10px">
                {{ scope.row.serviceCharge }}
              </el-descriptions-item>
              <el-descriptions-item label="物管费" label-align="center" align="left" width="80px">
                {{ scope.row.propertyManagementFee }}
              </el-descriptions-item>
              <el-descriptions-item label="合同总金额" label-align="center" align="left" width="80px">
                {{ scope.row.contractGrandTotal }}
              </el-descriptions-item>
              <el-descriptions-item label="设置费" label-align="center" align="left" width="80px">
                {{ scope.row.setUpFee }}
              </el-descriptions-item>
              <el-descriptions-item label="保证金" label-align="center" align="left" width="10px">
                {{ scope.row.earnestMoney }}
              </el-descriptions-item>
              <el-descriptions-item label="备注" label-align="center" align="left" width="80px">
                {{ scope.row.remark }}
              </el-descriptions-item>
              <el-descriptions-item label="创建日期" label-align="center" align="left" width="80px">
                {{ formatDate(scope.row.CreatedAt) }}
              </el-descriptions-item>
            </el-descriptions>
          </template>
        </el-table-column>
        <el-table-column align="left" label="合同编号" prop="contractNumber" width="120" />
        <el-table-column align="left" label="项目名称" prop="project" width="120" />
        <el-table-column align="left" label="商家名称" prop="merchant" width="120" />
        <el-table-column align="left" label="负责人" prop="principal" width="120" />
        <el-table-column align="left" label="创建日期" width="180" prop="date" sortable>
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="text" icon="edit" size="small" class="table-button" @click="updateContract(scope.row);handleRowClick(scope.row)">变更</el-button>
            <el-button type="text" icon="delete" size="mini" @click="deleteRow(scope.row);handleRowClick(scope.row)">删除</el-button>
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
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="新增合同">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="合同编号">
          <el-input v-model="formData.contractNumber" clearable placeholder="请输入" disabled/>
        </el-form-item>
        <el-form-item label="项目名称">
          <el-input v-model="formData.project" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="楼栋名称">
          <el-input v-model="formData.building" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="楼层名称">
          <el-input v-model="formData.floor" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="房源名称">
          <el-input v-model="formData.housing" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商家名称">
          <el-input v-model="formData.merchant" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="合同类型">
          <el-select v-model="formData.contractType" placeholder="请选择" style="width:100%" clearable>
            <el-option v-for="(item,key) in contractTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="合同签署">
          <el-select v-model="formData.contractSigning" placeholder="请选择" style="width:100%" clearable>
            <el-option v-for="(item,key) in contractSigningOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="是否续签">
          <el-select v-model="formData.renew" placeholder="请选择" style="width:100%" clearable>
            <el-option v-for="(item,key) in renewOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="关联合同编号">
          <el-input v-model="formData.associatedContractNumber" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="是否中介介入">
          <el-select v-model="formData.intermediary" placeholder="请选择" style="width:100%" clearable>
            <el-option v-for="(item,key) in IntermediaryOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="中介公司">
          <el-input v-model="formData.agency" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="中介联系人">
          <el-input v-model="formData.intermediaryContact" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="负责人">
          <el-input v-model="formData.principal" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="交付日">
          <el-date-picker v-model="formData.deliveryDate" type="date" style="width:100%" placeholder="选择日期" clearable />
        </el-form-item>
        <el-form-item label="合同开始时间">
          <el-date-picker v-model="formData.startTime" type="date" style="width:100%" placeholder="选择日期" clearable />
        </el-form-item>
        <el-form-item label="合同结束时间">
          <el-date-picker v-model="formData.endTime" type="date" style="width:100%" placeholder="选择日期" clearable />
        </el-form-item>
        <el-form-item label="支付周期">
          <el-select v-model="formData.paymentCycle" placeholder="请选择" style="width:100%" clearable>
            <el-option v-for="(item,key) in paymentCycleOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="单价">
          <el-select v-model="formData.univalence" placeholder="请选择" style="width:100%" clearable>
            <el-option v-for="(item,key) in univalenceOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="租金">
          <el-input-number v-model="formData.rent"  style="width:100%" :precision="2" clearable />
        </el-form-item>
        <el-form-item label="服务费">
          <el-input-number v-model="formData.serviceCharge"  style="width:100%" :precision="2" clearable />
        </el-form-item>
        <el-form-item label="物管费">
          <el-input-number v-model="formData.propertyManagementFee"  style="width:100%" :precision="2" clearable />
        </el-form-item>
        <el-form-item label="合同总金额">
          <el-input-number v-model="formData.contractGrandTotal"  style="width:100%" :precision="2" clearable />
        </el-form-item>
        <el-form-item label="设置费">
          <el-input-number v-model="formData.setUpFee"  style="width:100%" :precision="2" clearable />
        </el-form-item>
        <el-form-item label="保证金">
          <el-input-number v-model="formData.earnestMoney"  style="width:100%" :precision="2" clearable />
        </el-form-item>
        <el-form-item label="备注">
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
import { getUserList } from '@/api/user'
import { getProjectInformationList } from '@/api/projectInformation' // 此处请自行替换地址
import * as XLSX from 'xlsx';
import FileSaver from 'file-saver'
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
      principalOptions: [],
      projectOptions: [],
      contractNumberOptions: [],
      merchantOptions: [],
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
    async exportExcel() {
      /* generate workbook object from table */
      let oldPageSize = this.pageSize
      this.pageSize = 100
      await this.getTableData()
      let xlsxParam = { raw: true } // 导出的内容只做解析，不进行格式转换
      let wb = XLSX.utils.table_to_book(document.querySelector('#rebateSetTable'), xlsxParam);
      let ref = wb.Sheets.Sheet1['!ref']
      let flag = 0
      let maxCols = ""
      let maxRows = ""
      for(let i = 0; i < ref.length; i++){
        if(ref[i] === ":"){
          flag = 1;
          continue
        }
        if(flag === 1){
          if(ref[i]>="A"&&ref[i]<="Z"){
            maxCols += ref[i]
          }
          else{
            maxRows += ref[i];
          }
        }
      }
      for(let i = 0; i<=maxRows; i++){
        let cell = maxCols + i
        wb.Sheets.Sheet1[cell] = ''
        let firstCell = 'A' + i
        if(i === 1){
            wb.Sheets.Sheet1[firstCell] = {
            t: "s",
            v: "序号"
          }
        }
        else{
            wb.Sheets.Sheet1[firstCell] = {
            t: "s",
            v: (i - 1).toString()
          }
        }
      }
      /* get binary string as output */
      console.log(wb.Sheets.Sheet1)
      let wbout = XLSX.write(wb, { bookType: 'xlsx', bookSST: true, type: 'array' });
      // console.log(wbout)
      try {
          FileSaver.saveAs(new Blob([wbout], { type: 'application/octet-stream' }), '商机台账表.xlsx');
      } catch (e)
      {
        if (typeof console !== 'undefined')
            console.log(e, wbout)
      }

      this.pageSize = oldPageSize
      await this.getTableData()
      
      return wbout
    },
    handleRowClick(row) {
      row.expanded = !row.expanded;
      this.$refs.multipleTable.toggleRowExpansion(row, row.expanded);
    },
    async setContractNumberOptions() {
      this.contractNumberOptions = []
      const res = await getContractList({ page: 1, pageSize: 999 })
      res.data.list && res.data.list.forEach(item => {
        const option = {
          label: item.contractNumber,
          value: item.contractNumber
        }
        this.contractNumberOptions.push(option)
      })
    },
    async setMerchantOptions() {
      let list = []
      const res = await getContractList({ page: 1, pageSize: 999 })
      if(this.searchInfo.merchant === undefined || this.searchInfo.merchant === '') {
        res.data.list && res.data.list.forEach(item => {
          if(item.merchant != '')
            list.push(item.merchant)
        })
      }
      else {
        res.data.list && res.data.list.forEach(item => {
          if(item.merchant === this.searchInfo.merchant){
            if(item.merchant != '')
              list.push(item.merchant)          
          }
        })
      }
      let newArr = list.filter((item, index) => list.indexOf(item) === index); 

      this.merchantOptions = []
      newArr && newArr.forEach(item => {
        const option ={
          label: item,
          value: item
        }
        this.merchantOptions.push(option)
      })
    },
    async setProjectOptions() {
      this.projectOptions = []
      const res = await getProjectInformationList({ page: 1, pageSize: 999 })
      res.data.list && res.data.list.forEach(item => {
        const option = {
          label: item.name,
          value: item.name
        }
        this.projectOptions.push(option)
      })
    },
    async setPrincipalOptions() {
      this.principalOptions = []
      const res = await getUserList({ page: 1, pageSize: 999 })
      res.data.list && res.data.list.forEach(item => {
        const option = {
          label: item.nickName,
          value: item.nickName
        }
        this.principalOptions.push(option)
      })
    },
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
