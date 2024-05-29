<template>
    <el-radio-group v-model="chartType" class="groupItem">
        <el-radio-button v-for="(value, key) in groups" :key="key" :label="key" :value="value" />
    </el-radio-group>
    <div ref="variousChart" class="chart"></div>
</template>

<script setup>
import { onMounted, ref } from 'vue';
import * as echarts from 'echarts/core';
import { PieChart, BarChart } from 'echarts/charts';
import { TooltipComponent, VisualMapComponent, LegendComponent } from 'echarts/components';
import { CanvasRenderer } from 'echarts/renderers';
import { getChart } from '@/api/api';
import { useStore } from 'vuex'
import { computed, watch } from 'vue'
import 'echarts-wordcloud';

const store = useStore()

const groups = {
    "类别": "category",
    "申报批次": "batch",
    "民族": "ethnic",
    "关键词": "keyword",
};

let chartType = ref("category");

const updateCharts = computed({
    get: () => store.state.updateCharts,
    set: (value) => store.commit('setUpdateCharts', value)
})

echarts.use([TooltipComponent, VisualMapComponent, PieChart, LegendComponent, BarChart, CanvasRenderer]);

const variousChart = ref(null);
let chart;

const update = async (val) => {
    const jsonData = await getChart(chartType.value);
    let chartName = Object.entries(groups).find(([_, value]) => value === chartType.value)[0];
    let option;
    if (chartName === "关键词") {

        option = {
            series: [
                {
                    name: chartName,
                    type: 'wordCloud',
                    data: jsonData.data,
                    shape: 'circle',
                    gridSize: 1,
                    sizeRange: [12, 55],
                    rotationRange: [-90, 90],
                    drawOutOfBound: false,
                    textStyle: {
                        color: function () {
                            return 'yellow';
                        },
                        emphasis: {
                            shadowBlur: 10,
                            shadowColor: '#333'
                        }
                    }
                }
            ]
        };
    } else {
        option = {
            tooltip: {
                trigger: 'item'
            },
            series: [
                {
                    name: chartName,
                    type: 'pie',
                    data: jsonData.data,
                    radius: '50%',
                    label: {
                        color: '#ffffff'
                    }
                }
            ]
        };
    }
    if (jsonData.data) {
        chart.setOption(option);
    } else {
        chart.clear();
    }
    updateCharts.value = false;
}

watch(updateCharts, update)

watch(chartType, update)

onMounted(async () => {
    chart = echarts.init(variousChart.value);
    await update(chartType.value)
});

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

.groupItem {
    width: 100%;
    align-items: center;
    justify-content: center;
}

.chart {
    position: absolute;
    top: 25%;
    height: 75%;
    width: 100%;
}
</style>