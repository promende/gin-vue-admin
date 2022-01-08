<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="所属项目:">
          <el-input v-model="formData.project" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="所属楼栋:">
          <el-input v-model="formData.build" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="所属楼层:">
          <el-input v-model="formData.floor" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="房源名称:">
          <el-input v-model="formData.name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="房源状态:">
          <el-select v-model="formData.state" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in buildStateOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="房源类型:">
          <el-select v-model="formData.type" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in HousingTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="计租面积:">
          <el-input-number v-model="formData.meterRentArea" :precision="2" clearable></el-input-number>
        </el-form-item>
        <el-form-item label="实用面积:">
          <el-input-number v-model="formData.usableArea" :precision="2" clearable></el-input-number>
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
  createHousingMaintenance,
  updateHousingMaintenance,
  findHousingMaintenance
} from '@/api/housingMaintenance' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'HousingMaintenance',
  mixins: [infoList],
  data() {
    return {
      type: '',
      buildStateOptions: [],
      HousingTypeOptions: [],
      formData: {
        project: '',
        build: '',
        floor: '',
        name: '',
        state: undefined,
        type: undefined,
        meterRentArea: 0,
        usableArea: 0,
      }
    }
  },
  async created() {
    // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (this.$route.query.id) {
      const res = await findHousingMaintenance({ ID: this.$route.query.id })
      if (res.code === 0) {
        this.formData = res.data.rehousingMaintenance
        this.type = 'update'
      }
    } else {
      this.type = 'create'
    }
    await this.getDict('buildState')
    await this.getDict('HousingType')
  },
  methods: {
    async save() {
      let res
      switch (this.type) {
        case 'create':
          res = await createHousingMaintenance(this.formData)
          break
        case 'update':
          res = await updateHousingMaintenance(this.formData)
          break
        default:
          res = await createHousingMaintenance(this.formData)
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
