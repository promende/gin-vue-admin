<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="中介公司名称">
          <el-select v-model="searchInfo.name" placeholder="请选择" style="width:100%" default-first-option clearable filterable @visible-change="setNameOptions">
            <el-option v-for="(item,key) in nameOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="联系电话">
          <el-select v-model="searchInfo.telephoneNumber" placeholder="请选择" style="width:100%" default-first-option clearable filterable @visible-change="setTelephoneNumberOptions">
            <el-option v-for="(item,key) in telephoneNumberOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="创建人">
          <el-select v-model="searchInfo.creator" placeholder="请选择" style="width:100%" default-first-option clearable filterable @visible-change="setPrincipalOptions">
            <el-option v-for="(item,key) in principalOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="审核状态">
          <el-select v-model="searchInfo.auditType" placeholder="请选择" style="width:100%" default-first-option clearable filterable>
            <el-option v-for="(item,key) in auditTypeOptions" :key="key" :label="item.label" :value="item.value" />
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
            <el-button size="mini" type="primary" icon="plus" @click="openDialog();setCreator();setAuditType();">新增</el-button>
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
        >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="中介公司名称" prop="name" width="200" />
        <el-table-column align="left" label="联系电话" prop="telephoneNumber" width="120" />
        <el-table-column align="left" label="社会信用代码" prop="code" width="170" />
        <el-table-column align="left" label="备注" prop="remark" width="200" />
        <el-table-column align="left" label="审核状态" prop="auditType" width="120">
            <template #default="scope">
              <el-tag v-if="scope.row.auditType===0" class="ml-2" type="success">{{ filterDict(scope.row.auditType,"auditType") }}</el-tag>
              <el-tag v-else class="ml-2" type="danger">{{ filterDict(scope.row.auditType,"auditType") }}</el-tag>
            </template>
        </el-table-column>
        <el-table-column align="left" label="创建人" prop="creator" width="120" />
        <el-table-column align="left" label="创建日期" width="180" prop="date"  sortable>
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="text" icon="edit" size="small" class="table-button" @click="updateIntermediaryCompany(scope.row)">变更</el-button>
            <el-button type="text" icon="delete" size="small" @click="deleteRow(scope.row)">删除</el-button>
            <el-button v-if="scope.row.auditType===0" type="text" icon="tools" size="small" @click="changeAuditType1(scope.row)">取消审核</el-button>
            <el-button v-else type="text" icon="tools" size="small" @click="changeAuditType(scope.row)">审核</el-button>      
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
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="新增中介公司">
      <el-form :model="formData" label-position="right" label-width="110px" ref="formData" :rules="rules">
        <el-form-item label="中介公司名称" prop="name">
          <el-input v-model="formData.name" clearable placeholder="请输入" :disabled="this.formData.auditType===0"/>
        </el-form-item>
        <el-form-item label="联系电话">
          <el-input v-model="formData.telephoneNumber" clearable placeholder="" :disabled="this.formData.auditType===0"/>
        </el-form-item>
        <el-form-item label="社会信用代码">
          <el-input v-model="formData.code" clearable placeholder="" :disabled="this.formData.auditType===0"/>
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="formData.remark" clearable placeholder="" :disabled="this.formData.auditType===0"/>
        </el-form-item>
        <el-form-item label="审核状态" prop="auditType">
          <el-select v-model="formData.auditType" placeholder="请选择" style="width:100%" clearable disabled>
            <el-option v-for="(item,key) in auditTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="创建人" prop="creator">
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
  createIntermediaryCompany,
  deleteIntermediaryCompany,
  deleteIntermediaryCompanyByIds,
  updateIntermediaryCompany,
  findIntermediaryCompany,
  getIntermediaryCompanyList
} from '@/api/intermediaryCompany' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
import { getUserList } from '@/api/user'
import { mapGetters, mapActions } from 'vuex'
export default {
  name: 'IntermediaryCompany',
  mixins: [infoList],
  data() {
    return {
      listApi: getIntermediaryCompanyList,
      dialogFormVisible: false,
      type: '',
      deleteVisible: false,
      multipleSelection: [],
      auditTypeOptions: [],
      principalOptions: [],
      nameOptions: [],
      telephoneNumberOptions: [],
      codeOptions: [],
      nameList: [],
      tName: "",
      formData: {
        name: '',
        telephoneNumber: '',
        code: '',
        creator: '',
        auditType: undefined,
        remark: '',
      },
      rules: {
        name:                   [{ required: true, message: '请输入中介公司名称',  trigger: 'blur' }],
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
    async setNameList() {
      this.nameList = []
      const res = await getIntermediaryCompanyList({ page: 1, pageSize: 999 })
      res.data.list && res.data.list.forEach(item => {
        this.nameList.push(item.name)
      })
    },
    setAuditType() {
      this.formData.auditType = 1
    },
    setCreator() {
      this.formData.creator = this.userInfo.nickName
    },
    async setNameOptions() {
      this.nameOptions = []
      const res = await getIntermediaryCompanyList({ page: 1, pageSize: 999 })
      res.data.list && res.data.list.forEach(item => {
        const option = {
          label: item.name,
          value: item.name
        }
        this.nameOptions.push(option)
      })
    },
    async setTelephoneNumberOptions() {
      let list = []
      const res = await getIntermediaryCompanyList({ page: 1, pageSize: 999 })
      if(this.searchInfo.telephoneNumber === undefined || this.searchInfo.telephoneNumber === '') {
        res.data.list && res.data.list.forEach(item => {
          if(item.telephoneNumber != '')
            list.push(item.telephoneNumber)
        })
      }
      else {
        res.data.list && res.data.list.forEach(item => {
          if(item.telephoneNumber === this.searchInfo.telephoneNumber){
            if(item.telephoneNumber != '')
              list.push(item.telephoneNumber)          
          }
        })
      }
      let newArr = list.filter((item, index) => list.indexOf(item) === index); 

      this.telephoneNumberOptions = []
      newArr && newArr.forEach(item => {
        const option ={
          label: item,
          value: item
        }
        this.telephoneNumberOptions.push(option)
      })
    },
    async setCodeOptions() {
      this.codeOptions = []
      const res = await getIntermediaryCompanyList({ page: 1, pageSize: 999 })
      res.data.list && res.data.list.forEach(item => {
        if (item.code != "" && item.code != undefined) {
          const option = {
            label: item.code,
            value: item.code
          }
          this.codeOptions.push(option)
        }
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
        this.deleteIntermediaryCompany(row)
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
      const res = await deleteIntermediaryCompanyByIds({ ids })
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
      const res = await findIntermediaryCompany({ ID: row.ID })
      if (res.code === 0) {
        this.formData = res.data.reintermediaryCompany
        this.formData.auditType = this.formData.auditType===0 ? 1 : 0
      }
      let flag = await updateIntermediaryCompany(this.formData)
      if (flag.code === 0) {
        this.$message({
          type: 'success',
          message: '更改成功'
        })
        this.getTableData()
      }
    },
    async updateIntermediaryCompany(row) {
      const res = await findIntermediaryCompany({ ID: row.ID })
      this.type = 'update'
      if (res.code === 0) {
        this.formData = res.data.reintermediaryCompany
        this.dialogFormVisible = true
      }
      this.tName = this.formData.name
    },
    closeDialog() {
      this.$refs.formData.resetFields()
      this.dialogFormVisible = false
      this.formData = {
        name: '',
        telephoneNumber: '',
        code: '',
        creator: '',
        auditType: undefined,
        remark: '',
      }
    },
    async deleteIntermediaryCompany(row) {
      const res = await deleteIntermediaryCompany({ ID: row.ID })
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
      await this.setNameList()
      if(this.type === 'update'){
        if(this.formData.name != "" && this.formData.auditType != undefined) {
          let index = -1
          for(let i = 0 ; i < this.nameList.length; i++){
            if(this.formData.name === this.nameList[i]){
              index = i
              break
            }
          }
          if(index != -1 && this.formData.name===this.tName || index === -1){
            let res = await updateIntermediaryCompany(this.formData)
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
            message: '中介公司名称重复'
            })
          }
        }
      }
      else{
        let flag = 0
        for(let i = 0 ; i < this.nameList.length; i++){
            if(this.formData.name === this.nameList[i]){
            flag = 1
            break
          }
        }
        if(this.formData.name != "" && this.formData.auditType != undefined && flag!= 1) {
          let res
          res = await createIntermediaryCompany(this.formData)
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
            message: '中介公司名称重复'
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
