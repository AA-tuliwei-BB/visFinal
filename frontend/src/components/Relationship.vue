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
        return;
    }
    let jsonData = await getRel("province", currentProvince);
    let nodes = [{ name: currentProvince }]
    let links = []
    jsonData.data.forEach(element => {

        if (nodes.find(node => node.name === element.name)) {
            return;
        }
        nodes.push({ name: element.name });
        links.push({ source: currentProvince, target: element.name, value: element.value });
    });
    console.log(nodes, links)
    const option = {
        tooltip: {
            trigger: 'item',
            triggerOn: 'mousemove'
        },
        series: {
            type: 'sankey',
            layout: 'none',
            data: nodes,
            links: links,
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
    top: 25%;
    height: 75%;
    width: 100%;
}
</style>