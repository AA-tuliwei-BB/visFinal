<template>
    <div class="bgpage">
        <Header></Header>
        <!-- 用两个box作为示例，具体使用请自行更改 -->
        <div>
            <div class="container">
                <div class="filterBox">
                    <Viewbox title="筛选器" :boxb="true">
                        <Filter />
                    </Viewbox>
                </div>
                <div class="listBox">
                    <Viewbox title="项目列表" :boxb="true" />
                </div>
                <div class="chartBox">
                    <Viewbox title="图表展示" :boxb="true">
                        <Charts />
                    </Viewbox>
                </div>
            </div>
            <div class="container">
                <div class="mapBox">
                    <Viewbox title="中国地图" :boxb="true">
                        <Heatmap />
                    </Viewbox>
                </div>
                <div class="relationshipBox">
                    <Viewbox title="关系图" :boxb="true" />
                </div>
            </div>

        </div>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import Header from './Header.vue'
import Filter from './filter.vue';
import Viewbox from './viewbox/Viewbox.vue'
import Heatmap from './Heatmap.vue'
import request from '@/utils/request.js'
import Charts from './Charts.vue'
const subTitle = ref("子标题");

onMounted(async () => {
    const initJsonString = '{"category":["all"],"batch":["all"],"ethnic":"","keyword":"","province":["all"]}';
    await request.post("/filter", initJsonString)
});

</script>


<style>
.bgpage {
    background: url(src/assets/true.png);
    height: 100vh;
    width: 100vw;
}

.container {
    display: flex;
    /* 将容器设置为 flex 容器 */
}

.box {
    flex: 1;
    /* 每个子元素占据相等的空间 */
    height: 44vh;
    margin: 0.5vw;
}

.mapBox {
    position: absolute;
    left: 0.5%;
    top: 30.5%;
    width: 49%;
    height: 69%;
}

.filterBox {
    position: absolute;
    left: 0.5%;
    top: 10.5%;
    width: 49%;
    height: 19%;
}

.relationshipBox {
    position: absolute;
    left: 50.5%;
    top: 10.5%;
    width: 29%;
    height: 44%;
}

.chartBox {
    position: absolute;
    left: 80.5%;
    top: 10.5%;
    width: 19%;
    height: 44%;
}

.listBox {
    position: absolute;
    left: 50.5%;
    top: 55.5%;
    width: 49%;
    height: 44%;
}
</style>