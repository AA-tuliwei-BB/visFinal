<template>
    <div class="bgpage">
        <Header></Header>
        <!-- 用两个box作为示例，具体使用请自行更改 -->
        <div>
            <div class="container">
                <div class="box">
                    <Viewbox title="筛选器" :boxb="true">
                        <Filter />
                    </Viewbox>
                </div>
                <div class="box">
                    <Viewbox :title="subTitle" :boxb="true" />
                </div>
                <div class="box">
                    <Viewbox :title="subTitle" :boxb="true" />
                </div>
            </div>
            <div class="container">
                <div class="box">
                    <Viewbox title="中国地图" :boxb="true">
                        <Heatmap />
                    </Viewbox>
                </div>
                <div class="box">
                    <Viewbox title="测试子标题" :boxb="true" />
                </div>
            </div>

        </div>
    </div>
</template>

<script setup>
import { ref } from 'vue';
import Header from './Header.vue'
import Filter from './filter.vue';
import Viewbox from './viewbox/Viewbox.vue'
import Heatmap from './Heatmap.vue'
const subTitle = ref("子标题");
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
    height: 42vh;
    margin: 0.5vw;
}
</style>