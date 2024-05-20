<template>
    <div class="allFilterContainer">
        <div class="filterContainer">
            <span class="filterTitle">类别： </span>
            <el-select filterable class="filterItem" v-model="selectedCategories" multiple collapse-tags
                collapse-tags-tooltip placeholder="请选择">
                <el-option v-for="(item, index) in categories" :key="index" :label="item" :value="index">
                </el-option>
            </el-select>
        </div>

        <div class="filterContainer">
            <span class="filterTitle">批次： </span>
            <el-select filterable class="filterItem" v-model="selectedBatches" multiple collapse-tags
                collapse-tags-tooltip placeholder="请选择">
                <el-option v-for="(item, index) in declarationTimes" :key="index" :label="item" :value="index">
                </el-option>
            </el-select>
        </div>

        <div class="filterContainer">
            <span class="filterTitle">民族： </span>
            <el-select filterable class="filterItem" v-model="selectedEthnicity" multiple collapse-tags
                collapse-tags-tooltip placeholder="请选择">
                <el-option v-for="item in ethnicities" :key="item" :label="item" :value="item">
                </el-option>
            </el-select>
        </div>

        <div class="filterContainer">
            <span class="filterTitle">关键字： </span>
            <el-input v-model="selectedKeyword" class="inputItem" placeholder="请输入" clearable />
        </div>
        <el-button type="primary" class="submitButton" @click="submit">提交</el-button>
    </div>
</template>


<script setup>
import { computed } from 'vue'
import { useStore } from 'vuex'
import { postFilter } from '@/api/api';
import { categories, declarationTimes, ethnicities } from '@/utils/constants'

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

const updateCharts = computed({
    get: () => store.state.updateCharts,
    set: (value) => store.commit('setUpdateCharts', value)
})

const updateList = computed({
    get: () => store.state.updateList,
    set: (value) => store.commit('setUpdateList', value)
})

function submit() {
    postFilter(selectedCategories, selectedBatches, selectedEthnicity, selectedKeyword, selectedProvinces, updateHeatMap, updateCharts, updateList)
}

</script>

<style scoped>
.allFilterContainer {
    display: grid;
    grid-template-columns: repeat(2, 1fr);
    grid-template-rows: repeat(2, 1fr);
    align-items: center;
}

.filterContainer {
    display: flex;
    align-items: center;
}

.filterTitle {
    color: white;
    margin: 1.5vh 0.5vw 1.5vh 1vw;
    width: 15%;
    font-size: 0.9vw;
}

.inputItem {
    width: 70%;
}

.filterItem {
    width: 70%;
}

.submitButton {
    position: absolute;
    right: 4.3%;
    top: 15%;
}
</style>