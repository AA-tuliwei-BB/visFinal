<template>

</template>

<script setup>
import { onMounted, ref } from 'vue';
import * as echarts from 'echarts/core';
import { MapChart } from 'echarts/charts';
import { TooltipComponent, VisualMapComponent } from 'echarts/components';
import { CanvasRenderer } from 'echarts/renderers';
import chinaMapData from '@/assets/china.json';
import { getHeat, postFilter } from '@/api/api';
import { useStore } from 'vuex'
import { computed, watch } from 'vue'

const store = useStore()

const selectedCategories = computed({
    get: () => store.state.selectedCategories,
    set: (value) => store.commit('setSelectedCategories', value)
})

const selectedEthnicity = computed({
    get: () => store.state.selectedEthnicity,
    set: (value) => store.commit('setSelectedEthnicity', value)
})

const selectedBatches = computed({
    get: () => store.state.selectedBatches,
    set: (value) => store.commit('setSelectedBatches', value)
})

const selectedKeyword = computed({
    get: () => store.state.selectedKeyword,
    set: (value) => store.commit('setSelectedKeyword', value)
})

const selectedProvinces = computed({
    get: () => store.state.selectedProvinces,
    set: (value) => store.commit('setSelectedProvinces', value)
})

const updateHeatMap = computed({
    get: () => store.state.updateHeatMap,
    set: (value) => store.commit('setUpdateHeatMap', value)
})

echarts.use([TooltipComponent, VisualMapComponent, MapChart, CanvasRenderer]);

const mapChart = ref(null);
const checkAll = ref(false)
const indeterminate = ref(false)

const response = (val) => {
    console.log(val)
    if (val.length === 0) {
        checkAll.value = false
        indeterminate.value = false
    } else if (val.length === provinces.length) {
        checkAll.value = true
        indeterminate.value = false
    } else {
        indeterminate.value = true
    }
    updateMap();
}

watch(selectedProvinces, response)



const handleCheckAll = (val) => {
    indeterminate.value = false
    if (val) {
        selectedProvinces.value = provinces.map((_) => _)
    } else {
        selectedProvinces.value = []
    }
}

const provinces = ['上海', '云南', '内蒙古', '北京', '吉林', '四川', '天津', '宁夏', '安徽', '山东', '山西', '广东', '广西', '新疆', '江苏', '江西', '河北', '河南', '浙江', '海南', '湖北', '湖南', '澳门', '甘肃', '福建', '西藏', '贵州', '辽宁', '重庆', '陕西', '青海', '香港', '黑龙江']

watch(updateHeatMap, async (val) => {
    const jsonData = await getHeat();

    const mapData = provinces.map(province => {
        let value = 0;
        if (jsonData && jsonData.data) {
            const item = jsonData.data.find(item => item.province === province);
            value = item ? item.heat : 0;
        }
        return {
            name: province,
            value: value,
        };
    });

    const chart = echarts.getInstanceByDom(mapChart.value);
    const option = chart.getOption();
    option.series[0].data = mapData;
    let maxVal = Math.max(...mapData.map(item => item.value));
    option.visualMap[0].max = maxVal + 1;
    option.visualMap[0].range = [0, maxVal + 1];
    console.log(option.visualMap[0])
    chart.setOption(option, true);
    updateHeatMap.value = false;
})
onMounted(async () => {
    const chart = echarts.init(mapChart.value);

    echarts.registerMap('china', chinaMapData);

    const jsonData = await getHeat();
    const mapData = provinces.map(province => {
        let value = 0;
        if (jsonData && jsonData.data) {
            const item = jsonData.data.find(item => item.province === province);
            value = item ? item.heat : 0;
        }
        return {
            name: province,
            value: value,
        };
    });
    let maxVal = Math.max(...mapData.map(item => item.value));

    let option = {
        tooltip: {
            trigger: 'item'
        },
        visualMap: {
            min: 0,
            max: maxVal + 1,
            left: 'left',
            top: 'bottom',
            text: ['High', 'Low'],
            range: [0, maxVal],
            inRange: {
                color: ['white', 'blue']
            },
            calculable: true,
            textStyle: {
                color: "white"
            }
        },
        series: [
            {
                name: '中国',
                type: 'map',
                map: 'china',
                roam: true,
                selectedMode: false,
                data: mapData, // 使用转换后的数据
                emphasis: {
                    itemStyle: {
                        areaColor: null,
                        borderColor: "purple",
                        borderWidth: 1
                    },
                }
            }
        ]
    };

    chart.setOption(option);


    // 添加click事件监听器
    chart.on('click', function (params) {
        const provinceName = params.name;
        option = chart.getOption();
        // 获取省份在selectedProvinces数组中的索引
        const selectedIndex = selectedProvinces.value.indexOf(provinceName);
        // 获取省份在option.series[0].data数组中的索引
        const dataIndex = option.series[0].data.findIndex(item => item.name === provinceName);

        if (selectedIndex > -1) {
            selectedProvinces.value.splice(selectedIndex, 1);
        } else {
            selectedProvinces.value.push(provinceName);
        }

        response(selectedProvinces.value)
    });
});

function submit() {
    postFilter(selectedCategories, selectedBatches, selectedEthnicity, selectedKeyword, selectedProvinces, updateHeatMap);
}

function updateMap() {
    const chart = echarts.getInstanceByDom(mapChart.value);
    const option = chart.getOption();

    // 清除所有省份的自定义样式
    option.series[0].data.forEach(item => {
        item.itemStyle = {};
    });

    // 为选中的省份设置自定义样式
    selectedProvinces.value.forEach(provinceName => {
        const dataIndex = option.series[0].data.findIndex(item => item.name === provinceName);
        option.series[0].data[dataIndex].itemStyle = {
            borderColor: 'purple', // 选中时的边框颜色
            borderWidth: 3 // 选中时的边框宽度
        };
    });

    chart.setOption(option); // 更新选项
}
</script>

<style scoped>
.mapItem {
    position: absolute;
    left: 5%;
    top: 15%;
    width: 95%;
    height: 85%;
}

.submitButton {
    position: absolute;
    right: 3%;
    top: 6%;
}

.filterItem {
    position: absolute;
    left: 25%;
    top: 6%;
    width: 60%;
}
</style>