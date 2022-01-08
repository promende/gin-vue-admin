<template>
  <div>
    <div class="gva-form-box">
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
          <el-input v-model="formData.principal" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="备注:">
          <el-input v-model="formData.remark" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" @click="save">保存</el-button>
          <el-button size="mini" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
import {
  createIntermediaryContact,
  updateIntermediaryContact,
  findIntermediaryContact
} from '@/api/intermediaryContact' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'IntermediaryContact',
  mixins: [infoList],
  data() {
    return {
      type: '',
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
    // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (this.$route.query.id) {
      const res = await findIntermediaryContact({ ID: this.$route.query.id })
      if (res.code === 0) {
        this.formData = res.data.reintermediarycontact
        this.type = 'update'
      }
    } else {
      this.type = 'create'
    }
  },
  methods: {
    async save() {
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
      }
    },
    back() {
      this.$router.go(-1)
    }
  }
}
</script>

<style>
</style>
