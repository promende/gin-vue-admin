<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="姓名">
          <el-select v-model="searchInfo.name" placeholder="请选择" style="width:100%" default-first-option clearable filterable @visible-change="setNameOptions">
            <el-option v-for="(item,key) in nameOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="电话号码">
          <el-select v-model="searchInfo.telephoneNumber" placeholder="请选择" style="width:100%" default-first-option clearable filterable @visible-change="setTelephoneNumberOptions">
            <el-option v-for="(item,key) in telephoneNumberOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="公司">
          <el-select v-model="searchInfo.company" placeholder="请选择" style="width:100%" default-first-option clearable filterable @visible-change="setCompanyOptions">
            <el-option v-for="(item,key) in companyOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="商机来源">
          <el-select v-model="searchInfo.source" placeholder="请选择" style="width:100%" default-first-option clearable filterable >
            <el-option v-for="(item,key) in SourceOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="意向项目">
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
        row-key="ID"
        :default-sort = "{prop: 'date', order: 'descending'}"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="姓名" prop="name" width="120" />
        <el-table-column align="left" label="电话号码" prop="telephoneNumber" width="120" />
        <el-table-column align="left" label="公司" prop="company" width="300" />
        <el-table-column align="left" label="职务" prop="duty" width="150" />
        <el-table-column align="left" label="商机来源" prop="source" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.source,"Source") }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="意向项目" prop="project" width="120" />
        <el-table-column align="left" label="负责人" prop="principal" width="120" />
        <el-table-column align="left" label="创建日期" width="180" prop="date"  sortable>
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="text" icon="edit" size="small" class="table-button" @click="updateBusiness(scope.row)">变更</el-button>
            <el-button type="text" icon="delete" size="mini" @click="deleteRow(scope.row)">删除</el-button>
            <el-button type="text" icon="tools" size="small" @click="hintChangeToCustomer(scope.row)">转换客户</el-button>    
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
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="新增商机">
      <el-form :model="formData" label-position="right" label-width="80px" ref="formData" :rules="rules">
        <el-form-item label="姓名" prop="name">
          <el-input v-model.trim="formData.name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="电话号码" prop="telephoneNumber">
          <el-input v-model.trim="formData.telephoneNumber" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="公司">
          <el-input v-model="formData.company" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="职务">
          <el-input v-model="formData.duty" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商机来源" prop="source">
          <el-select v-model="formData.source" placeholder="请选择" style="width:100%" clearable>
            <el-option v-for="(item,key) in SourceOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="意向项目" prop="project">
          <el-select v-model="formData.project" placeholder="请选择" style="width:100%" clearable @visible-change="setProjectOptions">
            <el-option v-for="(item,key) in projectOptions" :key="key" :label="item.label" :value="item.value"/>
          </el-select>
        </el-form-item>
        <el-form-item label="负责人" prop="principal">
          <el-select v-model="formData.principal" placeholder="请选择" style="width:100%" clearable @visible-change="setPrincipalOptions">
            <el-option v-for="(item,key) in principalOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
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
  createBusiness,
  deleteBusiness,
  deleteBusinessByIds,
  updateBusiness,
  findBusiness,
  getBusinessList
} from '@/api/business' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
import { getProjectInformationList } from '@/api/projectInformation'
import { getUserList } from '@/api/user'
import { createCustomer, getCustomerList } from '@/api/customerData'
import * as XLSX from 'xlsx';
import FileSaver from 'file-saver'
export default {
  name: 'Business',
  mixins: [infoList],
  data() {
    return {
      listApi: getBusinessList,
      dialogFormVisible: false,
      type: '',
      deleteVisible: false,
      multipleSelection: [],
      SourceOptions: [],
      projectOptions: [],
      principalOptions: [],
      telephoneNumberOptions: [],
      nameOptions: [],
      companyOptions: [],
      telephoneNumberList: [],
      tTelephoneNumber: '',
      formData: {
        name: '',
        telephoneNumber: '',
        company: '',
        duty: '',
        source: undefined,
        project: '',
        principal: '',
      },
      rules: {
        name:                   [{ required: true, message: '请输入姓名',  trigger: 'blur' }],
        telephoneNumber:        [{ required: true, message: '请输入电话号码',  trigger: 'blur' }],
        source:                 [{ required: true, message: '请输入商机来源',  trigger: 'blur' }],
        project:                [{ required: true, message: '请输入意向项目',  trigger: 'blur' }],
        principal:              [{ required: true, message: '请输入负责人',  trigger: 'blur' }],
      }
    }
  },
  async created() {
    await this.getTableData()
    await this.getDict('Source')
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
    async hintChangeToCustomer(row) {
      this.$confirm('确定要转换客户吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.changeToCustomer(row)
      })
    },
    async changeToCustomer(row) {
      const res1 = await findBusiness({ ID: row.ID })
      if (res1.code === 0) {
        this.formData = res1.data.rebusiness
      }
      const res = await getCustomerList({ page: 1, pageSize: 999 })
      let flag = 0
      res.data.list && res.data.list.forEach(item => {
        if(item.telephone === this.formData.telephoneNumber) {
          flag = 1
        }
      })

      if (flag) {
        this.$message({
          type: 'warning',
          message: '电话号码重复'
        })
      }
      else {
        let temp  = {
          name: this.formData.company==="" ? this.formData.name : this.formData.company,
          type: this.formData.company==="" ? 0 : 1,
          linkman: this.formData.name,
          iDNumber: '',
          address: '',
          telephone: this.formData.telephoneNumber,
          invoice: '',
          bank: '',
          account: '',
          remark: this.formData.remark,
          audit: 1,
          principal: this.formData.principal,
        }
        let res = await createCustomer(temp)
        if (res.code === 0) {
          this.$message({
            type: 'success',
            message: '转换客户成功'
          })
        }
        else {
          this.$message({
            type: 'warning',
            message: '转换客户失败'
          })
        }
      }
    },
    async setCompanyOptions() {
      let list = []
      const res = await getBusinessList({ page: 1, pageSize: 999 })
      if(this.searchInfo.company === undefined || this.searchInfo.company === '') {
        res.data.list && res.data.list.forEach(item => {
          if(item.company != '')
            list.push(item.company)
        })
      }
      else {
        res.data.list && res.data.list.forEach(item => {
          if(item.company === this.searchInfo.company){
            if(item.company != '')
              list.push(item.company)          
          }
        })
      }
      let newArr = list.filter((item, index) => list.indexOf(item) === index); 

      this.companyOptions = []
      newArr && newArr.forEach(item => {
        const option ={
          label: item,
          value: item
        }
        this.companyOptions.push(option)
      })
    },
    async setNameOptions() {
      let list = []
      const res = await getBusinessList({ page: 1, pageSize: 999 })
      if(this.searchInfo.name === undefined || this.searchInfo.name === '') {
        res.data.list && res.data.list.forEach(item => {
          if(item.name != '')
            list.push(item.name)
        })
      }
      else {
        res.data.list && res.data.list.forEach(item => {
          if(item.name === this.searchInfo.name){
            if(item.name != '')
              list.push(item.name)          
          }
        })
      }
      let newArr = list.filter((item, index) => list.indexOf(item) === index); 

      this.nameOptions = []
      newArr && newArr.forEach(item => {
        const option ={
          label: item,
          value: item
        }
        this.nameOptions.push(option)
      })
    },
    async setTelephoneNumberOptions(){
      this.telephoneNumberOptions = []
      const res = await getBusinessList({ page: 1, pageSize: 999 })
      res.data.list && res.data.list.forEach(item => {
        const option = {
          label: item.telephoneNumber,
          value: item.telephoneNumber
        }
        this.telephoneNumberOptions.push(option)
      })
    },
    async setTelephoneNumberList(){
      this.telephoneNumberList = []
      const res = await getBusinessList({ page: 1, pageSize: 999 })
      res.data.list && res.data.list.forEach(item => {
        this.telephoneNumberList.push(item.telephoneNumber)
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
        this.deleteBusiness(row)
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
      const res = await deleteBusinessByIds({ ids })
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
    async updateBusiness(row) {
      const res = await findBusiness({ ID: row.ID })
      this.type = 'update'
      if (res.code === 0) {
        this.formData = res.data.rebusiness
        this.dialogFormVisible = true
      }
      this.tTelephoneNumber = this.formData.telephoneNumber
    },
    closeDialog() {
      this.$refs.formData.resetFields()
      this.dialogFormVisible = false
      this.formData = {
        name: '',
        telephoneNumber: '',
        company: '',
        duty: '',
        source: undefined,
        project: '',
        principal: '',
      }
    },
    async deleteBusiness(row) {
      const res = await deleteBusiness({ ID: row.ID })
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
      await this.setTelephoneNumberList()
      if(this.type === 'update'){
        if(this.formData.name != "" && this.formData.telephoneNumber != "" && this.formData.source != undefined 
        && this.formData.project != undefined && this.formData.principal != "" ) {
          let index = -1
          for(let i = 0 ; i < this.telephoneNumberList.length; i++){
            if(this.formData.telephoneNumber === this.telephoneNumberList[i]){
              index = i
              break
            }
          }
          if(index != -1 && this.formData.telephoneNumber===this.tTelephoneNumber || index === -1){
            let res
            res = await updateBusiness(this.formData)
            if (res.code === 0) {
              this.$message({
                type: 'success',
                message: '更改成功'
              })
              this.closeDialog()
              this.getTableData()
            }
          }
          else{
            this.$message({
            type: 'warning',
            message: '电话号码重复'
            })
          }
        }
      }
      else{
        let flag = 0
        for(let i = 0 ; i < this.telephoneNumberList.length; i++){
            if(this.formData.telephoneNumber === this.telephoneNumberList[i]){
            flag = 1
            break
          }
        }
        if(this.formData.name != "" && this.formData.telephoneNumber != "" && this.formData.source != undefined 
        && this.formData.project != undefined && this.formData.principal != "" && flag!= 1) {
          let res
          res = await createBusiness(this.formData)
          if (res.code === 0) {
            this.$message({
              type: 'success',
              message: '创建成功'
            })
            this.closeDialog()
            this.getTableData()
          }
        }
        else{
          if (flag === 1){
            this.$message({
            type: 'warning',
            message: '电话号码重复'
            })
          }
          else{
            this.$message({
            type: 'warning',
            message: '请填写必填项'
            })
          }
        }
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
