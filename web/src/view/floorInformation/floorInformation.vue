<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="所属项目">
          <el-select v-model="searchInfo.project" placeholder="请选择" style="width:100%" default-first-option clearable filterable @change="getSearchBuildingOptions();setSearchBuild();getSearchFloorOptions()" @visible-change="getSearchBuildingOptions">
            <el-option v-for="(item,key) in projectOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="所属楼栋">
          <el-select v-model="searchInfo.build" placeholder="请选择" style="width:100%" default-first-option clearable filterable @change="getSearchBuildingOptions();setSearchFloor();getSearchFloorOptions()" @visible-change="getSearchBuildingOptions">
            <el-option v-for="(item,key) in searchBuildingOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="楼层名称">
          <el-select v-model="searchInfo.name" placeholder="请选择" style="width:100%" default-first-option clearable filterable @change="getSearchBuildingOptions();getSearchFloorOptions()" @visible-change="getSearchBuildingOptions">
            <el-option v-for="(item,key) in searchFloorOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="楼层状态">
          <el-select v-model="searchInfo.floorState" placeholder="请选择" style="width:100%" default-first-option clearable filterable >
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
            <el-button size="mini" type="primary" icon="plus" @click="openDialog();getFloorList()">新增</el-button>
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
        <el-table-column align="left" label="所属楼栋" prop="build" width="120" />
        <el-table-column align="left" label="楼层名称" prop="name" width="120" />
        <el-table-column align="left" label="楼层状态" prop="floorState" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.floorState,"buildState") }}
            </template>
        </el-table-column>
        <el-table-column align="right" label="建筑面积（㎡)" prop="coveredArea" width="120" :formatter="rounding"/>
        <el-table-column align="right" label="经营面积（㎡）" prop="operatingArea" width="120" :formatter="rounding"/>
        <el-table-column align="left" label="日期" width="180" prop="date" sortable>
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="text" icon="edit" size="small" class="table-button" @click="updateFloorInformation(scope.row)">变更</el-button>
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
      <el-form :model="formData" label-position="right" label-width="120px" ref="floorForm" :rules="rules">
        <el-form-item label="所属项目" prop="project">
          <el-select v-model="formData.project" placeholder="请选择" style="width:100%" default-first-option clearable filterable  @change="getSearchBuildingOptions();setFormDataBuild();getFloorList()" @visible-change="getFormDataBuildingOptions">
            <el-option v-for="(item,key) in projectOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="所属楼栋" prop="build">
          <el-select v-model="formData.build" placeholder="请选择" style="width:100%" default-first-option clearable filterable @change="getFormDataBuildingOptions();getFloorList();getFloorList()" @visible-change="getFormDataBuildingOptions">
            <el-option v-for="(item,key) in formDataBuildingOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="楼层名称" prop="name">
          <el-input v-model="formData.name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="楼层状态" prop="floorState">
          <el-select v-model="formData.floorState" placeholder="请选择" style="width:100%" clearable>
            <el-option v-for="(item,key) in buildStateOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="建筑面积（㎡）" prop="coveredArea">
          <el-input-number v-model="formData.coveredArea"  style="width:100%" :precision="2" clearable />
        </el-form-item>
        <el-form-item label="经营面积（㎡）" prop="operatingArea">
          <el-input-number v-model="formData.operatingArea"  style="width:100%" :precision="2" clearable />
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
  createFloorInformation,
  deleteFloorInformation,
  deleteFloorInformationByIds,
  updateFloorInformation,
  findFloorInformation,
  getFloorInformationList
} from '@/api/floorInformation' //  此处请自行替换地址
import { getProjectInformationList } from '@/api/projectInformation'
import { getBuildingInformationList } from '@/api/buildingInformation' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'FloorInformation',
  mixins: [infoList],
  data() {
    return {
      listApi: getFloorInformationList,
      dialogFormVisible: false,
      type: '',
      tFloor: '',
      deleteVisible: false,
      searchBuildingOptions: [],
      formDataBuildingOptions: [],
      multipleSelection: [],
      buildStateOptions: [],
      projectOptions: [],
      searchFloorOptions: [],
      floorList: [],
      formData: {
        project: '',
        build: '',
        name: '',
        floorState: undefined,
        coveredArea: 0,
        operatingArea: 0,
      },
      rules: {
        project:                   [{ required: true, message: '请输入所属项目',  trigger: 'blur' }],
        build:                     [{ required: true, message: '请输入所属楼栋',  trigger: 'blur' }],
        name:                      [{ required: true, message: '请输入楼层名称',  trigger: 'blur' }],
        floorState:                [{ required: true, message: '请输入楼层状态',  trigger: 'blur' }],
        coveredArea:               [{ required: true, message: '请输入建筑面积（㎡)',  trigger: 'blur' }],
        operatingArea:             [{ required: true, message: '请输入经营面积（㎡）',  trigger: 'blur' }],
      }
    }
  },
  async created() {
    await this.getTableData()
    await this.getDict('buildState')
    await this.setProjectOptions()
    await this.getSearchBuildingOptions()
    await this.getSearchFloorOptions()
  },
  methods: {
    setSearchFloor() {
      this.searchInfo.name = ''
    },
    async getSearchFloorOptions() {
      let searchFloorList = []
      const res = await getFloorInformationList({ page: 1, pageSize: 999 })

      if(this.searchInfo.project === undefined || this.searchInfo.project === '') {
        res.data.list && res.data.list.forEach(item => {
          searchFloorList.push(item.name)
        })
      }
      else {
        if(this.searchInfo.build === undefined || this.searchInfo.build === ''){
          res.data.list && res.data.list.forEach(item => {
            if(item.project === this.searchInfo.project)
              searchFloorList.push(item.name)
          })
        }
        else{
          res.data.list && res.data.list.forEach(item => {
            if(item.project === this.searchInfo.project && item.build === this.searchInfo.build)
              searchFloorList.push(item.name)
          })
        }
      }

      let newArr = searchFloorList.filter((item, index) => searchFloorList.indexOf(item) === index)
      this.searchFloorOptions = []
      newArr && newArr.forEach(item => {
        const option ={
          label: item,
          value: item
        }
        this.searchFloorOptions.push(option)
      })
    },
    async getFloorList() {
      this.floorList = []
      const res = await getFloorInformationList({ page: 1, pageSize: 999 })
      res.data.list && res.data.list.forEach(item => {
        if(item.project === this.formData.project && item.build === this.formData.build){
          this.floorList.push(item.name)
        }
      })
      console.log(this.floorList)
    },
    setFormDataBuild() {
      this.formData.build = ''
    },
    setSearchBuild() {
      this.searchInfo.build = ''
    },
    rounding(row, column) {
      return parseFloat(row[column.property]).toFixed(2)
    },
    async getSearchBuildingOptions(){
      let list = []
      const res = await getBuildingInformationList({ page: 1, pageSize: 999 })
      if(this.searchInfo.project === undefined || this.searchInfo.project === '') {
        res.data.list && res.data.list.forEach(item => {
          list.push(item.name)
        })
      }
      else {
        res.data.list && res.data.list.forEach(item => {
          if(item.project === this.searchInfo.project){
            list.push(item.name)
          }
        })
      }
      let newArr = list.filter((item, index) => list.indexOf(item) === index); 

      this.searchBuildingOptions = []
      newArr && newArr.forEach(item => {
        const option ={
          label: item,
          value: item
        }
        this.searchBuildingOptions.push(option)
      })
    },
    async getFormDataBuildingOptions(){
      let list = []
      const res = await getBuildingInformationList({ page: 1, pageSize: 999 })
      if(this.formData.project === undefined || this.formData.project === '') {
        res.data.list && res.data.list.forEach(item => {
          list.push(item.name)
        })
      }
      else {
        res.data.list && res.data.list.forEach(item => {
          if(item.project === this.formData.project){
            list.push(item.name)
          }
        })
      }
      let newArr = list.filter((item, index) => list.indexOf(item) === index); 

      this.formDataBuildingOptions = []
      newArr && newArr.forEach(item => {
        const option ={
          label: item,
          value: item
        }
        this.formDataBuildingOptions.push(option)
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
        this.deleteFloorInformation(row)
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
      const res = await deleteFloorInformationByIds({ ids })
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
    async updateFloorInformation(row) {
      const res = await findFloorInformation({ ID: row.ID })
      this.type = 'update'
      if (res.code === 0) {
        this.formData = res.data.refloorInformation
        this.dialogFormVisible = true
      }
      this.tFloor = this.formData.name
    },
    closeDialog() {
      this.$refs.floorForm.resetFields()
      this.dialogFormVisible = false
      this.formData = {
        project: '',
        build: '',
        name: '',
        floorState: undefined,
        coveredArea: 0,
        operatingArea: 0,
      }
    },
    async deleteFloorInformation(row) {
      const res = await deleteFloorInformation({ ID: row.ID })
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
        await this.getFloorList()
        if(this.formData.project != undefined && this.formData.build != undefined && this.formData.name != "" && 
        this.formData.floorState != undefined && this.formData.coveredArea != undefined && this.formData.operatingArea != undefined) {
          let index = -1
          for(let i = 0 ; i < this.floorList.length; i++){
            if(this.formData.name === this.floorList[i]){
              index = i
              break
            }
          }
          if(index != -1 && this.formData.name===this.tFloor || index === -1){
            let res
            res = await updateFloorInformation(this.formData)
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
            message: '楼层名称重复'
            })
          }
        }
      }
      else {
        let flag = 0
        for(let i = 0 ; i < this.floorList.length; i++){
            if(this.formData.name === this.floorList[i]){
            flag = 1
            break
          }
        }
        if(this.formData.project != undefined && this.formData.build != undefined && this.formData.name != "" && 
        this.formData.floorState != undefined && this.formData.coveredArea != undefined && this.formData.operatingArea != undefined 
        && flag != 1){
          let res
          res = await createFloorInformation(this.formData)
          
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
            message: '楼层名称重复'
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
