<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="所属项目">
          <el-select v-model="searchInfo.project" placeholder="请选择" style="width:100%" default-first-option clearable filterable @change="getSearchBuildingOptions();setSearchName()" @visible-change="setProjectOptions">
            <el-option v-for="(item,key) in projectOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="楼栋名称">
          <el-select v-model="searchInfo.name" placeholder="请选择" style="width:100%" default-first-option clearable filterable @change="getSearchBuildingOptions" @visible-change="getSearchBuildingOptions">
            <el-option v-for="(item,key) in searchBuildingOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="楼栋状态">
          <el-select v-model="searchInfo.buildState" placeholder="请选择" style="width:100%" default-first-option clearable filterable >
            <el-option v-for="(item,key) in buildStateOptions" :key="key" :label="item.label" :value="item.value" />
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
            <el-button size="mini" type="primary" icon="plus" @click="openDialog();getBuildingList();setProjectOptions();">新增</el-button>
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
        
        <el-table-column align="left" label="所属项目" prop="project" width="120" />
        <el-table-column align="left" label="楼栋名称" prop="name" width="120" />
        <el-table-column align="left" label="楼栋状态" prop="buildState" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.buildState,"buildState") }}
            </template>
        </el-table-column>
        <el-table-column align="right" label="建筑面积（㎡)" prop="coveredArea" width="120" :formatter="rounding"/>
        <el-table-column align="right" label="经营面积（㎡）" prop="businessArea" width="120" :formatter="rounding"/>
        <el-table-column align="right" label="楼上楼层数" prop="upstairs" width="120" />
        <el-table-column align="right" label="地下楼层数" prop="downstair" width="120" />
        <el-table-column align="left" label="日期" width="180" prop="date" sortable>
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="text" icon="edit" size="small" class="table-button" @click="updateBuildingInformation(scope.row)">变更</el-button>
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
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="新增楼栋">
      <el-form :model="formData" label-position="right" label-width="120px" ref="searchInfo" :rules="rules">
        <el-form-item label="所属项目" prop="project">
          <el-select v-model="formData.project" placeholder="请选择" style="width:100%" default-first-option clearable filterable @change="getBuildingList">
            <el-option v-for="(item,key) in projectOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="楼栋名称" prop="name">
          <el-input v-model.trim="formData.name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="楼栋状态" prop="buildState">
          <el-select v-model="formData.buildState" placeholder="请选择" style="width:100%" clearable>
            <el-option v-for="(item,key) in buildStateOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="建筑面积（㎡）" prop="coveredArea">
          <el-input-number v-model="formData.coveredArea"  style="width:100%" :precision="2" clearable />
        </el-form-item>
        <el-form-item label="经营面积（㎡）" prop="businessArea">
          <el-input-number v-model="formData.businessArea"  style="width:100%" :precision="2" clearable />
        </el-form-item>
        <el-form-item label="楼上楼层数" prop="upstairs">
          <el-input v-model.number="formData.upstairs" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="地下楼层数" prop="downstair">
          <el-input v-model.number="formData.downstair" clearable placeholder="请输入" />
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
  createBuildingInformation,
  deleteBuildingInformation,
  deleteBuildingInformationByIds,
  updateBuildingInformation,
  findBuildingInformation,
  getBuildingInformationList
} from '@/api/buildingInformation' //  此处请自行替换地址
import { getProjectInformationList } from '@/api/projectInformation'
import infoList from '@/mixins/infoList'
export default {
  name: 'BuildingInformation',
  mixins: [infoList],
  data() {
    return {
      listApi: getBuildingInformationList,
      dialogFormVisible: false,
      type: '',
      tBuilding: '',
      deleteVisible: false,
      multipleSelection: [],
      projectOptions: [],
      buildStateOptions: [],
      buildingList: [],
      searchBuildingOptions: [],
      projectOptionsInterval: '',
      buildingListInterval: '',
      formData: {
        project: '',
        name: '',
        buildState: undefined,
        coveredArea: 0,
        businessArea: 0,
        upstairs: 0,
        downstair: 0,
      },
      rules: {
        project:              [{ required: true, message: '请输入所属项目',  trigger: 'blur' }],
        name:                 [{ required: true, message: '请输入楼栋名称',  trigger: 'blur' }],
        buildState:           [{ required: true, message: '请输入楼栋状态',  trigger: 'blur' }],
        coveredArea:          [{ required: true, message: '请输入建筑面积（㎡)',  trigger: 'blur' }],
        businessArea:         [{ required: true, message: '请输入经营面积（㎡）',  trigger: 'blur' }],
        upstairs:             [{ required: true, message: '请输入楼上楼层数',  trigger: 'blur' }],
        downstair:            [{ required: true, message: '请输入地下楼层数',  trigger: 'blur' }],
      },
    }
  },
  async created() {
    await this.getTableData()
    await this.getDict('buildState')
    await this.setProjectOptions()
    await this.getBuildingList()
    await this.getSearchBuildingOptions()
  },
  methods: {
    setSearchName(){
      this.searchInfo.name = ''
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
    async getSearchBuildingOptions(){
      let searchBuildingList = []
      const res = await getBuildingInformationList({ page: 1, pageSize: 999 })
      if(this.searchInfo.project === undefined || this.searchInfo.project === '') {
        res.data.list && res.data.list.forEach(item => {
          searchBuildingList.push(item.name)
        })
      }
      else {
        res.data.list && res.data.list.forEach(item => {
          if(item.project === this.searchInfo.project){
            searchBuildingList.push(item.name)
          }
        })
      }
      let newArr = searchBuildingList.filter((item, index) => searchBuildingList.indexOf(item) === index);  
      searchBuildingList = newArr

      this.searchBuildingOptions = []
      newArr && newArr.forEach(item => {
        const option ={
          label: item,
          value: item
        }
        this.searchBuildingOptions.push(option)
      })
    },
    async getBuildingList() {
      this.buildingList = []
      const res = await getBuildingInformationList({ page: 1, pageSize: 999 })
      res.data.list && res.data.list.forEach(item => {
        if(item.project === this.formData.project){
          this.buildingList.push(item.name)
        }
      })
    },
    rounding(row, column) {
      return parseFloat(row[column.property]).toFixed(2)
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
        this.deleteBuildingInformation(row)
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
      const res = await deleteBuildingInformationByIds({ ids })
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
    async updateBuildingInformation(row) {
      const res = await findBuildingInformation({ ID: row.ID })
      this.type = 'update'
      if (res.code === 0) {
        this.formData = res.data.rebuildingInformation
        this.dialogFormVisible = true
      }
      this.tBuilding = this.formData.name
    },
    closeDialog() {
      this.dialogFormVisible = false
      this.formData = {
        project: '',
        name: '',
        buildState: undefined,
        coveredArea: 0,
        businessArea: 0,
        upstairs: 0,
        downstair: 0,
      }
    },
    async deleteBuildingInformation(row) {
      const res = await deleteBuildingInformation({ ID: row.ID })
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
      if(this.type === 'update') {
        if(this.formData.project != undefined &&  this.formData.name != "" && this.formData.buildState != undefined 
        && this.formData.upstairs != undefined && this.formData.downstair != undefined) {
          let index = -1
          for(let i = 0 ; i < this.buildingList.length; i++){
            if(this.formData.name === this.buildingList[i]){
              index = i
              break
            }
          }
          if(index != -1 && this.formData.name===this.tBuilding || index === -1){
            let res
            res = await updateBuildingInformation(this.formData)
            if (res.code === 0) {
              this.$message({
                type: 'success',
                message: '更改成功'
              })
              this.closeDialog()
              this.getTableData()
              this.getSearchBuildingOptions()
            }
          }
          else{
            this.$message({
            type: 'warning',
            message: '楼栋名称重复'
            })
          }
        }
        else{
          this.$message({
            type: 'warning',
            message: '请填写必填项'
            })
        }
      }
      else {
        let flag = 0
        for(let i = 0 ; i < this.buildingList.length; i++){
          if(this.formData.name === this.buildingList[i]){
            flag = 1
            break
          }
        }
        if(this.formData.project != undefined &&  this.formData.name != "" && this.formData.buildState != undefined && this.formData.upstairs != undefined 
          && this.formData.downstair != undefined && flag != 1){
          let res
            res = await createBuildingInformation(this.formData)
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
            message: '楼栋名称重复'
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
