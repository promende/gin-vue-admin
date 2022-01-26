<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="合同编号">
          <el-select v-model="searchInfo.contractNumber" placeholder="请选择" style="width:100%" default-first-option clearable filterable @visible-change="setContractNumberOptions">
            <el-option v-for="(item,key) in contractNumberOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="商家名称">
          <el-select v-model="searchInfo.merchant" placeholder="请选择" style="width:100%" default-first-option clearable filterable @visible-change="setMerchantOptions">
            <el-option v-for="(item,key) in merchantOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="项目名称">
          <el-select v-model="searchInfo.project" placeholder="请选择" style="width:100%" default-first-option clearable filterable @visible-change="setProjectOptions">
            <el-option v-for="(item,key) in projectOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="审核状态">
          <el-select v-model="searchInfo.auditType" placeholder="请选择" style="width:100%" default-first-option clearable filterable >
            <el-option v-for="(item,key) in auditTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="负责人">
          <el-select v-model="searchInfo.principal" placeholder="请选择" style="width:100%" default-first-option clearable filterable @visible-change="setPrincipalOptions">
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
            <el-button size="mini" type="primary" icon="plus" @click="openDialog();setAuditType()">新增</el-button>
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
        @row-click="handleRowClick"
        >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" type="expand" >
          <template #default="scope">
            <el-descriptions :column="3" border>
              <el-descriptions-item label="合同编号" label-align="center" align="left" width="10px" label-class-name="background: #E1F3D8;">
                {{ scope.row.contractNumber }}
              </el-descriptions-item>
              <el-descriptions-item label="项目名称" label-align="center" align="left" width="10px">
                {{ scope.row.project }}
              </el-descriptions-item>
              <el-descriptions-item label="楼栋名称" label-align="center" align="left" width="10px">
                {{ scope.row.building }}
              </el-descriptions-item>
              <el-descriptions-item label="楼层名称" label-align="center" align="left" width="80px">
                {{ scope.row.floor }}
              </el-descriptions-item>
              <el-descriptions-item label="房源名称" label-align="center" align="left" width="80px">
                {{ scope.row.housing }}
              </el-descriptions-item>
              <el-descriptions-item label="商家名称" label-align="center" align="left" width="80px">
                {{ scope.row.merchant }}
              </el-descriptions-item>
              <el-descriptions-item label="合同类型" label-align="center" align="left" width="80px">
                {{ filterDict(scope.row.contractType,"contractType") }}
              </el-descriptions-item>
              <el-descriptions-item label="合同签署" label-align="center" align="left" width="80px">
                {{ filterDict(scope.row.contractSigning,"contractSigning") }}
              </el-descriptions-item>
              <el-descriptions-item label="是否续签" label-align="center" align="left" width="80px">
                {{ filterDict(scope.row.renew,"renew") }}
              </el-descriptions-item>
              <el-descriptions-item label="关联合同编号" label-align="center" align="left" width="80px">
                {{ scope.row.associatedContractNumber }}
              </el-descriptions-item>
              <el-descriptions-item label="是否中介介入" label-align="center" align="left" width="80px">
                {{ filterDict(scope.row.intermediary,"Intermediary") }}
              </el-descriptions-item>
              <el-descriptions-item label="中介公司" label-align="center" align="left" width="10px" label-class-name="background: #E1F3D8;">
                {{ scope.row.agency }}
              </el-descriptions-item>
              <el-descriptions-item label="中介联系人" label-align="center" align="left" width="10px">
                {{ scope.row.intermediaryContact }}
              </el-descriptions-item>
              <el-descriptions-item label="负责人" label-align="center" align="left" width="10px">
                {{ scope.row.principal }}
              </el-descriptions-item>
              <el-descriptions-item label="交付日" label-align="center" align="left" width="80px">
                {{ formatDate(scope.row.deliveryDate) }}
              </el-descriptions-item>
              <el-descriptions-item label="合同开始时间" label-align="center" align="left" width="80px">
                {{ formatDate(scope.row.startTime) }}
              </el-descriptions-item>
              <el-descriptions-item label="合同结束时间" label-align="center" align="left" width="80px">
                {{ formatDate(scope.row.endTime) }}
              </el-descriptions-item>
              <el-descriptions-item label="支付周期" label-align="center" align="left" width="80px">
                {{ filterDict(scope.row.paymentCycle,"paymentCycle") }}
              </el-descriptions-item>
              <el-descriptions-item label="单价" label-align="center" align="left" width="10px" label-class-name="background: #E1F3D8;">
                {{ filterDict(scope.row.univalence,"univalence") }}
              </el-descriptions-item>
              <el-descriptions-item label="租金" label-align="center" align="left" width="10px">
                {{ scope.row.rent }}
              </el-descriptions-item>
              <el-descriptions-item label="服务费" label-align="center" align="left" width="10px">
                {{ scope.row.serviceCharge }}
              </el-descriptions-item>
              <el-descriptions-item label="物管费" label-align="center" align="left" width="80px">
                {{ scope.row.propertyManagementFee }}
              </el-descriptions-item>
              <el-descriptions-item label="合同总金额" label-align="center" align="left" width="80px">
                {{ scope.row.contractGrandTotal }}
              </el-descriptions-item>
              <el-descriptions-item label="设置费" label-align="center" align="left" width="80px">
                {{ scope.row.setUpFee }}
              </el-descriptions-item>
              <el-descriptions-item label="保证金" label-align="center" align="left" width="10px">
                {{ scope.row.earnestMoney }}
              </el-descriptions-item>
              <el-descriptions-item label="备注" label-align="center" align="left" width="80px">
                {{ scope.row.remark }}
              </el-descriptions-item>
              <el-descriptions-item label="创建日期" label-align="center" align="left" width="80px">
                {{ formatDate(scope.row.CreatedAt) }}
              </el-descriptions-item>
              <el-descriptions-item label="审核状态" label-align="center" align="left" width="80px">
                <el-tag v-if="scope.row.auditType===0" class="ml-2" type="success">{{ filterDict(scope.row.auditType,"auditType") }}</el-tag>
                <el-tag v-else class="ml-2" type="danger">{{ filterDict(scope.row.auditType,"auditType") }}</el-tag>
              </el-descriptions-item>
            </el-descriptions>
          </template>
        </el-table-column>
        <el-table-column align="left" label="合同编号" prop="contractNumber" width="120" />
        <el-table-column align="left" label="商家名称" prop="merchant" width="120" />
        <el-table-column align="left" label="项目名称" prop="project" width="120" />
        <el-table-column align="left" label="审核状态" prop="auditType" width="120">
            <template #default="scope">
              <el-tag v-if="scope.row.auditType===0" class="ml-2" type="success">{{ filterDict(scope.row.auditType,"auditType") }}</el-tag>
              <el-tag v-else class="ml-2" type="danger">{{ filterDict(scope.row.auditType,"auditType") }}</el-tag>
            </template>
        </el-table-column>
        <el-table-column align="left" label="负责人" prop="principal" width="120" />
        <el-table-column align="left" label="创建日期" width="180" prop="date" sortable>
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="text" icon="edit" size="small" class="table-button" @click="updatePact(scope.row);handleRowClick(scope.row)">变更</el-button>
            <el-button type="text" icon="delete" size="mini" @click="deletePact(scope.row);handleRowClick(scope.row)">删除</el-button>
            <el-button v-if="scope.row.auditType===0" type="text" icon="tools" size="small" @click="handleRowClick(scope.row);changeAuditType1(scope.row);">取消审核</el-button>
            <el-button v-else type="text" icon="tools" size="small" @click="handleRowClick(scope.row);changeAuditType(scope.row);">审核</el-button> 
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
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="新增合同">
      <el-form :model="formData" label-position="right" label-width="120px" ref="formData" :rules="rules">
        <el-form-item label="租赁范围" prop="a">
          <el-cascader
            v-model="this.housingSelects"
            style="width:100%"
            :options="housingOptions"
            :props="{ multiple:true, label:'label', value:'value', disabled:'disabled', checkStrictly:false}"
            @visible-change="setHousingOptions"
          />
        </el-form-item>
        <el-row style="margin-top:-10px">
            <el-col :span="12">
              <el-form-item label="商家名称" prop="b">
                <el-select v-model="formData.merchant" placeholder="请选择" style="width:100%" default-first-option clearable filterable @visible-change="setCustomerOptions" @change="setPrincipal">
                  <el-option v-for="(item,key) in customerOptions" :key="key" :label="item.label" :value="item.value" />
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="负责人">
                <el-input v-model="formData.principal" clearable placeholder="" disabled/>
              </el-form-item>
            </el-col>
        </el-row>
        <el-row style="margin-top:-20px">
            <el-col :span="12">
              <el-form-item label="合同类型" prop="c">
                <el-select v-model="formData.contractType" placeholder="请选择" style="width:100%" clearable>
                  <el-option v-for="(item,key) in contractTypeOptions" :key="key" :label="item.label" :value="item.value" />
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="合同签署" prop="d">
                <el-select v-model="formData.contractSigning" placeholder="请选择" style="width:100%" clearable>
                  <el-option v-for="(item,key) in contractSigningOptions" :key="key" :label="item.label" :value="item.value" />
                </el-select>
              </el-form-item>
            </el-col>
        </el-row>
        <el-row style="margin-top:-20px">
            <el-col :span="12">
              <el-form-item label="是否续签" prop="e">
                <el-select v-model="formData.renew" placeholder="请选择" style="width:30%" clearable>
                  <el-option v-for="(item,key) in renewOptions" :key="key" :label="item.label" :value="item.value" />
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="关联合同编号" v-if="this.formData.renew === 0" prop="f">
                <el-input v-model="formData.associatedContractNumber" clearable placeholder="请输入" />
              </el-form-item>
            </el-col>
        </el-row>
        <el-row style="margin-top:-20px">
            <el-col :span="12">
              <el-form-item label="是否中介介入" prop="g">
                <el-select v-model="formData.intermediary" placeholder="请选择" style="width:30%" clearable>
                  <el-option v-for="(item,key) in IntermediaryOptions" :key="key" :label="item.label" :value="item.value" />
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="中介人" v-if="this.formData.intermediary===0" prop="h">
                <el-cascader
                  v-model="this.middlemanSelects"
                  style="width:100%"
                  :options="middlemanOptions"
                  :props="{ multiple:false, label:'label', value:'value', disabled:'disabled'}"
                  @visible-change="setMiddlemanOptions"
                />
              </el-form-item>
            </el-col>
        </el-row>
        <el-row style="margin-top:-20px">
            <el-col :span="12">
              <el-form-item label="交付日" style="margin-top:-10px" prop="date1">
                <el-date-picker v-model="formData.deliveryDate" type="date" style="width:100%" placeholder="选择日期" clearable />
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="合同起止时间" style="margin-top:-10px" prop="j">
                <el-date-picker
                  style="width:100%"
                  v-model="beginToEnd"
                  type="daterange"
                  unlink-panels
                  range-separator="-"
                  start-placeholder="合同开始时间"
                  end-placeholder="合同结束时间"
                  :shortcuts="shortcuts"
                >
                </el-date-picker>
              </el-form-item>
            </el-col>
        </el-row>
        <el-row style="margin-top:-20px">
            <el-col :span="12">
              <el-form-item label="支付周期" prop="k">
                <el-select v-model="formData.paymentCycle" placeholder="请选择" style="width:100%" clearable>
                  <el-option v-for="(item,key) in paymentCycleOptions" :key="key" :label="item.label" :value="item.value" />
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="单价" prop="h">
                <el-select v-model="formData.univalence" placeholder="请选择" style="width:100%" clearable>
                  <el-option v-for="(item,key) in univalenceOptions" :key="key" :label="item.label" :value="item.value" />
                </el-select>
              </el-form-item>
            </el-col>
        </el-row>
        <el-form-item label="租金" prop="a">
          <el-input-number v-model="formData.rent"  style="width:100%" :precision="2" clearable />
        </el-form-item>
        <el-form-item label="服务费" prop="a">
          <el-input-number v-model="formData.serviceCharge"  style="width:100%" :precision="2" clearable />
        </el-form-item>
        <el-form-item label="物管费" prop="a">
          <el-input-number v-model="formData.propertyManagementFee"  style="width:100%" :precision="2" clearable />
        </el-form-item>
        <el-form-item label="合同总金额" prop="a">
          <el-input-number v-model="formData.contractGrandTotal"  style="width:100%" :precision="2" clearable />
        </el-form-item>
        <el-form-item label="设置费" prop="a">
          <el-input-number v-model="formData.setUpFee"  style="width:100%" :precision="2" clearable />
        </el-form-item>
        <el-form-item label="保证金" prop="a">
          <el-input-number v-model="formData.earnestMoney"  style="width:100%" :precision="2" clearable />
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="formData.remark" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="审核状态">
          <el-select v-model="formData.auditType" placeholder="请选择" style="width:100%" clearable disabled>
            <el-option v-for="(item,key) in auditTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
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
  createPact,
  deletePact,
  deletePactByIds,
  updatePact,
  findPact,
  getPactList
} from '@/api/pact' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
import { getUserList } from '@/api/user'
import { getProjectInformationList } from '@/api/projectInformation' 
import { getHousingMaintenanceList } from '@/api/housingMaintenance' //  此处请自行替换地址
import { getBuildingInformationList } from '@/api/buildingInformation' 
import { getFloorInformationList } from '@/api/floorInformation'
import { getCustomerList } from '@/api/customerData' //  此处请自行替换地址
import { getMiddlemanList } from '@/api/middleman' //  此处请自行替换地址
import {getIntermediaryCompanyList } from '@/api/intermediaryCompany' //  此处请自行替换地址
export default {
  name: 'Pact',
  mixins: [infoList],
  data() {
    return {
      listApi: getPactList,
      dialogFormVisible: false,
      type: '',
      deleteVisible: false,
      multipleSelection: [],
      paymentCycleOptions: [],
      univalenceOptions: [],
      auditTypeOptions: [],
      contractTypeOptions: [],
      contractSigningOptions: [],
      renewOptions: [],
      IntermediaryOptions: [],
      principalOptions: [],
      projectOptions: [],
      contractNumberOptions: [],
      merchantOptions: [],
      housingOptions: [],
      housingSelects: [],
      middlemanSelects: [],
      middlemanOptions: [],
      customerOptions: [],
      beginToEnd: [],
      formData: {
        contractNumber: '',
        project: '',
        building: '',
        floor: '',
        housing: '',
        merchant: '',
        contractType: undefined,
        contractSigning: undefined,
        renew: undefined,
        associatedContractNumber: '',
        intermediary: undefined,
        agency: '',
        intermediaryContact: '',
        principal: '',
        deliveryDate: '',
        startTime: '',
        endTime: '',
        paymentCycle: undefined,
        univalence: undefined,
        rent: 0,
        serviceCharge: 0,
        propertyManagementFee: 0,
        contractGrandTotal: 0,
        setUpFee: 0,
        earnestMoney: 0,
        remark: '',
        auditType: undefined,
      },
      rules: {
        a:                [{ required: true, message: '请输入租赁范围',  trigger: 'blur' }],
        b:                [{ required: true, message: '请输入商家名称',  trigger: 'blur' }],
        c:                [{ required: true, message: '请输入合同类型',  trigger: 'blur' }],
        d:                [{ required: true, message: '请输入合同签署',  trigger: 'blur' }],
        e:                [{ required: true, message: '请输入是否续签',  trigger: 'blur' }],
        f:                [{ required: true, message: '请输入关联合同编号',  trigger: 'blur' }],
        g:                [{ required: true, message: '请输入是否中介介入',  trigger: 'blur' }],
        h:                [{ required: true, message: '请输入中介人',  trigger: 'blur' }],date1: [
    {
      type: 'date',
      required: true,
      message: 'Please pick a date',
      trigger: 'change',
    },
  ],
        j:                [{ type: 'date', required: true, message: '请输入合同起止时间',  trigger: 'change' }],
        k:                [{ required: true, message: '请输入支付周期',  trigger: 'blur' }],
        l:                [{ required: true, message: '请输入单价',  trigger: 'blur' }],
      }
    }
  },
  async created() {
    await this.getTableData()
    await this.getDict('paymentCycle')
    await this.getDict('univalence')
    await this.getDict('auditType')
    await this.getDict('contractType')
    await this.getDict('contractSigning')
    await this.getDict('renew')
    await this.getDict('Intermediary')
  },
  methods: {
    setAuditType() {
      this.formData.auditType = 1
    },
    async setPrincipal(){
      const res = await getCustomerList({ page: 1, pageSize: 999 })
      res.data.list && res.data.list.forEach(item => {
        if(this.formData.merchant===item.name){
          this.formData.principal = item.principal
        }
      })
    },
    gethousingSelects (){
      console.log(this.housingSelects)
    },
    async setCustomerOptions() {      
      this.customerOptions = []
      const res = await getCustomerList({ page: 1, pageSize: 999 })
      res.data.list &&res.data.list.forEach(item => {
        if(item.audit === 0){
          const option = {
            label: item.name,
            value: item.name
          }
          this.customerOptions.push(option)
        }
      })
    },
    async setMiddlemanOptions() {
      this.middlemanOptions = []
      const resIntermediaryCompany = await getIntermediaryCompanyList({ page: 1, pageSize: 999 })
      const resMiddleman = await getMiddlemanList({ page: 1, pageSize: 999 })
      resIntermediaryCompany.data.list && resIntermediaryCompany.data.list.forEach(item => {
        const option = {
          label: item.name,
          value: item.name,
          children: [],
        }
        resMiddleman.data.list && resMiddleman.data.list.forEach(middlemanItem => {
          if(middlemanItem.company === option.value && middlemanItem.auditType===0){
            const middlemanOption = {
              label: middlemanItem.name,
              value: middlemanItem.name,
              children: [],
            }
            option.children.push(middlemanOption)
          }
        })
        if(option.children!=""){
          this.middlemanOptions.push(option)
        }
      })
    },
    async setHousingOptions() {
      this.housingOptions = []
      const resProject = await getProjectInformationList({ page: 1, pageSize: 999 })
      const resBuilding = await getBuildingInformationList({ page: 1, pageSize: 999 })
      const resFloor = await getFloorInformationList({ page: 1, pageSize: 999 })
      const resHousing = await getHousingMaintenanceList({ page: 1, pageSize: 999 })
      resProject.data.list && resProject.data.list.forEach(item => {
        const option = {
          label: item.name,
          value: item.name,
          children: [],
        }
        resBuilding.data.list && resBuilding.data.list.forEach(buildingItem => {
        if(buildingItem.project === option.value){
          const buildingOption = {
            label: buildingItem.name,
            value: buildingItem.name,
            children: []
          }
          resFloor.data.list && resFloor.data.list.forEach(floorItem => {
            if(floorItem.build === buildingOption.value && floorItem.project === option.value){
              const floorOption = {
                label: floorItem.name,
                value: floorItem.name,
                children: []
              }
              resHousing.data.list && resHousing.data.list.forEach(housingItem => {
                if(housingItem.floor === floorOption.value && housingItem.build === buildingOption.value && housingItem.project === option.value){
                  const housingOption = {
                    label: housingItem.name,
                    value: housingItem.name,
                    children: []
                  }
                  floorOption.children.push(housingOption)
                }
              })
              if(floorOption.children.length != 0)
                buildingOption.children.push(floorOption)
            }
          })
          if(buildingOption.children.length != 0)
            option.children.push(buildingOption)
        }
      })
        if(option.children.length != 0)
          this.housingOptions.push(option)
      })

      console.log(this.housingOptions)
    },
    async setHousingOptions1() {
      this.housingOptions = []
      const resProject = await getProjectInformationList({ page: 1, pageSize: 999 })
      const resBuilding = await getBuildingInformationList({ page: 1, pageSize: 999 })
      const resFloor = await getFloorInformationList({ page: 1, pageSize: 999 })
      const resHousing = await getHousingMaintenanceList({ page: 1, pageSize: 999 })
      resProject.data.list && resProject.data.list.forEach(item => {
        const option = {
          label: item.name,
          value: item.name,
          children: [],
        }
        resBuilding.data.list && resBuilding.data.list.forEach(buildingItem => {
        if(buildingItem.project === option.value){
          const buildingOption = {
            label: buildingItem.name,
            value: buildingItem.name,
            children: []
          }
          resFloor.data.list && resFloor.data.list.forEach(floorItem => {
            if(floorItem.build === buildingOption.value && floorItem.project === option.value){
              const floorOption = {
                label: floorItem.name,
                value: floorItem.name,
                children: []
              }
              resHousing.data.list && resHousing.data.list.forEach(housingItem => {
                if(housingItem.floor === floorOption.value && housingItem.build === buildingOption.value && housingItem.project === option.value){
                  const housingOption = {
                    label: housingItem.name,
                    value: housingItem.name,
                    children: []
                  }
                  floorOption.children.push(housingOption)
                }
              })
              buildingOption.children.push(floorOption)
            }
          })
          option.children.push(buildingOption)
        }
      })
        this.housingOptions.push(option)
      })
    },
    changeAuditType(row) {
      this.$confirm('确定要审核吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.changeAuditType2(row)
      })
    },
    changeAuditType1(row) {
      this.$confirm('确定要取消审核吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.changeAuditType2(row)
      })
    },
    async changeAuditType2(row) {
      const res = await findPact({ ID: row.ID })
      if (res.code === 0) {
        this.formData = res.data.rept
        this.formData.auditType = this.formData.auditType===0 ? 1 : 0
      }
      let flag = await updatePact(this.formData)
      if (flag.code === 0) {
        this.$message({
          type: 'success',
          message: '更改成功'
        })
        this.getTableData()
      }
    },
    handleRowClick(row) {
      row.expanded = !row.expanded;
      this.$refs.multipleTable.toggleRowExpansion(row, row.expanded);
    },
    async setContractNumberOptions() {
      this.contractNumberOptions = []
      const res = await getPactList({ page: 1, pageSize: 999 })
      res.data.list && res.data.list.forEach(item => {
        const option = {
          label: item.contractNumber,
          value: item.contractNumber
        }
        this.contractNumberOptions.push(option)
      })
    },
    async setMerchantOptions() {
      let list = []
      const res = await getPactList({ page: 1, pageSize: 999 })
      if(this.searchInfo.merchant === undefined || this.searchInfo.merchant === '') {
        res.data.list && res.data.list.forEach(item => {
          if(item.merchant != '')
            list.push(item.merchant)
        })
      }
      else {
        res.data.list && res.data.list.forEach(item => {
          if(item.merchant === this.searchInfo.merchant){
            if(item.merchant != '')
              list.push(item.merchant)          
          }
        })
      }
      let newArr = list.filter((item, index) => list.indexOf(item) === index); 

      this.merchantOptions = []
      newArr && newArr.forEach(item => {
        const option ={
          label: item,
          value: item
        }
        this.merchantOptions.push(option)
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
        this.deletePact(row)
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
      const res = await deletePactByIds({ ids })
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
    async updatePact(row) {
      const res = await findPact({ ID: row.ID })
      this.type = 'update'
      if (res.code === 0) {
        this.formData = res.data.rept
        this.dialogFormVisible = true
      }
    },
    closeDialog() {
      this.$refs.formData.resetFields()
      this.dialogFormVisible = false
      this.housingSelects = []
      this.middlemanSelects = []
      this.beginToEnd = []
      this.formData = {
        contractNumber: '',
        project: '',
        building: '',
        floor: '',
        housing: '',
        merchant: '',
        contractType: undefined,
        contractSigning: undefined,
        renew: undefined,
        associatedContractNumber: '',
        intermediary: undefined,
        agency: '',
        intermediaryContact: '',
        principal: '',
        deliveryDate: new Date(),
        startTime: new Date(),
        endTime: new Date(),
        paymentCycle: undefined,
        univalence: undefined,
        rent: 0,
        serviceCharge: 0,
        propertyManagementFee: 0,
        contractGrandTotal: 0,
        setUpFee: 0,
        earnestMoney: 0,
        remark: '',
        auditType: undefined,
      }
    },
    async deletePact(row) {
      const res = await deletePact({ ID: row.ID })
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
      this.housingSelects && this.housingSelects.forEach(item => {
        console.log(item)
      })
      console.log(this.middlemanSelects)
      console.log(this.beginToEnd)
      let res
      switch (this.type) {
        case 'create':
          res = await createPact(this.formData)
          break
        case 'update':
          res = await updatePact(this.formData)
          break
        default:
          res = await createPact(this.formData)
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
