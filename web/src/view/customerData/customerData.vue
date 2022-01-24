<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="商家名称">
          <el-select v-model="searchInfo.name" placeholder="请选择" style="width:100%" default-first-option clearable filterable @visible-change="setNameOptions">
            <el-option v-for="(item,key) in nameOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="商家性质">
          <el-select v-model="searchInfo.type" placeholder="请选择" style="width:100%" default-first-option clearable filterable>
            <el-option v-for="(item,key) in intermediaryTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="联系人">
          <el-select v-model="searchInfo.linkman" placeholder="请选择" style="width:100%" default-first-option clearable filterable @visible-change="setLinkmanOptions">
            <el-option v-for="(item,key) in linkmanOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="联系电话">
          <el-select v-model="searchInfo.telephone" placeholder="请选择" style="width:100%" default-first-option clearable filterable @visible-change="setTelephoneOptions">
            <el-option v-for="(item,key) in telephoneOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="审核状态">
          <el-select v-model="searchInfo.audit" placeholder="请选择" style="width:100%" clearable>
            <el-option v-for="(item,key) in auditTypeOptions" :key="key" :label="item.label" :value="item.value" />
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
            <el-button size="mini" type="primary" icon="plus" @click="openDialog();setAudit()">新增</el-button>
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
        :default-sort = "{prop: 'date', order: 'descending'}"
        row-key="ID"
        @selection-change="handleSelectionChange"
        @row-click="handleRowClick"
        >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" type="expand">
          <template #default="scope">
            <el-descriptions :column="5" border>
              <el-descriptions-item label="商家名称" label-align="center" align="left" width="10px" label-class-name="background: #E1F3D8;">
                {{ scope.row.name }}
              </el-descriptions-item>
              <el-descriptions-item label="商家性质" label-align="center" align="left" width="10px">
                {{ filterDict(scope.row.type,"intermediaryType") }}
              </el-descriptions-item>
              <el-descriptions-item label="联系人" label-align="center" align="left" width="10px">
                {{ scope.row.linkman }}
              </el-descriptions-item>
              <el-descriptions-item label="联系电话" label-align="center" align="left" width="80px">
                {{ scope.row.telephone }}
              </el-descriptions-item>
              <el-descriptions-item label="身份证号" label-align="center" align="left" width="80px">
                {{ scope.row.iDNumber }}
              </el-descriptions-item>
              <el-descriptions-item label="详细地址" label-align="center" align="left" width="80px">
                {{ scope.row.address }}
              </el-descriptions-item>
              <el-descriptions-item label="开票名称" label-align="center" align="left" width="80px">
                {{ scope.row.invoice }}
              </el-descriptions-item>
              <el-descriptions-item label="开户银行" label-align="center" align="left" width="80px">
                {{ scope.row.bank }}
              </el-descriptions-item>
              <el-descriptions-item label="开户账号" label-align="center" align="left" width="80px">
                {{ scope.row.account }}
              </el-descriptions-item>
              <el-descriptions-item label="负责人" label-align="center" align="left" width="80px">
                {{ scope.row.principal }}
              </el-descriptions-item>
              <el-descriptions-item label="备注" label-align="center" align="left" width="80px">
                {{ scope.row.remark }}
              </el-descriptions-item>
            </el-descriptions>
          </template>
        </el-table-column>
        <el-table-column align="left" label="商家名称" prop="name" width="120"/>
        <el-table-column align="left" label="商家性质" prop="type" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.type,"intermediaryType") }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="联系人" prop="linkman" width="120" />
        <el-table-column align="left" label="联系电话" prop="telephone" width="120" />
        <!-- <el-table-column align="left" label="身份证号" prop="iDNumber" width="120" />
        <el-table-column align="left" label="详细地址" prop="address" width="120" /> -->
        <!-- <el-table-column align="left" label="开票名称" prop="invoice" width="120" />
        <el-table-column align="left" label="开户银行" prop="bank" width="120" />
        <el-table-column align="left" label="开户账号" prop="account" width="120" />
        <el-table-column align="left" label="备注" prop="remark" width="120" /> -->
        <el-table-column align="left" label="审核状态" prop="audit" width="120">
            <template #default="scope">
              <el-tag v-if="scope.row.audit===0" class="ml-2" type="success">{{ filterDict(scope.row.audit,"auditType") }}</el-tag>
              <el-tag v-else class="ml-2" type="danger">{{ filterDict(scope.row.audit,"auditType") }}</el-tag>
            </template>
        </el-table-column>
        <el-table-column align="left" label="负责人" prop="principal" width="120" />
        <el-table-column align="left" label="创建日期" width="180" prop="date"  sortable>
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="text" icon="edit" size="small" class="table-button" @click="updateCustomer(scope.row);handleRowClick(scope.row)">变更</el-button>
            <el-button type="text" icon="delete" size="mini" @click="deleteRow(scope.row);handleRowClick(scope.row)">删除</el-button>
            <el-button v-if="scope.row.audit===0" type="text" icon="tools" size="small" @click="changeAuditType1(scope.row);handleRowClick(scope.row)">取消审核</el-button>
            <el-button v-else type="text" icon="tools" size="small" @click="changeAuditType(scope.row);handleRowClick(scope.row)">审核</el-button> 
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
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="新增客户">
      <el-form :model="formData" label-position="right" label-width="80px" ref="formData" :rules="rules">
        <el-form-item label="商家名称" prop="name">
          <el-input v-model.trim="formData.name" clearable placeholder="请输入" :disabled="this.formData.audit===0"/>
        </el-form-item>
        <el-form-item label="商家性质" prop="type">
          <el-select v-model="formData.type" placeholder="请选择" style="width:100%" clearable :disabled="this.formData.audit===0">
            <el-option v-for="(item,key) in intermediaryTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="联系人" prop="linkman">
          <el-input v-model.trim="formData.linkman" clearable placeholder="请输入" :disabled="this.formData.audit===0"/>
        </el-form-item>
        <el-form-item label="联系电话" prop="telephone">
          <el-input v-model.trim="formData.telephone" clearable placeholder="请输入" :disabled="this.formData.audit===0"/>
        </el-form-item>
        <el-form-item label="身份证号" prop="iDNumber">
          <el-input v-model.trim="formData.iDNumber" clearable placeholder="" :disabled="this.formData.audit===0"/>
        </el-form-item>
        <el-form-item label="详细地址" prop="address">
          <el-input v-model.trim="formData.address" clearable placeholder="" :disabled="this.formData.audit===0"/>
        </el-form-item>
        <el-form-item label="开票名称" prop="invoice">
          <el-input v-model.trim="formData.invoice" clearable placeholder="" :disabled="this.formData.audit===0"/>
        </el-form-item>
        <el-form-item label="开户银行" prop="bank">
          <el-input v-model.trim="formData.bank" clearable placeholder="" :disabled="this.formData.audit===0"/>
        </el-form-item>
        <el-form-item label="开户账号" prop="account">
          <el-input v-model.trim="formData.account" clearable placeholder="" :disabled="this.formData.audit===0"/>
        </el-form-item>
        <el-form-item label="备注" prop="remark">
          <el-input v-model.trim="formData.remark" clearable placeholder="" :disabled="this.formData.audit===0"/>
        </el-form-item>
        <el-form-item label="审核状态" prop="audit">
          <el-select v-model="formData.audit" placeholder="请选择" style="width:100%" clearable disabled>
            <el-option v-for="(item,key) in auditTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="负责人" prop="principal">
          <el-select v-model="formData.principal" placeholder="请选择" style="width:100%" default-first-option clearable filterable @visible-change="setPrincipalOptions" :disabled="this.formData.audit===0">
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
  createCustomer,
  deleteCustomer,
  deleteCustomerByIds,
  updateCustomer,
  findCustomer,
  getCustomerList
} from '@/api/customerData' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
import { getUserList } from '@/api/user'
export default {
  name: 'Customer',
  mixins: [infoList],
  data() {
    return {
      listApi: getCustomerList,
      dialogFormVisible: false,
      type: '',
      deleteVisible: false,
      multipleSelection: [],
      intermediaryTypeOptions: [],
      auditTypeOptions: [],
      nameOptions: [],
      linkmanOptions: [],      
      telephoneOptions: [],
      principalOptions: [],
      telephoneNumberList: [],
      tTelephone: "",
      formData: {
        name: '',
        type: undefined,
        linkman: '',
        iDNumber: '',
        address: '',
        telephone: '',
        invoice: '',
        bank: '',
        account: '',
        remark: '',
        audit: undefined,
        principal: '',
      },
      rules: {
        name:                   [{ required: true, message: '请输入商家名称',  trigger: 'blur' }],
        type:                   [{ required: true, message: '请输入商家性质',  trigger: 'blur' }],
        linkman:                [{ required: true, message: '请输入联系电话',  trigger: 'blur' }],
        telephone:              [{ required: true, message: '请输入联系人',  trigger: 'blur' }],
        audit:                  [{ required: true, message: '请输入审核状态',  trigger: 'blur' }],       
        principal:              [{ required: true, message: '请输入负责人',  trigger: 'blur' }],
      }
    }
  },
  async created() {
    await this.getTableData()
    await this.getDict('intermediaryType')
    await this.getDict('auditType')
  },
  methods: {
    handleRowClick(row) {
      row.expanded = !row.expanded;
      this.$refs.multipleTable.toggleRowExpansion(row, row.expanded);
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
      const res = await findCustomer({ ID: row.ID })
      if (res.code === 0) {
        this.formData = res.data.recustomerData
        this.formData.audit = this.formData.audit===0 ? 1 : 0
      }
      let flag = await updateCustomer(this.formData)
      if (flag.code === 0) {
        this.$message({
          type: 'success',
          message: '更改成功'
        })
        this.getTableData()
      }
    },
    setAudit() {
      this.formData.audit = 1
    },
    async setTelephoneNumberList(){
      this.telephoneNumberList = []
      const res = await getCustomerList({ page: 1, pageSize: 999 })
      res.data.list && res.data.list.forEach(item => {
        this.telephoneNumberList.push(item.telephone)
      })
    },
    async setLinkmanOptions() {
      let list = []
      const res = await getCustomerList({ page: 1, pageSize: 999 })
      if(this.searchInfo.linkman === undefined || this.searchInfo.linkman === '') {
        res.data.list && res.data.list.forEach(item => {
          if(item.linkman != '')
            list.push(item.linkman)
        })
      }
      else {
        res.data.list && res.data.list.forEach(item => {
          if(item.linkman === this.searchInfo.linkman){
            if(item.linkman != '')
              list.push(item.linkman)          
          }
        })
      }
      let newArr = list.filter((item, index) => list.indexOf(item) === index); 

      this.linkmanOptions = []
      newArr && newArr.forEach(item => {
        const option ={
          label: item,
          value: item
        }
        this.linkmanOptions.push(option)
      })
    },
    async setTelephoneOptions() {
      this.telephoneOptions = []
      const res = await getCustomerList({ page: 1, pageSize: 999 })
      res.data.list && res.data.list.forEach(item => {
        const option = {
          label: item.telephone,
          value: item.telephone
        }
        this.telephoneOptions.push(option)
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
    async setNameOptions() {
      let list = []
      const res = await getCustomerList({ page: 1, pageSize: 999 })
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
        this.deleteCustomer(row)
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
      const res = await deleteCustomerByIds({ ids })
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
    async updateCustomer(row) {
      const res = await findCustomer({ ID: row.ID })
      this.type = 'update'
      if (res.code === 0) {
        this.formData = res.data.recustomerData
        this.dialogFormVisible = true
      }
      this.tTelephone = this.formData.telephone
    },
    closeDialog() {
      this.$refs.formData.resetFields()
      this.dialogFormVisible = false
      this.formData = {
        name: '',
        type: undefined,
        linkman: '',
        iDNumber: '',
        address: '',
        telephone: '',
        invoice: '',
        bank: '',
        account: '',
        remark: '',
        audit: undefined,
        principal: '',
      }
    },
    async deleteCustomer(row) {
      const res = await deleteCustomer({ ID: row.ID })
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
        if(this.formData.name != "" && this.formData.linkman != "" && this.formData.telephone != "" && this.formData.audit != undefined 
        && this.formData.principal != "") {
          let index = -1
          for(let i = 0 ; i < this.telephoneNumberList.length; i++){
            if(this.formData.telephone === this.telephoneNumberList[i]){
              index = i
              break
            }
          }
          if(index != -1 && this.formData.telephone===this.tTelephone || index === -1){
            let res
            res = await updateCustomer(this.formData)
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
            if(this.formData.telephone === this.telephoneNumberList[i]){
            flag = 1
            break
          }
        }
        if(this.formData.name != "" && this.formData.linkman != "" && this.formData.telephone != "" && this.formData.audit != undefined 
        && this.formData.principal != "" && flag!= 1) {
          let res
          res = await createCustomer(this.formData)
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
