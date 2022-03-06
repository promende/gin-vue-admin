<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="所属公司">
          <el-select v-model="searchInfo.company" placeholder="请选择" style="width:100%" default-first-option clearable filterable @visible-change="setCompanyOptions">
            <el-option v-for="(item,key) in companyOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="联系人">
          <el-select v-model="searchInfo.name" placeholder="请选择" style="width:100%" default-first-option clearable filterable @visible-change="setNameOptions">
            <el-option v-for="(item,key) in nameOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="联系电话">
          <el-select v-model="searchInfo.telephoneNumber" placeholder="请选择" style="width:100%" default-first-option clearable filterable @visible-change="setTelephoneNumberOptions">
            <el-option v-for="(item,key) in telephoneNumberOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="审核状态">
          <el-select v-model="searchInfo.auditType" placeholder="请选择" style="width:100%" default-first-option clearable filterable>
            <el-option v-for="(item,key) in auditTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="创建人">
          <el-select v-model="searchInfo.auditType" placeholder="请选择" style="width:100%" default-first-option clearable filterable @visible-change="setPrincipalOptions">
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
        >
        <el-table-column align="left" label="所属公司" prop="company" width="200" />
        <el-table-column align="left" label="联系人" prop="name" width="120" />
        <el-table-column align="left" label="联系电话" prop="telephoneNumber" width="120" />
        <el-table-column align="left" label="备注" prop="remark" width="200" />
        <el-table-column align="left" label="审核状态" prop="auditType" width="120">
            <template #default="scope">
              <el-tag v-if="scope.row.auditType===0" class="ml-2" type="success">{{ filterDict(scope.row.auditType,"auditType") }}</el-tag>
              <el-tag v-else class="ml-2" type="danger">{{ filterDict(scope.row.auditType,"auditType") }}</el-tag>
            </template>
        </el-table-column>
        <el-table-column align="left" label="创建人" prop="creator" width="120" />
        <el-table-column align="left" label="创建日期" width="180" prop="date" sortable>
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="">
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
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="新增中介联系人">
      <el-form :model="formData" label-position="right" label-width="80px" ref="formData" :rules="rules">
        <el-form-item label="所属公司" prop="company">
          <el-select v-model="formData.company" placeholder="请选择" style="width:100%" default-first-option clearable filterable @visible-change="setCompanyOptions" :disabled="this.formData.auditType===0">
            <el-option v-for="(item,key) in companyOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="联系人" prop="name">
          <el-input v-model="formData.name" clearable placeholder="请输入" :disabled="this.formData.auditType===0"/>
        </el-form-item>
        <el-form-item label="联系电话" prop="telephoneNumber">
          <el-input v-model="formData.telephoneNumber" clearable placeholder="请输入" :disabled="this.formData.auditType===0"/>
        </el-form-item>
        <el-form-item label="备注" prop="remark">
          <el-input v-model="formData.remark" clearable placeholder="" :disabled="this.formData.auditType===0"/>
        </el-form-item>
        <el-form-item label="审核状态" prop="auditType">
          <el-select v-model="formData.auditType" placeholder="请选择" style="width:100%" clearable disabled>
            <el-option v-for="(item,key) in auditTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="创建人" prop="creator" >
          <el-input v-model="formData.creator" clearable placeholder="请输入" disabled/>
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
  createMiddleman,
  deleteMiddleman,
  deleteMiddlemanByIds,
  updateMiddleman,
  findMiddleman,
  getMiddlemanList
} from '@/api/middleman' //  此处请自行替换地址
import { getIntermediaryCompanyList } from '@/api/intermediaryCompany' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
import { getUserList } from '@/api/user'
import { mapGetters, mapActions } from 'vuex'
import * as XLSX from 'xlsx';
import FileSaver from 'file-saver'
export default {
  name: 'Middleman',
  mixins: [infoList],
  data() {
    return {
      listApi: getMiddlemanList,
      dialogFormVisible: false,
      type: '',
      deleteVisible: false,
      multipleSelection: [],
      auditTypeOptions: [],
      companyOptions: [],
      nameOptions: [],
      telephoneNumberOptions: [],
      principalOptions: [],
      telephoneNumberList: [],
      tTelephoneNumber: '',
      formData: {
        company: '',
        name: '',
        telephoneNumber: '',
        remark: '',
        auditType: undefined,
        creator: '',
      },
      rules: {
        company:                [{ required: true, message: '请输入中介公司名称',  trigger: 'blur' }],
        name:                   [{ required: true, message: '请输入联系人',  trigger: 'blur' }],
        telephoneNumber:        [{ required: true, message: '请输入电话号码',  trigger: 'blur' }],
        creator:                [{ required: true, message: '请输入中介公司名称',  trigger: 'blur' }],
        auditType:              [{ required: true, message: '请输入中介公司名称',  trigger: 'blur' }],
      }
    }
  },
  async created() {
    await this.getTableData()
    await this.getDict('auditType')
  },
  computed: {
    ...mapGetters('user', ['userInfo', 'sideMode', 'baseColor']),
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
          FileSaver.saveAs(new Blob([wbout], { type: 'application/octet-stream' }), '中介联系人台账表.xlsx');
      } catch (e)
      {
        if (typeof console !== 'undefined')
            console.log(e, wbout)
      }

      this.pageSize = oldPageSize
      await this.getTableData()
      
      return wbout
    },
    async setTelephoneNumberList() {
      this.telephoneNumberList = []
      const res = await getMiddlemanList({ page: 1, pageSize: 999 })
      res.data.list && res.data.list.forEach(item => {
        this.telephoneNumberList.push(item.telephoneNumber)
      })
    },
    setAuditType() {
      this.formData.auditType = 1
    },
    setCreator() {
      this.formData.creator = this.userInfo.nickName
    },
    changeAuditType(row) {
      this.$confirm('确定要审核吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.changeAuditType2(row)
      })
    },
    changeAuditType1(row) {
      this.$confirm('确定要取消审核吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.changeAuditType2(row)
      })
    },
    async changeAuditType2(row) {
      const res = await findMiddleman({ ID: row.ID })
      if (res.code === 0) {
        this.formData = res.data.remiddleman
        this.formData.auditType = this.formData.auditType===0 ? 1 : 0
      }
      let flag = await updateMiddleman(this.formData)
      if (flag.code === 0) {
        this.$message({
          type: 'success',
          message: '更改成功'
        })
        this.getTableData()
      }
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
    async setTelephoneNumberOptions() {
      this.telephoneNumberOptions = []
      const res = await getMiddlemanList({ page: 1, pageSize: 999 })
      res.data.list && res.data.list.forEach(item => {
        const option = {
          label: item.telephoneNumber,
          value: item.telephoneNumber
        }
        this.telephoneNumberOptions.push(option)
      })
    },
    async setNameOptions() {
      let list = []
      const res = await getMiddlemanList({ page: 1, pageSize: 999 })
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
    async setCompanyOptions() {
      this.companyOptions = []
      const res = await getIntermediaryCompanyList({ page: 1, pageSize: 999 })
      res.data.list && res.data.list.forEach(item => {
        const option = {
          label: item.name,
          value: item.name
        }
        this.companyOptions.push(option)
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
        this.deleteMiddleman(row)
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
      const res = await deleteMiddlemanByIds({ ids })
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
    async updateMiddleman(row) {
      const res = await findMiddleman({ ID: row.ID })
      this.type = 'update'
      if (res.code === 0) {
        this.formData = res.data.remiddleman
        this.dialogFormVisible = true
      }
      this.tTelephoneNumber = this.formData.telephoneNumber
    },
    closeDialog() {
      this.$refs.formData.resetFields()
      this.dialogFormVisible = false
      this.formData = {
        company: '',
        name: '',
        telephoneNumber: '',
        remark: '',
        auditType: undefined,
        creator: '',
      }
    },
    async deleteMiddleman(row) {
      const res = await deleteMiddleman({ ID: row.ID })
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
        if(this.formData.company != undefined && this.formData.name != "" && this.formData.telephoneNumber != "") {
          let index = -1
          for(let i = 0 ; i < this.telephoneNumberList.length; i++){
            if(this.formData.telephoneNumber === this.telephoneNumberList[i]){
              index = i
              break
            }
          }
          if(index != -1 && this.formData.telephoneNumber===this.tTelephoneNumber || index === -1){
            let res = await updateMiddleman(this.formData)
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
            message: '联系电话重复'
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
        if(this.formData.company != undefined && this.formData.name != "" && this.formData.telephoneNumber != "" && flag!= 1) {
          let res
          res = await createMiddleman(this.formData)
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
            message: '联系电话重复'
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
