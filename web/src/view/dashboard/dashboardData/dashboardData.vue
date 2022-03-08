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
    <el-card class="box-card" style="width: 700px;margin-top:20px">
      <template #header>
        <div class="card-header" style="margin-top:30px;margin-left:30px">
          <span style="font-size:24px;font-weight:bolder;">运营数据</span>
        </div>
      </template>
      <el-row :gutter="20">
        <el-col :span="6"><div class="grid-content"><font style="font-size:16px">商机同比增长（月）</font></div></el-col>
        <el-col :span="6"><div class="grid-content"><font style="font-size:16px">商机同比增长（季）</font></div></el-col>
        <el-col :span="6"><div class="grid-content"><font style="font-size:16px">商机环比增长（月）</font></div></el-col>
        <el-col :span="6"><div class="grid-content"><font style="font-size:16px">商机环比增长（季）</font></div></el-col>
      </el-row> 
      <el-row :gutter="20" style="margin-top:-30px">
        <el-col :span="6"><div class="grid-content"><font style="font-size:14px">{{bNumber1}}</font></div></el-col>
        <el-col :span="6"><div class="grid-content"><font style="font-size:14px">{{bNumber2}}</font></div></el-col>
        <el-col :span="6"><div class="grid-content"><font style="font-size:14px">{{bNumber3}}</font></div></el-col>
        <el-col :span="6"><div class="grid-content"><font style="font-size:14px">{{bNumber4}}</font></div></el-col>
      </el-row> 
      <el-row :gutter="20">
        <el-col :span="6"><div class="grid-content"><font style="font-size:16px">客户同比增长（月）</font></div></el-col>
        <el-col :span="6"><div class="grid-content"><font style="font-size:16px">客户同比增长（季）</font></div></el-col>
        <el-col :span="6"><div class="grid-content"><font style="font-size:16px">客户环比增长（月）</font></div></el-col>
        <el-col :span="6"><div class="grid-content"><font style="font-size:16px">客户环比增长（季）</font></div></el-col>
      </el-row> 
      <el-row :gutter="20" style="margin-top:-30px">
        <el-col :span="6"><div class="grid-content"><font style="font-size:14px">{{cNumber1}}</font></div></el-col>
        <el-col :span="6"><div class="grid-content"><font style="font-size:14px">{{cNumber2}}</font></div></el-col>
        <el-col :span="6"><div class="grid-content"><font style="font-size:14px">{{cNumber3}}</font></div></el-col>
        <el-col :span="6"><div class="grid-content"><font style="font-size:14px">{{cNumber4}}</font></div></el-col>
      </el-row> 
    </el-card>
    <el-card class="box-card" style="width: 700px;margin-left:100px;margin-top:20px">
      <template #header>
        <div class="card-header" style="margin-top:30px;margin-left:30px">
          <span style="font-size:24px;font-weight:bolder;">招商数据</span>
        </div>
      </template>
      <el-row :gutter="20">
        <el-col :span="6"><div class="grid-content"><font style="font-size:16px">本年商机王</font></div></el-col>
        <el-col :span="6"><div class="grid-content"><font style="font-size:16px">本季商机王</font></div></el-col>
        <el-col :span="6"><div class="grid-content"><font style="font-size:16px">本月商机王</font></div></el-col>
        <el-col :span="6"><div class="grid-content"><font style="font-size:16px">本日商机王</font></div></el-col>
      </el-row> 
      <el-row :gutter="20" style="margin-top:-30px">
        <el-col :span="6"><div class="grid-content"><font style="font-size:14px">{{bNumber5}}</font></div></el-col>
        <el-col :span="6"><div class="grid-content"><font style="font-size:14px">{{bNumber6}}</font></div></el-col>
        <el-col :span="6"><div class="grid-content"><font style="font-size:14px">{{bNumber7}}</font></div></el-col>
        <el-col :span="6"><div class="grid-content"><font style="font-size:14px">{{bNumber8}}</font></div></el-col>
      </el-row> 
      <el-row :gutter="20">
        <el-col :span="6"><div class="grid-content"><font style="font-size:16px">本年客户王</font></div></el-col>
        <el-col :span="6"><div class="grid-content"><font style="font-size:16px">本季客户王</font></div></el-col>
        <el-col :span="6"><div class="grid-content"><font style="font-size:16px">本月客户王</font></div></el-col>
        <el-col :span="6"><div class="grid-content"><font style="font-size:16px">本日客户王</font></div></el-col>
      </el-row> 
      <el-row :gutter="20" style="margin-top:-30px">
        <el-col :span="6"><div class="grid-content"><font style="font-size:14px">{{cNumber5}}</font></div></el-col>
        <el-col :span="6"><div class="grid-content"><font style="font-size:14px">{{cNumber6}}</font></div></el-col>
        <el-col :span="6"><div class="grid-content"><font style="font-size:14px">{{cNumber7}}</font></div></el-col>
        <el-col :span="6"><div class="grid-content"><font style="font-size:14px">{{cNumber8}}</font></div></el-col>
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
      bNumber1:'',
      bNumber2:'',
      bNumber3:'',
      bNumber4:'',
      bNumber5:'',
      bNumber6:'',
      bNumber7:'',
      bNumber8:'',
      cNumber1:'',
      cNumber2:'',
      cNumber3:'',
      cNumber4:'',
      cNumber5:'',
      cNumber6:'',
      cNumber7:'',
      cNumber8:'',
    }
  },
  async created() {
    await this.setProject()
    await this.setBuilding()
    await this.setFloor()
    await this.setHousing()
    await this.setBusinessNumber()
    await this.setClientNumber()
    await this.setbNumber()
    await this.setcNumber()
  },
  methods:{
    async setbNumber(){
      var now = new Date()
      var nowYear = now.getFullYear()
      var nowMonth = now.getMonth() + 1
      var nowDay = now.getDate()
      const res = await getBusinessList({ page: 1, pageSize: 999 })
      let a = 0;
      let b = 0
      let c = 0;
      let d = 0
      let e = 0;
      let f = 0
      let g = 0;
      let h = 0
      let dic1 = new Array();
      let dic2 = new Array();
      let dic3 = new Array();
      let dic4 = new Array();
      res.data.list && res.data.list.forEach(item => {
        let t = item.CreatedAt
        let today = new Date(t)
        let year = today.getFullYear()
        let month = today.getMonth() + 1
        let day = today.getDate()
        
        if(nowYear === year && nowMonth == month){
          a++
        }
        if(nowYear - 1 === year && nowMonth == month){
          b++
        }
        if(nowMonth === 1){
          if(nowYear === year && nowMonth === month){
            e++
          }
          if(nowYear-1 === year && month === 12){
            f++
          }
        }
        else{
          if(nowYear === year && nowMonth === month){
            e++
          }
          if(nowYear === year && nowMonth - 1 === month){
            f++
          }
        }

        if(nowMonth <= 3){
          if(nowYear === year && month <= 3){
            c++
          }
          if(nowYear - 1 === year && month <= 3){
            d++
          }
        }
        else if(nowMonth <= 6){
          if(nowYear === year && month > 3 && month <= 6){
            c++
          }
          if(nowYear - 1 === year && month > 3 && month <= 6){
            d++
          }
        }
        else if(nowMonth <= 9){
          if(nowYear === year && month > 6 && month <= 9){
            c++
          }
          if(nowYear - 1 === year && month > 6 && month <= 9){
            d++
          }
        }
        else if(nowMonth <= 12){
          if(nowYear === year && month > 9 && month <= 12){
            c++
          }
          if(nowYear - 1 === year && month > 9 && month <= 12){
            d++
          }
        }
        
        if(nowMonth <= 3){
          if(nowYear === year && month <= 3){
            g++
          }
          if(nowYear - 1 === year && month > 9 && month <= 12){
            h++
          }
        }
        else if(nowMonth <= 6){
          if(nowYear === year && month > 3 && month <= 6){
            g++
          }
          if(nowYear - 1 === year && month > 0 && month <= 3){
            h++
          }
        }
        else if(nowMonth <= 9){
          if(nowYear === year && month > 6 && month <= 9){
            g++
          }
          if(nowYear - 1 === year && month > 3 && month <= 6){
            h++
          }
        }
        else if(nowMonth <= 12){
          if(nowYear === year && month > 9 && month <= 12){
            g++
          }
          if(nowYear - 1 === year && month > 6 && month <= 9){
            h++
          }
        }
        
        if(nowYear === year){
          dic1[item.principal] ++;
        }

        if(nowMonth <= 3){
          if(nowYear === year && month <= 3){
            dic2[item.principal] ++;
          }
        }
        else if(nowMonth <= 6){
          if(nowYear === year && month > 3 && month <= 6){
            dic2[item.principal] ++;
          }
        }
        else if(nowMonth <= 9){
          if(nowYear === year && month > 6 && month <= 9){
            dic2[item.principal] ++;
          }
        }
        else if(nowMonth <= 12){
          if(nowYear === year && month > 9 && month <= 12){
            dic2[item.principal] ++;
          }
        }

        if(nowMonth === month){
          dic3[item.principal] ++;
        }
        if(nowDay === day){
          dic4[item.principal] ++;
        }
      })
      this.bNumber1 = ((a - b) / b * 10000 / 100).toFixed(2) + '%'
      this.bNumber2 = ((c - d) / d * 10000 / 100).toFixed(2) + '%'
      this.bNumber3 = ((e - f) / f * 10000 / 100).toFixed(2) + '%'
      this.bNumber4 = ((g - h) / h * 10000 / 100).toFixed(2) + '%'
      let res1 =Object.keys(dic1).sort(function(a,b){ return dic1[a]-dic1[b]; });
      this.bNumber5 = res1[0]
      res1 =Object.keys(dic2).sort(function(a,b){ return dic2[a]-dic2[b]; });
      this.bNumber6 = res1[0]
      res1 =Object.keys(dic3).sort(function(a,b){ return dic3[a]-dic3[b]; });
      this.bNumber7 = res1[0]
      res1 =Object.keys(dic4).sort(function(a,b){ return dic4[a]-dic4[b]; });
      this.bNumber8 = res1[0]
      
    },
    async setcNumber(){
      var now = new Date()
      var nowYear = now.getFullYear()
      var nowMonth = now.getMonth() + 1
      var nowDay = now.getDate()
      const res = await getCustomerList({ page: 1, pageSize: 999 })
      let a = 0;
      let b = 0
      let c = 0;
      let d = 0
      let e = 0;
      let f = 0
      let g = 0;
      let h = 0
      let dic1 = new Array();
      let dic2 = new Array();
      let dic3 = new Array();
      let dic4 = new Array();
      res.data.list && res.data.list.forEach(item => {
        if(item.audit === 0) {
          let t = item.CreatedAt
          let today = new Date(t)
          let year = today.getFullYear()
          let month = today.getMonth() + 1
          let day = today.getDate()
          
          if(nowYear === year && nowMonth == month){
            a++
          }
          if(nowYear - 1 === year && nowMonth == month){
            b++
          }
          if(nowMonth === 1){
            if(nowYear === year && nowMonth === month){
              e++
            }
            if(nowYear-1 === year && month === 12){
              f++
            }
          }
          else{
            if(nowYear === year && nowMonth === month){
              e++
            }
            if(nowYear === year && nowMonth-1 === month){
              f++
            }
          }

          if(nowMonth <= 3){
            if(nowYear === year && month <= 3){
              c++
            }
            if(nowYear - 1 === year && month <= 3){
              d++
            }
          }
          else if(nowMonth <= 6){
            if(nowYear === year && month > 3 && month <= 6){
              c++
            }
            if(nowYear - 1 === year && month > 3 && month <= 6){
              d++
            }
          }
          else if(nowMonth <= 9){
            if(nowYear === year && month > 6 && month <= 9){
              c++
            }
            if(nowYear - 1 === year && month > 6 && month <= 9){
              d++
            }
          }
          else if(nowMonth <= 12){
            if(nowYear === year && month > 9 && month <= 12){
              c++
            }
            if(nowYear - 1 === year && month > 9 && month <= 12){
              d++
            }
          }
          
          if(nowMonth <= 3){
            if(nowYear === year && month <= 3){
              g++
            }
            if(nowYear - 1 === year && month > 9 && month <= 12){
              h++
            }
          }
          else if(nowMonth <= 6){
            if(nowYear === year && month > 3 && month <= 6){
              g++
            }
            if(nowYear - 1 === year && month > 0 && month <= 3){
              h++
            }
          }
          else if(nowMonth <= 9){
            if(nowYear === year && month > 6 && month <= 9){
              g++
            }
            if(nowYear - 1 === year && month > 3 && month <= 6){
              h++
            }
          }
          else if(nowMonth <= 12){
            if(nowYear === year && month > 9 && month <= 12){
              g++
            }
            if(nowYear - 1 === year && month > 6 && month <= 9){
              h++
            }
          }
          if(nowYear === year){
            dic1[item.principal] ++;
          }

          if(nowMonth <= 3){
            if(nowYear === year && month <= 3){
              dic2[item.principal] ++;
            }
          }
          else if(nowMonth <= 6){
            if(nowYear === year && month > 3 && month <= 6){
              dic2[item.principal] ++;
            }
          }
          else if(nowMonth <= 9){
            if(nowYear === year && month > 6 && month <= 9){
              dic2[item.principal] ++;
            }
          }
          else if(nowMonth <= 12){
            if(nowYear === year && month > 9 && month <= 12){
              dic2[item.principal] ++;
            }
          }

          if(nowMonth === month){
            dic3[item.principal] ++;
          }
          if(nowDay === day){
            dic4[item.principal] ++;
          }
        }
      })
      this.cNumber1 = ((a - b) / b * 10000 / 100).toFixed(2) + '%'
      this.cNumber2 = ((c - d) / d * 10000 / 100).toFixed(2) + '%'
      this.cNumber3 = ((e - f) / f * 10000 / 100).toFixed(2) + '%'
      this.cNumber4 = ((g - h) / h * 10000 / 100).toFixed(2) + '%'
      let res1 =Object.keys(dic1).sort(function(a,b){ return dic1[a]-dic1[b]; });
      this.cNumber5 = res1[0]
      res1 =Object.keys(dic2).sort(function(a,b){ return dic2[a]-dic2[b]; });
      this.cNumber6 = res1[0]
      res1 =Object.keys(dic3).sort(function(a,b){ return dic3[a]-dic3[b]; });
      this.cNumber7 = res1[0]
      res1 =Object.keys(dic4).sort(function(a,b){ return dic4[a]-dic4[b]; });
      this.cNumber8 = res1[0]
    },
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