<template>
  <el-space wrap>
    <el-card class="box-card" style="width: 700px;margin-top:20px">
      <template #header>
        <div class="card-header" style="margin-top:30px;margin-left:30px">
          <span style="font-size:24px;font-weight:bolder;">基础数据</span>
        </div>
      </template>
      <el-row :gutter="20">
        <el-col :span="6"><div class="grid-content"><font style="font-size:16px">项目（个）</font></div></el-col>
        <el-col :span="6"><div class="grid-content"><font style="font-size:16px">楼栋（栋）</font></div></el-col>
        <el-col :span="6"><div class="grid-content"><font style="font-size:16px">楼层（层）</font></div></el-col>
        <el-col :span="6"><div class="grid-content"><font style="font-size:16px">房源（间）</font></div></el-col>
      </el-row> 
      <el-row :gutter="20" style="margin-top:-30px">
        <el-col :span="6"><div class="grid-content"><font style="font-size:14px">{{projectNumber}}</font></div></el-col>
        <el-col :span="6"><div class="grid-content"><font style="font-size:14px">{{buildingNumber}}</font></div></el-col>
        <el-col :span="6"><div class="grid-content"><font style="font-size:14px">{{floorNumber}}</font></div></el-col>
        <el-col :span="6"><div class="grid-content"><font style="font-size:14px">{{housingNumber}}</font></div></el-col>
      </el-row> 
      <el-row :gutter="20">
        <el-col :span="6"><div class="grid-content"><font style="font-size:16px">建筑面积（㎡）</font></div></el-col>
        <el-col :span="6"><div class="grid-content"><font style="font-size:16px">经营面积（㎡）</font></div></el-col>
        <el-col :span="6"><div class="grid-content"><font style="font-size:16px">物业自持管理（个）</font></div></el-col>
        <el-col :span="6"><div class="grid-content"><font style="font-size:16px">物业外包管理（个）</font></div></el-col>
      </el-row> 
      <el-row :gutter="20" style="margin-top:-30px">
        <el-col :span="6"><div class="grid-content"><font style="font-size:14px">{{coveredArea}}</font></div></el-col>
        <el-col :span="6"><div class="grid-content"><font style="font-size:14px">{{operatingArea}}</font></div></el-col>
        <el-col :span="6"><div class="grid-content"><font style="font-size:14px">{{controlOneself}}</font></div></el-col>
        <el-col :span="6"><div class="grid-content"><font style="font-size:14px">{{outSourcing}}</font></div></el-col>
      </el-row> 
    </el-card>
    <el-card class="box-card" style="width: 700px;margin-left:100px;margin-top:20px">
      <template #header>
        <div class="card-header" style="margin-top:30px;margin-left:30px">
          <span style="font-size:24px;font-weight:bolder;">招商数据</span>
        </div>
      </template>
      <el-row :gutter="20">
        <el-col :span="6"><div class="grid-content"><font style="font-size:16px">总商机（人）</font></div></el-col>
        <el-col :span="6"><div class="grid-content"><font style="font-size:16px">本年新增商机（人）</font></div></el-col>
        <el-col :span="6"><div class="grid-content"><font style="font-size:16px">本月新增商机（人）</font></div></el-col>
        <el-col :span="6"><div class="grid-content"><font style="font-size:16px">本日新增商机（人）</font></div></el-col>
      </el-row> 
      <el-row :gutter="20" style="margin-top:-30px">
        <el-col :span="6"><div class="grid-content"><font style="font-size:14px">{{sumBusinessNumber}}</font></div></el-col>
        <el-col :span="6"><div class="grid-content"><font style="font-size:14px">{{yearBusinessNumber}}</font></div></el-col>
        <el-col :span="6"><div class="grid-content"><font style="font-size:14px">{{monthBusinessNumber}}</font></div></el-col>
        <el-col :span="6"><div class="grid-content"><font style="font-size:14px">{{dayBusinessNumber}}</font></div></el-col>
      </el-row> 
      <el-row :gutter="20">
        <el-col :span="6"><div class="grid-content"><font style="font-size:16px">总客户（人）</font></div></el-col>
        <el-col :span="6"><div class="grid-content"><font style="font-size:16px">本年新增客户（人）</font></div></el-col>
        <el-col :span="6"><div class="grid-content"><font style="font-size:16px">本月新增客户（人）</font></div></el-col>
        <el-col :span="6"><div class="grid-content"><font style="font-size:16px">本日新增客户（人）</font></div></el-col>
      </el-row> 
      <el-row :gutter="20" style="margin-top:-30px">
        <el-col :span="6"><div class="grid-content"><font style="font-size:14px">{{sumClientNumber}}</font></div></el-col>
        <el-col :span="6"><div class="grid-content"><font style="font-size:14px">{{yearClientNumber}}</font></div></el-col>
        <el-col :span="6"><div class="grid-content"><font style="font-size:14px">{{monthClientNumber}}</font></div></el-col>
        <el-col :span="6"><div class="grid-content"><font style="font-size:14px">{{dayClientNumber}}</font></div></el-col>
      </el-row> 
    </el-card>
  </el-space>
</template>
<style lang="scss">
.el-row {
  margin-bottom: -5px;
  &:last-child {
    margin-bottom: 0;
  }
}
.el-col {
  border-radius: 4px;
}
.bg-purple-dark {
  background: #99a9bf;
}
.bg-purple {
  background: #d3dce6;
}
.bg-purple-light {
  background: #e5e9f2;
}
.grid-content {
  border-radius: 4px;
  min-height: 36px;
}
.row-bg {
  padding: 10px 0;
  background-color: #f9fafc;
}
</style>


<script>
import infoList from '@/mixins/infoList'
import { getHousingMaintenanceList } from '@/api/housingMaintenance' //  此处请自行替换地址
import { getProjectInformationList } from '@/api/projectInformation'
import { getBuildingInformationList } from '@/api/buildingInformation' 
import { getFloorInformationList } from '@/api/floorInformation'
import { getBusinessList } from '@/api/business' //  此处请自行替换地址
import { getCustomerList } from '@/api/customerData' //  此处请自行替换地址
export default {
  mixins: [infoList],
  data() {
    return {
      projectNumber:'',
      buildingNumber:'',
      floorNumber:'',
      housingNumber:'',
      coveredArea:'',
      operatingArea:'',
      controlOneself:'',
      outSourcing:'',
      sumBusinessNumber:'',
      yearBusinessNumber:'',
      monthBusinessNumber:'',
      dayBusinessNumber:'',
      sumClientNumber:'',
      yearClientNumber:'',
      monthClientNumber:'',
      dayClientNumber:'',
    }
  },
  async created() {
    await this.setProject()
    await this.setBuilding()
    await this.setFloor()
    await this.setHousing()
    await this.setBusinessNumber()
    await this.setClientNumber()
  },
  methods:{
    async setClientNumber() {
      this.sumClientNumber = 0
      this.yearClientNumber = 0
      this.monthClientNumber = 0
      this.dayClientNumber = 0
      const res = await getCustomerList({ page: 1, pageSize: 999 })
      var now = new Date()
      var nowYear = now.getFullYear()
      var nowMonth = now.getMonth() + 1
      var nowDay = now.getDate()
      res.data.list && res.data.list.forEach(item => {
        if(item.audit === 0) {
          this.sumClientNumber = this.sumClientNumber + 1 
          let t = item.CreatedAt
          let today = new Date(t)
          let year = today.getFullYear()
          let month = today.getMonth() + 1
          let day = today.getDate()
          if(nowYear === year){
            this.yearClientNumber = this.yearClientNumber*1 + 1
          }
          if(nowMonth === month){
            this.monthClientNumber = this.monthClientNumber*1 + 1
          }
          if(nowDay === day){
            this.dayClientNumber = this.dayClientNumber*1 + 1
          }
        }
      })
    },
    async setBusinessNumber() {
      this.sumBusinessNumber = 0
      this.yearBusinessNumber = 0
      this.monthBusinessNumber = 0
      this.dayBusinessNumber = 0
      const res = await getBusinessList({ page: 1, pageSize: 999 })
      this.sumBusinessNumber = res.data.list.length

      var now = new Date()
      var nowYear = now.getFullYear()
      var nowMonth = now.getMonth() + 1
      var nowDay = now.getDate()
      res.data.list && res.data.list.forEach(item => {
        let t = item.CreatedAt
        let today = new Date(t)
        let year = today.getFullYear()
        let month = today.getMonth() + 1
        let day = today.getDate()
        if(nowYear === year){
          this.yearBusinessNumber = this.yearBusinessNumber + 1
        }
        if(nowMonth === month){
          this.monthBusinessNumber = this.monthBusinessNumber + 1
        }
        if(nowDay === day){
          this.dayBusinessNumber = this.dayBusinessNumber + 1
        }
      })
    },
    async setProject() {
      this.projectNumber = ''
      this.coveredArea = ''
      this.operatingArea = ''
      this.controlOneself = ''
      this.outSourcing = ''
      const res = await getProjectInformationList({ page: 1, pageSize: 999 })
      this.projectNumber = res.data.list.length
      res.data.list && res.data.list.forEach(item => {
        this.coveredArea = this.coveredArea * 1 + item.coveredArea
        this.operatingArea = this.operatingArea * 1 + item.operatingArea
        if(item.managementType === 0){
          this.controlOneself = this.controlOneself + 1
        }
        else {
          this.outSourcing = this.outSourcing + 1
        }
      });
    },
    async setBuilding() {
      this.buildingNumber = ''
      const res = await getBuildingInformationList({ page: 1, pageSize: 999 })
      this.buildingNumber = res.data.list.length
    },
    async setFloor() {
      this.floorNumber = ''
      const res = await getFloorInformationList({ page: 1, pageSize: 999 })
      this.floorNumber = res.data.list.length
    },
    async setHousing() {
      this.housingNumber = ''
      const res = await getHousingMaintenanceList({ page: 1, pageSize: 999 })
      this.housingNumber = res.data.list.length
    },
  },
};
</script>