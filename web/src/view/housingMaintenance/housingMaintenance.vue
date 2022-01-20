<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="所属项目">
          <el-select v-model="searchInfo.project" placeholder="请选择" style="width:100%" default-first-option clearable filterable @change="getSearchBuildingOptions();setSearchBuild();getSearchFloorOptions();setSearchHousing();" @visible-change="setProjectOptions();">
            <el-option v-for="(item,key) in projectOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="所属楼栋">
          <el-select v-model="searchInfo.build" placeholder="请选择" style="width:100%" default-first-option clearable filterable @change="getSearchFloorOptions();setSearchFloor();setSearchHousing();" @visible-change="getSearchBuildingOptions">
            <el-option v-for="(item,key) in searchBuildingOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="所属楼层">
          <el-select v-model="searchInfo.floor" placeholder="请选择" style="width:100%" default-first-option clearable filterable @change="getSearchHousingMainOptions();setSearchHousing();" @visible-change="getSearchFloorOptions();">
            <el-option v-for="(item,key) in searchFloorOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="房源名称">
          <el-select v-model="searchInfo.name" placeholder="请选择" style="width:100%" default-first-option clearable filterable @visible-change="getSearchHousingMainOptions();">
            <el-option v-for="(item,key) in searchBuildingMainOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="房源状态">
          <el-select v-model="searchInfo.state" placeholder="请选择" style="width:100%" default-first-option clearable filterable >
            <el-option v-for="(item,key) in buildStateOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="房源类型">
          <el-select v-model="searchInfo.type" placeholder="请选择" style="width:100%" default-first-option clearable filterable >
            <el-option v-for="(item,key) in HousingTypeOptions" :key="key" :label="item.label" :value="item.value" />
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
            <el-button size="mini" type="primary" icon="plus" @click="openDialog();setProjectOptions()">新增</el-button>
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
        <el-table-column align="left" label="所属楼层" prop="floor" width="120" />
        <el-table-column align="left" label="房源名称" prop="name" width="120" />
        <el-table-column align="left" label="房源状态" prop="state" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.state,"buildState") }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="房源类型" prop="type" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.type,"HousingType") }}
            </template>
        </el-table-column>
        <el-table-column align="right" label="计租面积（㎡)" prop="meterRentArea" width="120" :formatter="rounding"/>
        <el-table-column align="right" label="实用面积（㎡)" prop="usableArea" width="120" :formatter="rounding"/>
        <el-table-column align="left" label="创建日期" width="180" prop="date" sortable>
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="text" icon="edit" size="small" class="table-button" @click="updateHousingMaintenance(scope.row)">变更</el-button>
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
      <el-form :model="formData" label-position="right" label-width="120px" ref="formData" :rules="rules">
        <el-form-item label="所属项目" prop="project">
          <el-select v-model="formData.project" placeholder="请选择" style="width:100%" default-first-option clearable filterable @change="getFormDataBuildingOptions();setFormDataBuild();">
            <el-option v-for="(item,key) in projectOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="所属楼栋" prop="build">
          <el-select v-model="formData.build" placeholder="请选择" style="width:100%" default-first-option clearable filterable @change="getFormDataFloorOptions();setFormDataFloor()">
            <el-option v-for="(item,key) in formDataBuildingOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="所属楼层" prop="floor">
          <el-select v-model="formData.floor" placeholder="请选择" style="width:100%" default-first-option clearable filterable >
            <el-option v-for="(item,key) in formDataFloorOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="房源名称" prop="name">
          <el-input v-model="formData.name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="房源状态" prop="state">
          <el-select v-model="formData.state" placeholder="请选择" style="width:100%" clearable>
            <el-option v-for="(item,key) in buildStateOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="房源类型" prop="type">
          <el-select v-model="formData.type" placeholder="请选择" style="width:100%" clearable>
            <el-option v-for="(item,key) in HousingTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="计租面积（㎡)" prop="meterRentArea">
          <el-input-number v-model="formData.meterRentArea"  style="width:100%" :precision="2" clearable />
        </el-form-item>
        <el-form-item label="实用面积（㎡)" prop="usableArea">
          <el-input-number v-model="formData.usableArea"  style="width:100%" :precision="2" clearable />
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
  createHousingMaintenance,
  deleteHousingMaintenance,
  deleteHousingMaintenanceByIds,
  updateHousingMaintenance,
  findHousingMaintenance,
  getHousingMaintenanceList
} from '@/api/housingMaintenance' //  此处请自行替换地址
import { getProjectInformationList } from '@/api/projectInformation'
import { getBuildingInformationList } from '@/api/buildingInformation' 
import { getFloorInformationList } from '@/api/floorInformation'
import infoList from '@/mixins/infoList'
export default {
  name: 'HousingMaintenance',
  mixins: [infoList],
  data() {
    return {
      listApi: getHousingMaintenanceList,
      dialogFormVisible: false,
      type: '',
      tBuilding: '',
      deleteVisible: false,
      multipleSelection: [],
      buildStateOptions: [],
      HousingTypeOptions: [],
      projectOptions: [],
      searchBuildingOptions: [],
      formDataBuildingOptions: [],
      searchFloorOptions: [],
      searchBuildingMainOptions: [],
      formDataFloorOptions: [],
      housingList: [],
      formData: {
        project: '',
        build: '',
        floor: '',
        name: '',
        state: undefined,
        type: undefined,
        meterRentArea: 0,
        usableArea: 0,
      },
      rules: {
        project:                   [{ required: true, message: '请输入所属项目',  trigger: 'blur' }],
        build:                     [{ required: true, message: '请输入所属楼栋',  trigger: 'blur' }],
        floor:                     [{ required: true, message: '请输入所属楼层',  trigger: 'blur' }],
        name:                      [{ required: true, message: '请输入房源名称',  trigger: 'blur' }],
        state:                     [{ required: true, message: '请输入房源状态',  trigger: 'blur' }],
        type:                      [{ required: true, message: '请输入房源类型',  trigger: 'blur' }],
        meterRentArea:             [{ required: true, message: '请输入计租面积（㎡)',  trigger: 'blur' }],
        usableArea:                [{ required: true, message: '请输入实用面积（㎡）',  trigger: 'blur' }],
      }
    }
  },
  async created() {
    await this.getTableData()
    await this.getDict('buildState')
    await this.getDict('HousingType')
    await this.setProjectOptions()
    await this.getSearchBuildingOptions()
    await this.getFormDataBuildingOptions()
    await this.getSearchFloorOptions()
    await this.getFormDataFloorOptions()
    await this.getSearchHousingMainOptions()
  },
  methods: {
    rounding(row, column) {
      return parseFloat(row[column.property]).toFixed(2)
    },
    async getHousingList(){
      this.housingList = []
      const res = await getHousingMaintenanceList({ page: 1, pageSize: 999 })
      res.data.list && res.data.list.forEach(item => {
        if(item.project === this.formData.project && item.build === this.formData.build && item.floor === this.formData.floor){
          this.housingList.push(item.name)
        }
      })
    },
    async getFormDataFloorOptions() {
      let list = []
      const res = await getFloorInformationList({ page: 1, pageSize: 999 })

      if(this.formData.project === undefined || this.formData.project === '') {
        res.data.list && res.data.list.forEach(item => {
          list.push(item.name)
        })
      }
      else {
        if(this.formData.build === undefined || this.formData.build === ''){
          res.data.list && res.data.list.forEach(item => {
            if(item.project === this.formData.project)
              list.push(item.name)
          })
        }
        else{
          res.data.list && res.data.list.forEach(item => {
            if(item.project === this.formData.project && item.build === this.formData.build)
              list.push(item.name)
          })
        }
      }

      let newArr = list.filter((item, index) => list.indexOf(item) === index)
      this.formDataFloorOptions = []
      newArr && newArr.forEach(item => {
        const option ={
          label: item,
          value: item
        }
        this.formDataFloorOptions.push(option)
      })
    },
    async getSearchHousingMainOptions() {
      let list = []
      const res = await getHousingMaintenanceList({ page: 1, pageSize: 999 })

      if(this.searchInfo.project === undefined || this.searchInfo.project === '') {
        res.data.list && res.data.list.forEach(item => {
          list.push(item.name)
        })
      }
      else {
        if(this.searchInfo.build === undefined || this.searchInfo.build === ''){
          res.data.list && res.data.list.forEach(item => {
            if(item.project === this.searchInfo.project)
              list.push(item.name)
          })
        }
        else{
          if(this.searchInfo.floor === undefined || this.searchInfo.floor === ''){
            res.data.list && res.data.list.forEach(item => {
              if(item.project === this.searchInfo.project && item.build === this.searchInfo.build)
                list.push(item.name)
            })
          }
          else{
            res.data.list && res.data.list.forEach(item => {
              if(item.project === this.searchInfo.project && item.build === this.searchInfo.build && item.floor === this.searchInfo.floor)
                list.push(item.name)
            })
          }
        }
      }

      let newArr = list.filter((item, index) => list.indexOf(item) === index)
      this.searchBuildingMainOptions = []
      newArr && newArr.forEach(item => {
        const option ={
          label: item,
          value: item
        }
        this.searchBuildingMainOptions.push(option)
      })
      console.log(this.searchBuildingMainOptions)
    },
    async getSearchFloorOptions() {
      let list = []
      const res = await getFloorInformationList({ page: 1, pageSize: 999 })

      if(this.searchInfo.project === undefined || this.searchInfo.project === '') {
        res.data.list && res.data.list.forEach(item => {
          list.push(item.name)
        })
      }
      else {
        if(this.searchInfo.build === undefined || this.searchInfo.build === ''){
          res.data.list && res.data.list.forEach(item => {
            if(item.project === this.searchInfo.project)
              list.push(item.name)
          })
        }
        else{
          res.data.list && res.data.list.forEach(item => {
            if(item.project === this.searchInfo.project && item.build === this.searchInfo.build)
              list.push(item.name)
          })
        }
      }

      let newArr = list.filter((item, index) => list.indexOf(item) === index)
      this.searchFloorOptions = []
      newArr && newArr.forEach(item => {
        const option ={
          label: item,
          value: item
        }
        this.searchFloorOptions.push(option)
      })
    },
    setSearchHousing() {
      this.searchInfo.name = ''
    },
    setFormDataFloor() {
      this.formData.floor = ''
    },
    setSearchFloor() {
      this.searchInfo.floor = ''
    },
    setFormDataBuild() {
      this.formData.build = ''
    },
    setSearchBuild() {
      this.searchInfo.build = ''
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
        this.deleteHousingMaintenance(row)
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
      const res = await deleteHousingMaintenanceByIds({ ids })
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
    async updateHousingMaintenance(row) {
      const res = await findHousingMaintenance({ ID: row.ID })
      this.type = 'update'
      if (res.code === 0) {
        this.formData = res.data.rehousingMaintenance
        this.dialogFormVisible = true
      }
      this.tBuilding = this.formData.name
    },
    closeDialog() {
      this.$refs.formData.resetFields()
      this.dialogFormVisible = false
      this.formData = {
        project: '',
        build: '',
        floor: '',
        name: '',
        state: undefined,
        type: undefined,
        meterRentArea: 0,
        usableArea: 0,
      }
    },
    async deleteHousingMaintenance(row) {
      const res = await deleteHousingMaintenance({ ID: row.ID })
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
      await this.getHousingList()
      if(this.type === 'update'){
        if(this.formData.project != undefined && this.formData.build != undefined && this.formData.floor != undefined && this.formData.name != "" && 
        this.formData.state != undefined && this.formData.type != undefined && this.formData.meterRentArea != undefined && this.formData.usableArea != undefined) {
          let index = -1
          for(let i = 0 ; i < this.housingList.length; i++){
            if(this.formData.name === this.housingList[i]){
              index = i
              break
            }
          }
          if(index != -1 && this.formData.name===this.tBuilding || index === -1){
            let res
            res = await updateHousingMaintenance(this.formData)
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
            message: '房源名称重复'
            })
          }
        }
      }
      else {
        let flag = 0
        for(let i = 0 ; i < this.housingList.length; i++){
            if(this.formData.name === this.housingList[i]){
            flag = 1
            break
          }
        }
        if(this.formData.project != undefined && this.formData.build != undefined && this.formData.floor != undefined && this.formData.name != "" && 
        this.formData.state != undefined && this.formData.type != undefined && this.formData.meterRentArea != undefined && this.formData.usableArea != undefined
        && flag!= 1) {
          let res
          res = await createHousingMaintenance(this.formData)
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
            message: '房源名称重复'
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
