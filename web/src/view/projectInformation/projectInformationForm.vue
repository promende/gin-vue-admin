<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="项目名称:">
          <el-input v-model="formData.name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="项目简称:">
          <el-input v-model="formData.abbreviation" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="项目地址:">
          <el-input v-model="formData.address" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="营运状态:">
          <el-select v-model="formData.operatingState" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in OperatingStateOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="经营类型:">
          <el-select v-model="formData.managementType" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in managementTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="物业管理类型:">
          <el-select v-model="formData.propertyManagementType" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in PropertyManagementTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="建筑面积（㎡):">
          <el-input-number v-model="formData.coveredArea" :precision="2" clearable></el-input-number>
        </el-form-item>
        <el-form-item label="经营面积（㎡）:">
          <el-input-number v-model="formData.operatingArea" :precision="2" clearable></el-input-number>
        </el-form-item>
        <el-form-item label="负责人:">
          <el-input v-model="formData.principal" clearable placeholder="请输入" />
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
  createProjectInformation,
  updateProjectInformation,
  findProjectInformation
} from '@/api/projectInformation' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'ProjectInformation',
  mixins: [infoList],
  data() {
    return {
      type: '',
      OperatingStateOptions: [],
      managementTypeOptions: [],
      PropertyManagementTypeOptions: [],
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
      }
    }
  },
  async created() {
    // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (this.$route.query.id) {
      const res = await findProjectInformation({ ID: this.$route.query.id })
      if (res.code === 0) {
        this.formData = res.data.reproject
        this.type = 'update'
      }
    } else {
      this.type = 'create'
    }
    await this.getDict('OperatingState')
    await this.getDict('managementType')
    await this.getDict('PropertyManagementType')
  },
  methods: {
    async save() {
      let res
      switch (this.type) {
        case 'create':
          res = await createProjectInformation(this.formData)
          break
        case 'update':
          res = await updateProjectInformation(this.formData)
          break
        default:
          res = await createProjectInformation(this.formData)
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
