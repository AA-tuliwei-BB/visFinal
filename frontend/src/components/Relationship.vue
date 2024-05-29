<template>
    <div ref="relationChart" class="chart"></div>
</template>

<script setup>
import { onMounted, ref } from 'vue';
import * as echarts from 'echarts/core';
import { SankeyChart } from 'echarts/charts';
import { TooltipComponent, LegendComponent } from 'echarts/components';
import { CanvasRenderer } from 'echarts/renderers';
import { getRel } from '@/api/api';
import { useStore } from 'vuex'
import { computed, watch } from 'vue'

const store = useStore()

const updateRelationship = computed({
    get: () => store.state.updateRelationship,
    set: (value) => store.commit('setUpdateRelationship', value)
})
const selectedProvinces = computed({
    get: () => store.state.selectedProvinces,
    set: (value) => store.commit('selectedProvinces', value)
})

echarts.use([TooltipComponent, SankeyChart, LegendComponent, CanvasRenderer]);

const relationChart = ref(null);
let relChart;

const update = async () => {
    let currentProvince = selectedProvinces.value[0];
    console.log(currentProvince)
    if (!currentProvince || currentProvince === "all") {
        relChart.clear();
        updateRelationship.value = false;
        return;
    }
    let jsonData = await getRel("province", currentProvince);
    let nodes = [{ name: currentProvince }]
    let links = []
    console.log(jsonData)
    if (jsonData.data === null) {
        relChart.clear();
        updateRelationship.value = false;
        return;
    }
    // console.log([jsonData.provinces, jsonData.keywords])
    const option = {
        tooltip: {
            trigger: 'item',
            triggerOn: 'mousemove'
        },
        series: {
            type: 'sankey',
            layout: 'none',
            data: [...jsonData.provinces, ...jsonData.keywords],
            links: jsonData.links,
            emphasis: {
                focus: 'adjacency'
            },
            lineStyle: {
                color: 'source',
                curveness: 0.5
            }
        }
    };
    relChart.setOption(option);
    updateRelationship.value = false;
}

watch(updateRelationship, update)

onMounted(async () => {
    relChart = echarts.init(relationChart.value);
    await update()
});

</script>

<style scoped>
.chart {
    position: absolute;
    top: 8vh;
    height: calc(98% - 8vh);
    width: 100%;
}
</style>