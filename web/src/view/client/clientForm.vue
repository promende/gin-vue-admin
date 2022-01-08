<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="商家名称:">
          <el-input v-model="formData.name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商家性质:">
          <el-select v-model="formData.type" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in intermediaryTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="联系人:">
          <el-input v-model="formData.linkman" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="身份证号:">
          <el-input v-model="formData.iDNumber" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="地址:">
          <el-input v-model="formData.address" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="联系电话:">
          <el-input v-model="formData.telephone" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="开票名称:">
          <el-input v-model="formData.invoice" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="开户银行:">
          <el-input v-model="formData.bank" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="开户账号:">
          <el-input v-model="formData.account" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="审核状态:">
          <el-select v-model="formData.audit" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in auditTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
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
  createClient,
  updateClient,
  findClient
} from '@/api/client' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'Client',
  mixins: [infoList],
  data() {
    return {
      type: '',
      intermediaryTypeOptions: [],
      auditTypeOptions: [],
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
        audit: undefined,
      }
    }
  },
  async created() {
    // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (this.$route.query.id) {
      const res = await findClient({ ID: this.$route.query.id })
      if (res.code === 0) {
        this.formData = res.data.reclient
        this.type = 'update'
      }
    } else {
      this.type = 'create'
    }
    await this.getDict('intermediaryType')
    await this.getDict('auditType')
  },
  methods: {
    async save() {
      let res
      switch (this.type) {
        case 'create':
          res = await createClient(this.formData)
          break
        case 'update':
          res = await updateClient(this.formData)
          break
        default:
          res = await createClient(this.formData)
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
