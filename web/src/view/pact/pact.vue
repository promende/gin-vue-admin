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
            <!-- <el-button style="margin-left:10px" size="mini" type="primary" icon="download" @click="exportExcel">导出</el-button> -->
        </div>
        <el-table
        id="rebateSetTable"
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
            <el-descriptions :column="5" border>
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
              <el-descriptions-item label="负责人" label-align="center" align="left" width="10px" span="2">
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
              <template v-if="scope.row.univalence===0">
                <el-descriptions-item label="日租金" label-align="center" align="left" width="10px">
                  {{ scope.row.rent }}
                </el-descriptions-item>
                <el-descriptions-item label="日服务费" label-align="center" align="left" width="10px">
                  {{ scope.row.serviceCharge }}
                </el-descriptions-item>
                <el-descriptions-item label="日物管费" label-align="center" align="left" width="80px">
                  {{ scope.row.propertyManagementFee }}
                </el-descriptions-item>
              </template>
              <template v-if="scope.row.univalence===1">
                <el-descriptions-item label="月租金" label-align="center" align="left" width="10px">
                  {{ scope.row.rent }}
                </el-descriptions-item>
                <el-descriptions-item label="月服务费" label-align="center" align="left" width="10px">
                  {{ scope.row.serviceCharge }}
                </el-descriptions-item>
                <el-descriptions-item label="月物管费" label-align="center" align="left" width="80px">
                  {{ scope.row.propertyManagementFee }}
                </el-descriptions-item>
              </template>
              <el-descriptions-item label="合同总金额" label-align="center" align="left" width="80px" span="2">
                {{ parseFloat(scope.row.contractGrandTotal).toFixed(2) }}
              </el-descriptions-item>
              <el-descriptions-item label="设置费" label-align="center" align="left" width="80px">
                {{ scope.row.setUpFee }}
              </el-descriptions-item>
              <el-descriptions-item label="保证金" label-align="center" align="left" width="10px">
                {{ scope.row.earnestMoney }}
              </el-descriptions-item>
              <el-descriptions-item label="审核状态" label-align="center" align="left" width="80px">
                <el-tag v-if="scope.row.auditType===0" class="ml-2" type="success">{{ filterDict(scope.row.auditType,"auditType") }}</el-tag>
                <el-tag v-else class="ml-2" type="danger">{{ filterDict(scope.row.auditType,"auditType") }}</el-tag>
              </el-descriptions-item>
              <el-descriptions-item label="创建日期" label-align="center" align="left" width="80px">
                {{ formatDate(scope.row.CreatedAt) }}
              </el-descriptions-item>  
              <el-descriptions-item label="备注" label-align="center" align="left" width="80px">
                {{ scope.row.remark }}
              </el-descriptions-item>
            </el-descriptions>
          </template>
        </el-table-column>
        <el-table-column align="left" label="合同编号" prop="contractNumber" width="160" />
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
            <el-button type="text" icon="delete" size="mini" @click="deleteRow(scope.row);handleRowClick(scope.row)">删除</el-button>
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
            :props="{ multiple:false, label:'label', value:'value', disabled:'disabled', checkStrictly:false}"
            @visible-change="setHousingOptions" 
            :disabled="this.formData.auditType===0"
          />
        </el-form-item>
        <el-row style="margin-top:-10px">
            <el-col :span="12">
              <el-form-item label="商家名称" prop="b">
                <el-select v-model="formData.merchant" placeholder="请选择" style="width:100%" default-first-option clearable filterable @visible-change="setCustomerOptions" @change="setPrincipal" :disabled="this.formData.auditType===0">
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
                <el-select v-model="formData.contractType" placeholder="请选择" style="width:100%" clearable :disabled="this.formData.auditType===0">
                  <el-option v-for="(item,key) in contractTypeOptions" :key="key" :label="item.label" :value="item.value" />
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="合同签署" prop="d">
                <el-select v-model="formData.contractSigning" placeholder="请选择" style="width:100%" clearable :disabled="this.formData.auditType===0">
                  <el-option v-for="(item,key) in contractSigningOptions" :key="key" :label="item.label" :value="item.value" />
                </el-select>
              </el-form-item>
            </el-col>
        </el-row>
        <el-row style="margin-top:-20px">
            <el-col :span="12">
              <el-form-item label="是否续签" prop="e">
                <el-select v-model="formData.renew" placeholder="请选择" style="width:30%" clearable  @visible-change="setAssociatedContractNumber" :disabled="this.formData.auditType===0">
                  <el-option v-for="(item,key) in renewOptions" :key="key" :label="item.label" :value="item.value" />
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="关联合同编号" v-if="this.formData.renew === 0" prop="associatedContractNumber">
                <el-input v-model="formData.associatedContractNumber" clearable placeholder="请输入"  :disabled="this.formData.auditType===0"/>
              </el-form-item>
            </el-col>
        </el-row>
        <el-row style="margin-top:-20px">
            <el-col :span="12">
              <el-form-item label="是否中介介入" prop="g">
                <el-select v-model="formData.intermediary" placeholder="请选择" style="width:30%" clearable @visible-change="setMiddlemanSelects" :disabled="this.formData.auditType===0">
                  <el-option v-for="(item,key) in IntermediaryOptions" :key="key" :label="item.label" :value="item.value" />
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="中介人" v-if="this.formData.intermediary===0" prop="h">
                <el-cascader
                  clearable
                  v-model="this.middlemanSelects"
                  style="width:100%"
                  :options="middlemanOptions"
                  :props="{ multiple:false, label:'label', value:'value', disabled:'disabled'}"
                  @visible-change="setMiddlemanOptions"
                  :disabled="this.formData.auditType===0"
                />
              </el-form-item>
            </el-col>
        </el-row>
        <el-row style="margin-top:-20px">
            <el-col :span="12">
              <el-form-item label="交付日" style="margin-top:-10px" prop="deliveryDate">
                <el-date-picker v-model="formData.deliveryDate" type="date" style="width:100%" placeholder="选择日期" clearable  :disabled="this.formData.auditType===0"/>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="合同起止时间" style="margin-top:-10px" prop="beginToEnd">
                <el-date-picker v-model="this.beginToEnd" type="daterange" style="width:100%"
                  start-placeholder="合同开始时间"
                  range-separator="-"
                  end-placeholder="合同结束时间"
                  unlink-panels
                  @change="setContractGrandTotal"
                  :disabled="this.formData.auditType===0"
                />
              </el-form-item>
            </el-col>
        </el-row>
        <el-row style="margin-top:-20px">
            <el-col :span="12">
              <el-form-item label="单价" prop="univalence">
                <el-select v-model="formData.univalence" placeholder="请选择" style="width:100%" clearable @visible-change="setFeeClear()" :disabled="this.formData.auditType===0">
                  <el-option v-for="(item,key) in univalenceOptions" :key="key" :label="item.label" :value="item.value" />
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="支付周期" prop="paymentCycle">
                <el-select v-model="formData.paymentCycle" placeholder="请选择" style="width:100%" clearable :disabled="this.formData.auditType===0">
                  <el-option v-for="(item,key) in paymentCycleOptions" :key="key" :label="item.label" :value="item.value" />
                </el-select>
              </el-form-item>
            </el-col>
        </el-row>
        <template v-if="formData.univalence===0">
          <el-form-item label="日租金" prop="rent">
            <el-input-number v-model="formData.rent"  style="width:100%" :precision="2" clearable @change="setContractGrandTotal" :disabled="this.formData.auditType===0"/>
          </el-form-item>
          <el-form-item label="日服务费" prop="serviceCharge">
            <el-input-number v-model="formData.serviceCharge"  style="width:100%" :precision="2" clearable  @change="setContractGrandTotal" :disabled="this.formData.auditType===0"/>
          </el-form-item>
          <el-form-item label="日物管费" prop="propertyManagementFee">
            <el-input-number v-model="formData.propertyManagementFee"  style="width:100%" :precision="2" clearable  @change="setContractGrandTotal" :disabled="this.formData.auditType===0"/>
          </el-form-item>
          <el-form-item label="合同总金额" prop="contractGrandTotal">
            <el-input-number v-model="formData.contractGrandTotal"  style="width:100%" :precision="2" clearable disabled/>
          </el-form-item>
          <el-form-item label="设置费" prop="setUpFee">
            <el-input-number v-model="formData.setUpFee"  style="width:100%" :precision="2" clearable  :disabled="this.formData.auditType===0"/>
          </el-form-item>
          <el-form-item label="保证金" prop="earnestMoney">
            <el-input-number v-model="formData.earnestMoney"  style="width:100%" :precision="2" clearable  :disabled="this.formData.auditType===0"/>
          </el-form-item>
        </template>
        <template v-if="formData.univalence===1">
          <el-form-item label="月租金" prop="rent">
            <el-input-number v-model="formData.rent"  style="width:100%" :precision="2" clearable  @change="setContractGrandTotal" :disabled="this.formData.auditType===0"/>
          </el-form-item>
          <el-form-item label="月服务费" prop="serviceCharge">
            <el-input-number v-model="formData.serviceCharge"  style="width:100%" :precision="2" clearable  @change="setContractGrandTotal" :disabled="this.formData.auditType===0"/>
          </el-form-item>
          <el-form-item label="月物管费" prop="propertyManagementFee">
            <el-input-number v-model="formData.propertyManagementFee"  style="width:100%" :precision="2" clearable  @change="setContractGrandTotal" :disabled="this.formData.auditType===0"/>
          </el-form-item>
          <el-form-item label="合同总金额" prop="contractGrandTotal">
            <el-input-number v-model="formData.contractGrandTotal"  style="width:100%" :precision="2" clearable disabled/>
          </el-form-item>
          <el-form-item label="设置费" prop="setUpFee">
            <el-input-number v-model="formData.setUpFee"  style="width:100%" :precision="2" clearable  :disabled="this.formData.auditType===0"/>
          </el-form-item>
          <el-form-item label="保证金" prop="earnestMoney">
            <el-input-number v-model="formData.earnestMoney"  style="width:100%" :precision="2" clearable :disabled="this.formData.auditType===0"/>
          </el-form-item>
        </template>
        <el-form-item label="备注" style="margin-top:-10px">
          <el-input v-model="formData.remark" clearable placeholder="请输入" :disabled="this.formData.auditType===0"/>
        </el-form-item>
        <!-- <el-form-item label="审核状态">
          <el-select v-model="formData.auditType" placeholder="请选择" style="width:100%" clearable disabled>
            <el-option v-for="(item,key) in auditTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item> -->
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
import * as XLSX from 'xlsx';
import FileSaver from 'file-saver'
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
        associatedContractNumber:                [{ required: true, message: '请输入关联合同编号',  trigger: 'blur' }],
        g:                [{ required: true, message: '请输入是否中介介入',  trigger: 'blur' }],
        h:                [{ required: true, message: '请输入中介人',  trigger: 'blur' }],
        deliveryDate:     [{ type: 'date', required: true, message: '请选择交付日', trigger: 'change'}],
        beginToEnd:       [{ type: '1', required: true, message: '请选择合同起止时间', trigger: 'change'}],
        paymentCycle:                [{ required: true, message: '请输入支付周期',  trigger: 'blur' }],
        univalence:                [{ required: true, message: '请输入单价',  trigger: 'blur' }],
        rent:             [{ required: true, message: '请输入计租面积（㎡)',  trigger: 'blur' }],
        serviceCharge:                [{ required: true, message: '请输入实用面积（㎡）',  trigger: 'blur' }],
        propertyManagementFee:             [{ required: true, message: '请输入计租面积（㎡)',  trigger: 'blur' }],
        contractGrandTotal:                [{ required: true, message: '请输入实用面积（㎡）',  trigger: 'blur' }],
        setUpFee:             [{ required: true, message: '请输入计租面积（㎡)',  trigger: 'blur' }],
        earnestMoney:                [{ required: true, message: '请输入实用面积（㎡）',  trigger: 'blur' }],
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
    async exportExcel() {
      /* generate workbook object from table */
      let oldPageSize = this.pageSize
      this.pageSize = 100
      await this.getTableData()
      let xlsxParam = { raw: true } // 导出的内容只做解析，不进行格式转换
      let wb = XLSX.utils.table_to_book(document.querySelector('#rebateSetTable'), xlsxParam);
      let ref = wb.Sheets.Sheet1['!ref']
      let flag = 0
      let maxCols = ""
      let maxRows = ""
      for(let i = 0; i < ref.length; i++){
        if(ref[i] === ":"){
          flag = 1;
          continue
        }
        if(flag === 1){
          if(ref[i]>="A"&&ref[i]<="Z"){
            maxCols += ref[i]
          }
          else{
            maxRows += ref[i];
          }
        }
      }
      for(let i = 0; i<=maxRows; i++){
        let cell = maxCols + i
        wb.Sheets.Sheet1[cell] = ''
        let firstCell = 'A' + i
        if(i === 1){
            wb.Sheets.Sheet1[firstCell] = {
            t: "s",
            v: "序号"
          }
        }
        else{
            wb.Sheets.Sheet1[firstCell] = {
            t: "s",
            v: (i - 1).toString()
          }
        }
      }
      /* get binary string as output */
      console.log(wb.Sheets.Sheet1)
      let wbout = XLSX.write(wb, { bookType: 'xlsx', bookSST: true, type: 'array' });
      // console.log(wbout)
      try {
          FileSaver.saveAs(new Blob([wbout], { type: 'application/octet-stream' }), '合同台账表.xlsx');
      } catch (e)
      {
        if (typeof console !== 'undefined')
            console.log(e, wbout)
      }

      this.pageSize = oldPageSize
      await this.getTableData()
      
      return wbout
    },
    setFeeClear() {
      this.formData.rent = 0
      this.formData.serviceCharge = 0 
      this.formData.propertyManagementFee = 0
      this.formData.contractGrandTotal = 0
    },
    setContractGrandTotal() {
      if(this.beginToEnd.length === 2){
        let begin = this.beginToEnd[0]
        let end = this.beginToEnd[1]
        if(this.formData.univalence === 0){
          let day = (end - begin) / (1000 * 60 * 60 * 24) + 1
          this.formData.contractGrandTotal = day * (this.formData.rent + this.formData.serviceCharge + this.formData.propertyManagementFee)
        }
        else if (this.formData.univalence === 1){
          let endDate = new Date(end)
          let endYear = endDate.getFullYear()
          let endMonth = endDate.getMonth() + 1
          let endDay = endDate.getDate()
          
          let beginDate = new Date(begin)
          let beginYear = beginDate.getFullYear()
          let beginMonth = beginDate.getMonth() + 1
          let beginDay = beginDate.getDate()

          let diffYear = endYear - beginYear
          let diffMonth = endMonth - beginMonth
          let diffDay = endDay - beginDay

          let month = diffYear * 12 + diffMonth
          if(diffDay > -1) {
            month = month + (diffDay + 1) * 12 / 365
          }
          else if (diffDay < -1) {
            month = month - 1
            let preEnd = new Date(endYear, endMonth-1, beginDay)
            month = month + (end - preEnd + 1) / (1000 * 60 * 60 * 24)
          }
          this.formData.contractGrandTotal = month * (this.formData.rent + this.formData.serviceCharge + this.formData.propertyManagementFee)   
        }
      }
    },
    rounding(row, column) {
      return parseFloat(row[column.property]).toFixed(2)
    },
    async setContractNumber(){
      this.formData.contractNumber = ''
      const res = await getProjectInformationList({ page: 1, pageSize: 999 })
      res.data.list && res.data.list.forEach(item => {
        if(this.formData.project === item.name){
          this.formData.contractNumber = item.abbreviation
        }
      })
      let temp = new Date()
      let year = temp.getFullYear()
      let month = temp.getMonth() + 1
      let day = temp.getDate()
      month = month > 10 ? month.toString() : '0'+month.toString()
      day = day > 10 ? day.toString() : '0'+day.toString()
      this.formData.contractNumber = this.formData.contractNumber + year + month + day
      const resPact = await getPactList({ page: 1, pageSize: 999 })
      let s = ''
      resPact.data.list && resPact.data.list.forEach(item => {
        if(this.formData.contractNumber === item.contractNumber.slice(0,-4)){
          s = item.contractNumber.slice(-4)
        }
      })
      if(s === ''){
        this.formData.contractNumber = this.formData.contractNumber + '0001'
      }
      else{
        let number = parseInt(s) + 1
        for(var len = (number + "").length; len < 4; len++) {
            number = "0" + number;            
        }
        this.formData.contractNumber = this.formData.contractNumber + number
      }
    },
    setAssociatedContractNumber() {
      if(this.formData.renew===1){
        this.formData.associatedContractNumber = ''
      }
    },
    setMiddlemanSelects() {
      if(this.formData.intermediary===1){
        this.middlemanSelects = []
      }
    },
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
                    disabled: '',
                    children: []
                  }
                  if(housingItem.state === 1){
                      housingOption.disabled = "disabled"
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
        this.housingSelects = []
        this.housingSelects.push(this.formData.project)
        this.housingSelects.push(this.formData.building)
        this.housingSelects.push(this.formData.floor)
        this.housingSelects.push(this.formData.housing)
        this.setHousingOptions()
        this.middlemanSelects = []
        this.middlemanSelects.push(this.formData.agency)
        this.middlemanSelects.push(this.formData.intermediaryContact)
        this.setMiddlemanOptions()
        this.beginToEnd = []
        this.beginToEnd.push(this.formData.startTime)
        this.beginToEnd.push(this.formData.endTime)

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
      // this.housingSelects && this.housingSelects.forEach(item => {
      //   console.log(item)
      // })
      // console.log(this.middlemanSelects)
      // console.log(this.beginToEnd)
      // console.log(this.housingSelects.length)
      // console.log(this.beginToEnd.length)
      // console.log(this.formData)
      if(this.type === 'update'){
        if(this.housingSelects.length != 0 && this.formData.merchant != '' 
        && this.formData.contractType != undefined && this.formData.contractSigning != undefined 
        && this.formData.renew != undefined && this.formData.intermediary != undefined 
        && this.formData.deliveryDate != '' && this.beginToEnd.length != 0 
        && this.formData.univalence != undefined && this.formData.paymentCycle != undefined) {
          // console.log(this.formData.intermediary)
          // console.log(this.middlemanSelects.length)
          // console.log(this.formData.renew)
          // console.log(this.formData.associatedContractNumber)
          if(this.formData.intermediary === 0 && this.middlemanSelects.length ===0 || this.formData.renew === 0 && this.formData.associatedContractNumber === ""){
            this.$message({
              type: 'warning',
              message: '请填写必填项'
            })
          }
          else{
            this.formData.project = this.housingSelects[0]
            this.formData.building = this.housingSelects[1]
            this.formData.floor = this.housingSelects[2]
            this.formData.housing = this.housingSelects[3]
            this.formData.agency = this.middlemanSelects[0]
            this.formData.intermediaryContact = this.middlemanSelects[1]
            this.formData.startTime = this.beginToEnd[0]
            this.formData.endTime = this.beginToEnd[1]
            console.log(this.formData)
            let res = await updatePact(this.formData)
            if (res.code === 0) {
              this.$message({
                type: 'success',
                message: '更改成功'
              })
              this.closeDialog()
              this.getTableData()
            }
          }
        }
        else{
          this.$message({
            type: 'warning',
            message: '请填写必填项'
          })
        }
      }
      else{
        if(this.housingSelects.length != 0 && this.formData.merchant != '' 
        && this.formData.contractType != undefined && this.formData.contractSigning != undefined 
        && this.formData.renew != undefined && this.formData.intermediary != undefined 
        && this.formData.deliveryDate != '' && this.beginToEnd.length != 0 
        && this.formData.univalence != undefined && this.formData.paymentCycle != undefined) {
          // console.log(this.formData.intermediary)
          // console.log(this.middlemanSelects.length)
          // console.log(this.formData.renew)
          // console.log(this.formData.associatedContractNumber)
          if(this.formData.intermediary === 0 && this.middlemanSelects.length ===0 || this.formData.renew === 0 && this.formData.associatedContractNumber === ""){
            this.$message({
              type: 'warning',
              message: '请填写必填项'
            })
          }
          else{
            this.formData.project = this.housingSelects[0]
            this.formData.building = this.housingSelects[1]
            this.formData.floor = this.housingSelects[2]
            this.formData.housing = this.housingSelects[3]
            this.formData.agency = this.middlemanSelects[0]
            this.formData.intermediaryContact = this.middlemanSelects[1]
            this.formData.startTime = this.beginToEnd[0]
            this.formData.endTime = this.beginToEnd[1]
            await this.setContractNumber()
            console.log(this.formData)
            let res = await createPact(this.formData)
            if (res.code === 0) {
              this.$message({
                type: 'success',
                message: '创建成功'
              })
              this.closeDialog()
              this.getTableData()
            }
          }
        }
        else{
          this.$message({
            type: 'warning',
            message: '请填写必填项'
          })
        }
      }

      // let res
      // switch (this.type) {
      //   case 'create':
      //     res = await createPact(this.formData)
      //     break
      //   case 'update':
      //     res = await updatePact(this.formData)
      //     break
      //   default:
      //     res = await createPact(this.formData)
      //     break
      // }
      // if (res.code === 0) {
      //   this.$message({
      //     type: 'success',
      //     message: '创建/更改成功'
      //   })
      //   this.closeDialog()
      //   this.getTableData()
      // }
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
