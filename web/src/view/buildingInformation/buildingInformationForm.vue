<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="所属项目:">
          <el-input v-model="formData.project" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="楼栋名称:">
          <el-input v-model="formData.name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="楼栋状态:">
          <el-select v-model="formData.buildState" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in buildStateOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="建筑面积（㎡):">
          <el-input-number v-model="formData.coveredArea" :precision="2" clearable></el-input-number>
        </el-form-item>
        <el-form-item label="经营面积（㎡）:">
          <el-input-number v-model="formData.businessArea" :precision="2" clearable></el-input-number>
        </el-form-item>
        <el-form-item label="楼上楼层数:">
          <el-input v-model.number="formData.upstairs" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="地下楼层数:">
          <el-input v-model.number="formData.downstair" clearable placeholder="请输入" />
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
  createBuildingInformation,
  updateBuildingInformation,
  findBuildingInformation
} from '@/api/buildingInformation' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'BuildingInformation',
  mixins: [infoList],
  data() {
    return {
      type: '',
      buildStateOptions: [],
      formData: {
        project: '',
        name: '',
        buildState: undefined,
        coveredArea: 0,
        businessArea: 0,
        upstairs: 0,
        downstair: 0,
      }
    }
  },
  async created() {
    // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (this.$route.query.id) {
      const res = await findBuildingInformation({ ID: this.$route.query.id })
      if (res.code === 0) {
        this.formData = res.data.rebuildingInformation
        this.type = 'update'
      }
    } else {
      this.type = 'create'
    }
    await this.getDict('buildState')
  },
  methods: {
    async save() {
      let res
      switch (this.type) {
        case 'create':
          res = await createBuildingInformation(this.formData)
          break
        case 'update':
          res = await updateBuildingInformation(this.formData)
          break
        default:
          res = await createBuildingInformation(this.formData)
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
