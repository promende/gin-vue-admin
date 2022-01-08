<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="所属公司">
          <el-input v-model="searchInfo.company" placeholder="支持模糊查找" />
        </el-form-item>
        <el-form-item label="联系人">
          <el-input v-model="searchInfo.name" placeholder="支持模糊查找" />
        </el-form-item>
        <el-form-item label="联系电话">
          <el-input v-model="searchInfo.telephoneNumber" placeholder="支持模糊查找" />
        </el-form-item>
        <el-form-item label="创建人">
          <el-input v-model="searchInfo.principal" placeholder="支持模糊查找" />
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="mini" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button size="mini" type="primary" icon="plus" @click="openDialog();setPrincipal()">新增</el-button>
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
        <el-table-column align="left" label="所属公司" prop="company" width="300" />
        <el-table-column align="left" label="联系人" prop="name" width="120" />
        <el-table-column align="left" label="联系电话" prop="telephoneNumber" width="120" />
        <el-table-column align="left" label="创建人" prop="principal" width="120" />
        <el-table-column align="left" label="备注" prop="remark" width="600" />
        <el-table-column align="left" label="日期" width="180" prop="date" sortable>
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="text" icon="edit" size="small" class="table-button" @click="updateIntermediaryContact(scope.row)">变更</el-button>
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
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="所属公司:">
          <el-input v-model="formData.company" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="联系人:">
          <el-input v-model="formData.name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="联系电话:">
          <el-input v-model="formData.telephoneNumber" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="创建人:">
          <el-input v-model="formData.principal" clearable placeholder="请输入"  :disabled="true"/>
        </el-form-item>
        <el-form-item label="备注:">
          <el-input v-model="formData.remark" clearable placeholder="请输入" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog();">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
import {
  createIntermediaryContact,
  deleteIntermediaryContact,
  deleteIntermediaryContactByIds,
  updateIntermediaryContact,
  findIntermediaryContact,
  getIntermediaryContactList,
  TestT
} from '@/api/intermediaryContact' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
import { mapGetters, mapActions } from 'vuex'
export default {
  name: 'IntermediaryContact',
  mixins: [infoList],
  data() {
    return {
      listApi: getIntermediaryContactList,
      dialogFormVisible: false,
      type: '',
      deleteVisible: false,
      multipleSelection: [],
      formData: {
        company: '',
        name: '',
        telephoneNumber: '',
        principal: '',
        remark: '',
      }
    }
  },
  async created() {
    await this.getTableData()
  },
  computed: {
    ...mapGetters('user', ['userInfo', 'sideMode', 'baseColor']),
  },
  methods: {
    ...mapActions('user', ['LoginOut', 'GetUserInfo']),
    setPrincipal(){
      this.formData.principal = this.userInfo.userName
    },
    setTestT(){
      TestT()
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
        this.deleteIntermediaryContact(row)
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
      const res = await deleteIntermediaryContactByIds({ ids })
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
    async updateIntermediaryContact(row) {
      const res = await findIntermediaryContact({ ID: row.ID })
      this.type = 'update'
      if (res.code === 0) {
        this.formData = res.data.reintermediarycontact
        this.dialogFormVisible = true
      }
    },
    closeDialog() {
      this.dialogFormVisible = false
      this.formData = {
        company: '',
        name: '',
        telephoneNumber: '',
        principal: '',
        remark: '',
      }
    },
    async deleteIntermediaryContact(row) {
      const res = await deleteIntermediaryContact({ ID: row.ID })
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
          res = await createIntermediaryContact(this.formData)
          break
        case 'update':
          res = await updateIntermediaryContact(this.formData)
          break
        default:
          res = await createIntermediaryContact(this.formData)
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
