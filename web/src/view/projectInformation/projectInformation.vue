<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="项目名称">
          <el-select v-model="searchInfo.name" placeholder="请选择" style="width:100%" clearable filterable>
            <el-option v-for="(item,key) in projectOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="项目简称">
          <el-input v-model="searchInfo.abbreviation" placeholder="支持模糊搜索" />
        </el-form-item>
        <el-form-item label="负责人">
          <el-select v-model="searchInfo.principal" placeholder="请选择" style="width:100%" clearable filterable @visible-change="setPrincipalOptions">
            <el-option v-for="(item,key) in principalOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="营运状态">
          <el-select v-model="searchInfo.operatingState" placeholder="请选择" style="width:100%" default-first-option clearable filterable >
            <el-option v-for="(item,key) in OperatingStateOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="经营类型">
          <el-select v-model="searchInfo.managementType" placeholder="请选择" style="width:100%" default-first-option clearable filterable >
            <el-option v-for="(item,key) in managementTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="物业管理类型">
          <el-select v-model="searchInfo.propertyManagementType" placeholder="请选择" style="width:100%" default-first-option clearable filterable >
            <el-option v-for="(item,key) in PropertyManagementTypeOptions" :key="key" :label="item.label" :value="item.value" />
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
            <el-button size="mini" type="primary" icon="plus" @click="openDialog();setPrincipalOptions();setProjectList()">新增</el-button>
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
        <el-table-column align="left" label="项目名称" prop="name" width="120" />
        <el-table-column align="left" label="项目简称" prop="abbreviation" width="120" />
        <el-table-column align="left" label="项目地址" prop="address" width="120" />
        <el-table-column align="left" label="营运状态" prop="operatingState" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.operatingState,"OperatingState") }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="经营类型" prop="managementType" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.managementType,"managementType") }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="物业管理类型" prop="propertyManagementType" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.propertyManagementType,"PropertyManagementType") }}
            </template>
        </el-table-column>
        <el-table-column align="right" label="建筑面积（㎡)" prop="coveredArea" width="120" :formatter="rounding"/>
        <el-table-column align="right" label="经营面积（㎡）" prop="operatingArea" width="120" :formatter="rounding"/>
        <el-table-column align="left" label="负责人" prop="principal" width="120" />
        <el-table-column align="left" label="创建日期" width="180" prop="date" sortable>
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="text" icon="edit" size="small" class="table-button" @click="updateProjectInformation(scope.row)">变更</el-button>
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
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="新增项目">
      <el-form ref="projectForm" :model="formData" label-position="right" label-width="120px" :rules="rules">
        <el-form-item label="项目名称" prop="name">
          <el-input v-model.trim="formData.name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="项目简称" prop="abbreviation">
          <el-input v-model.trim="formData.abbreviation" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="项目地址" prop="address">
          <el-input v-model.trim="formData.address" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="营运状态" prop="operatingState">
          <el-select v-model="formData.operatingState" placeholder="请选择" style="width:100%" clearable filterable>
            <el-option v-for="(item,key) in OperatingStateOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="经营类型" prop="managementType">
          <el-select v-model="formData.managementType" placeholder="请选择" style="width:100%" clearable filterable>
            <el-option v-for="(item,key) in managementTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="物业管理类型" prop="propertyManagementType">
          <el-select v-model="formData.propertyManagementType" placeholder="请选择" style="width:100%" clearable filterable>
            <el-option v-for="(item,key) in PropertyManagementTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="建筑面积（㎡）" prop="coveredArea">
          <el-input-number v-model="formData.coveredArea"  style="width:100%" :precision="2" clearable />
        </el-form-item>
        <el-form-item label="经营面积（㎡）" prop="operatingArea">
          <el-input-number v-model="formData.operatingArea"  style="width:100%" :precision="2" clearable />
        </el-form-item>
        <el-form-item label="负责人" prop="principal">
          <el-select v-model="formData.principal" placeholder="请选择" style="width:100%" clearable filterable>
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
  createProjectInformation,
  deleteProjectInformation,
  deleteProjectInformationByIds,
  updateProjectInformation,
  findProjectInformation,
  getProjectInformationList
} from '@/api/projectInformation' // 此处请自行替换地址
import infoList from '@/mixins/infoList'
import { getUserList } from '@/api/user'
export default {
  name: 'ProjectInformation',
  mixins: [infoList],
  data() {
    const checkName = (rule, value, callback) => {
      if (value.length === 0) {
        return callback(new Error('请输入项目名称'))
      } else {
        callback()
      }
    }
    return {
      listApi: getProjectInformationList,
      dialogFormVisible: false,
      type: '',
      deleteVisible: false,
      ProjectList: [],
      principalOptions: [],
      multipleSelection: [],
      OperatingStateOptions: [],
      managementTypeOptions: [],
      PropertyManagementTypeOptions: [],
      principalOptionsInterval: '',
      projectListInterval: '',
      formData: {
        name: '',
        abbreviation: '',
        address: '',
        operatingState: undefined,
        managementType: undefined,
        propertyManagementType: undefined,
        coveredArea: 0,
        operatingArea: 0,
        principal: '',
      },
      rules: {
        name:                   [{ required: true, message: '请输入项目名称',  trigger: 'blur' }],
        abbreviation:           [{ required: true, message: '请输入项目简称',  trigger: 'blur' }],
        address:                [{ required: true, message: '请输入项目地址',  trigger: 'blur' }],
        operatingState:         [{ required: true, message: '请输入营运状态',  trigger: 'blur' }],
        managementType:         [{ required: true, message: '请输入经营类型',  trigger: 'blur' }],
        propertyManagementType: [{ required: true, message: '请输入物业管理类型',  trigger: 'blur' }],
        coveredArea:            [{ required: true, message: '请输入建筑面积',  trigger: 'blur' }],
        operatingArea:          [{ required: true, message: '请输入经营面积',  trigger: 'blur' }],
        principal:              [{ required: true, message: '请输入负责人',  trigger: 'blur' }],
      },
    }
  },
  async created() {
    await this.getTableData()
    await this.getDict('OperatingState')
    await this.getDict('managementType')
    await this.getDict('PropertyManagementType')
    await this.setPrincipalOptions()
    await this.setProjectList()
    await this.setProjectOptions()
  },
  methods: {
    onReset() {
      this.searchInfo = {}
    },
    async setProjectList() {
      this.ProjectList = []
      const res = await getProjectInformationList({ page: 1, pageSize: 999 })
      res.data.list && res.data.list.forEach(item => {
        this.ProjectList.push(item.name)
      })
    },
    async setProjectOptions() {
      this.projectOptions = []
      const res = await getProjectInformationList({ page: 1, pageSize: 999 })
      this.getProjectOptions(res.data.list, this.projectOptions)
    },
    getProjectOptions(data, options) {
      data && data.forEach(item => {
        const option = {
          label: item.name,
          value: item.name  
        }
        options.push(option)
      })
    },
    async setPrincipalOptions() {
      this.principalOptions = []
      const res = await getUserList({ page: 1, pageSize: 999 })
      this.getPrincipalOptions(res.data.list, this.principalOptions)
    },
    getPrincipalOptions(data, options) {
      data && data.forEach(item => {
        const option = {
          label: item.nickName,
          value: item.nickName
        }
        options.push(option)
      })
    },
    rounding(row, column) {
      return parseFloat(row[column.property]).toFixed(2)
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
        this.deleteProjectInformation(row)
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
      const res = await deleteProjectInformationByIds({ ids })
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
    async updateProjectInformation(row) {
      const res = await findProjectInformation({ ID: row.ID })
      this.type = 'update'
      if (res.code === 0) {
        this.formData = res.data.reproject
        this.dialogFormVisible = true
      }
    },
    closeDialog() {
      this.$refs.projectForm.resetFields()
      this.dialogFormVisible = false
      this.formData = {
        name: '',
        abbreviation: '',
        address: '',
        operatingState: undefined,
        managementType: undefined,
        propertyManagementType: undefined,
        coveredArea: 0,
        operatingArea: 0,
        principal: '',
      }
    },
    async deleteProjectInformation(row) {
      const res = await deleteProjectInformation({ ID: row.ID })
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
      let flag = 0
      for(let i = 0 ; i < this.ProjectList.length; i++){
        if(this.formData.name === this.ProjectList[i]){
          flag = 1
          break
        }
      }
      if(this.formData.name != '' &&  this.formData.abbreviation != '' && this.formData.address != '' && 
        this.formData.operatingState != undefined && this,this.formData.managementType != undefined &&
        this.formData.propertyManagementType != undefined && this.formData.principal != '' && flag != 1){
        let res
        switch (this.type) {
          case 'create':
            res = await createProjectInformation(this.formData)
            if (res.code === 0) {
              this.$message({
                type: 'success',
                message: '创建成功'
              })
            }
            break
          case 'update':
            res = await updateProjectInformation(this.formData)
            if (res.code === 0) {
              this.$message({
                type: 'success',
                message: '更改成功'
              })
            }
            break
          default:
            res = await createProjectInformation(this.formData)
            if (res.code === 0) {
              this.$message({
                type: 'success',
                message: '创建成功'
              })
            }
            break
        }
        if (res.code === 0) {
          this.closeDialog()
          this.getTableData()
        }
      }
      else{
        if (flag === 1){
          this.$message({
          type: 'warning',
          message: '项目名称重复'
          })
        }
        else{
          this.$message({
          type: 'warning',
          message: '请填写必填项'
          })
        }
      }
    },
    openDialog() {
      this.type = 'create'
      this.dialogFormVisible = true
    },
  },
}
</script>

<style>
</style>
