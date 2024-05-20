<template>
    <el-table ref="tableRef" row-class-name="tableItem" header-cell-class-name="headerRow" class="table"
        :data="exhibitData" highlight-current-row border show-overflow-tooltip @current-change="updateDetails">
        <el-table-column v-for="item in columns" :prop="item.property" :label="item.label" :width="item.width"
            header-align="center" align="center" />
    </el-table>
    <div class="pagination">
        <el-pagination layout="prev, pager, next" :total="totalSize" :sizes="pageSize" :pager-count="9"
            v-model:current-page="currentPage" @current-change="update" />
    </div>
    <el-scrollbar class="details">
        <el-collapse v-model="activeNames">
            <el-collapse-item title="关键词" name="1">
                {{ displayData.keyword }}
            </el-collapse-item>
            <el-collapse-item title="描述" name="2">
                {{ displayData.description }}
            </el-collapse-item>
        </el-collapse>
    </el-scrollbar>
</template>

<script setup>
import { onMounted, ref } from 'vue';
import { getList } from '@/api/api';
import { useStore } from 'vuex'
import { computed, watch } from 'vue'

const store = useStore()

const columns =
    [
        {
            property: "name",
            label: "名称",
            width: "176"
        }, {
            property: "ethnic",
            label: "民族",
            width: "68"
        }, {
            property: "category",
            label: "种类",
            width: "82"
        }, {
            property: "batch",
            label: "批次",
            width: "68"
        }, {
            property: "province",
            label: "省份",
            width: "68"
        }
    ];


const updateList = computed({
    get: () => store.state.updateList,
    set: (value) => store.commit('setUpdateList', value)
})

const tableRef = ref(null);
const pageSize = 7;
const totalSize = ref(0);
const currentPage = ref(1);
const activeNames = ref(null);

const fullData = ref(null);
const exhibitData = ref(null);
const nullData = {
    keyword: " ",
    description: " "
}
const displayData = ref(nullData);

const update = async (val) => {
    fullData.value = await getList(currentPage.value, pageSize);
    totalSize.value = fullData.value.num;
    exhibitData.value = fullData.value.data;
    updateList.value = false;
}

const updateDetails = (val) => {
    if (val) {
        displayData.value = val;
        if (!displayData.value.keyword) {
            displayData.value.keyword = " ";
        }
        if (!displayData.value.description) {
            displayData.value.description = " ";
        }
    } else {
        displayData.value = nullData;
    }
}


watch(updateList, () => {
    currentPage.value = 0;
    update();
})


onMounted(async () => {
    await update()
});

</script>

<style scoped>
.table {
    position: absolute;
    top: 15%;
    width: 49.5%;
    height: 76%;
    left: 0.5%;
}

/* :deep() .headerRow {
    background-color: aqua !important;
} */

/* :deep() .tableItem {
    background-color: aqua !important;
} */

.pagination {
    position: absolute;
    top: 91.5%;
    width: 50%;
    display: flex;
    justify-content: center;
}

.details {
    position: absolute;
    left: 50.5%;
    top: 15%;
    width: 48.5%;
    height: 84%;
}
</style>