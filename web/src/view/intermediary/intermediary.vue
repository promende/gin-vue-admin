<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="中介公司名称">
          <el-input v-model="searchInfo.name" placeholder="支持模糊查找" />
        </el-form-item>
        <el-form-item label="联系电话">
          <el-input v-model="searchInfo.telephoneNumber" placeholder="支持模糊查找" />
        </el-form-item>
        <el-form-item label="社会信用代码">
          <el-input v-model="searchInfo.code" placeholder="支持模糊查找" />
        </el-form-item>
        <el-form-item label="创建人">
          <el-select v-model="searchInfo.creator" placeholder="请选择" style="width:100%" default-first-option clearable filterable @visible-change="setPrincipalOptions">
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
            <el-button size="mini" type="primary" icon="plus" @click="openDialog();setCreator()">新增</el-button>
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
        
        <el-table-column align="left" label="中介公司名称" prop="name" width="300" />
        <el-table-column align="left" label="联系电话" prop="telephoneNumber" width="120" />
        <el-table-column align="left" label="社会信用代码" prop="code" width="180" />
        <el-table-column align="left" label="创建人" prop="creator" width="120" />
        <el-table-column align="left" label="备注" prop="remark" width="540" />
        <el-table-column align="left" label="创建日期" width="180" prop="date" sortable>
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="text" icon="edit" size="small" class="table-button" @click="updateIntermediary(scope.row)">变更</el-button>
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
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="新增中介公司">
      <el-form :model="formData" label-position="right" label-width="120px"  ref="formData" :rules="rules">
        <el-form-item label="中介名称" prop="name">
          <el-input v-model="formData.name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="联系电话" prop="telephoneNumber">
          <el-input v-model="formData.telephoneNumber" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="社会信用代码" prop="code">
          <el-input v-model="formData.code" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="创建人" prop="creator">
          <el-input v-model="formData.creator" clearable placeholder="请输入" :disabled="true"/>
        </el-form-item>
        <el-form-item label="备注" prop="remark">
          <el-input v-model="formData.remark" clearable placeholder="请输入" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="checkPhone();enterDialog();">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
import {
  createIntermediary,
  deleteIntermediary,
  deleteIntermediaryByIds,
  updateIntermediary,
  findIntermediary,
  getIntermediaryList
} from '@/api/intermediary' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
import { getUserList } from '@/api/user'
import { mapGetters, mapActions } from 'vuex'
export default {
  name: 'Intermediary',
  mixins: [infoList],
  data() {
    return {
      listApi: getIntermediaryList,
      dialogFormVisible: false,
      type: '',
      deleteVisible: false,
      multipleSelection: [],
      principalOptions: [],
      formData: {
        name: '',
        telephoneNumber: '',
        code: '',
        creator: '',
        remark: '',
      },
      rules: {
        name:                   [{ required: true, message: '请输入中介公司名称',  trigger: 'blur' }],
      },
      nameNull: 'false'
    }
  },
  computed: {
    ...mapGetters('user', ['userInfo', 'sideMode', 'baseColor']),
  },
  async created() {
    await this.getTableData()
  },
  methods: {
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
    ...mapActions('user', ['LoginOut', 'GetUserInfo']),
    setCreator(){
      this.formData.creator = this.userInfo.nickName
    },
    checkPhone(){
      
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
        this.deleteIntermediary(row)
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
      const res = await deleteIntermediaryByIds({ ids })
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
    async updateIntermediary(row) {
      const res = await findIntermediary({ ID: row.ID })
      this.type = 'update'
      if (res.code === 0) {
        this.formData = res.data.reintermediary
        this.dialogFormVisible = true
      }
    },
    closeDialog() {
      this.$refs.formData.resetFields()
      this.dialogFormVisible = false
      this.formData = {
        name: '',
        telephoneNumber: '',
        code: '',
        creator: '',
        remark: '',
      }
    },
    async deleteIntermediary(row) {
      const res = await deleteIntermediary({ ID: row.ID })
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
          res = await createIntermediary(this.formData)
          break
        case 'update':
          res = await updateIntermediary(this.formData)
          break
        default:
          res = await createIntermediary(this.formData)
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
