<template>
  <el-row>
    <el-col :span="24"><div class="grid-content"><font style="font-size:30px;font-weight:bolder">{{data.city}}</font></div></el-col>
  </el-row>
  <el-row :gutter="20">
    <el-col :span="12"><div class="grid-content"><font style="font-size:14px">实时气温</font><br><font style="font-size:14px">{{data.temperature}}°C</font></div></el-col>
    <el-col :span="12"><div class="grid-content"><font style="font-size:14px">天气</font><br><font style="font-size:14px">{{data.weather}}</font></div></el-col>
  </el-row>
  <el-row :gutter="20">
    <el-col :span="12"><div class="grid-content"><font style="font-size:14px">风向</font><br><font style="font-size:14px">{{data.windDirection}}</font></div></el-col>
    <el-col :span="12"><div class="grid-content"><font style="font-size:14px">风力</font><br><font style="font-size:14px">{{data.windPower}}</font></div></el-col>
  </el-row>
  <el-row :gutter="20">
    <el-col :span="12"><div class="grid-content"><font style="font-size:14px">空气湿度</font><br><font style="font-size:14px">{{data.humidity}}%</font></div></el-col>
    <el-col :span="12"><div class="grid-content"><font style="font-size:14px">发布时间</font><br><font style="font-size:14px">{{ formatDate(data.reportTime) }}</font></div></el-col>
  </el-row>
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
export default {
  mixins: [infoList],
  data() {
    return {
      data: {
        province: 'province',             //省份名
        city: 'city',                     //城市名
        humidity: 'humidity',             //空气湿度（百分比）
        reportTime: 'reportTime',         //数据发布的时间
        temperature: 'temperature',       //实时气温，单位：摄氏度
        weather: 'weather',               //天气预报，详见天气现象列表
        windDirection: 'windDirection',   //风向，风向编码对应描述
        windPower: 'windPower'            //风力，风力编码对应风力级别，单位：级
      }
    };
  },
  created() {
    this.getLngLatLocation();
  },
  mounted(){
    this.$nextTick(() => {
      setInterval(this.getLngLatLocation, 5000);
    });
  },
  methods:{
    // ip 定位
    getLngLatLocation() {
      var that = this
      var map = new AMap.Map('container', {
        resizeEnable: true
      })
      AMap.plugin("AMap.CitySearch", function () {
        var citySearch = new AMap.CitySearch();
        var result = '';
        citySearch.getLocalCity(function (status, result) {
          if (status === "complete" && result.info === "OK") {
            // 查询成功，result即为当前所在城市信息
            that.data.city = result.city;
            AMap.plugin('AMap.Weather', function() {
	              //创建天气查询实例
	              var weather = new AMap.Weather();
	              //执行实时天气信息查询
	              weather.getLive(result.city, function(err, data) {
	                  // console.log(err, data);
                    that.data.province = data.province
                    that.data.city = data.city
                    that.data.humidity = data.humidity
                    that.data.reportTime = data.reportTime
                    that.data.temperature = data.temperature
                    that.data.weather = data.weather
                    that.data.windDirection = data.windDirection
                    that.data.windPower = data.windPower
	              });
            });
          }
        });
      });
    },
    
  },
};
</script>