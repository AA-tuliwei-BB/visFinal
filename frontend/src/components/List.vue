<template>
    <el-table ref="tableRef" row-class-name="tableItem" header-cell-class-name="headerRow" class="table"
        :data="exhibitData" highlight-current-row show-overflow-tooltip @current-change="updateDetails"
        :row-style="{ backgroundColor: '#f5f5f5' }">
        <el-table-column v-for="item in columns" :prop="item.property" :label="item.label" :width="item.width"
            header-align="center" align="center" />
    </el-table>
    <div class="pagination">
        <el-pagination small layout="prev, pager, next" :total="totalSize" :sizes="pageSize" :pager-count="9"
            v-model:current-page="currentPage" @current-change="update" />
    </div>
    <el-scrollbar class="details">
        <el-collapse v-model="activeNames">
            <el-collapse-item title="关键词" name="1">
                {{ displayData.keyword }}
            </el-collapse-item>
            <el-collapse-item title="描述" name="2">
                <div v-html="displayData.description"></div>
                <!-- {{ displayData.description }} -->
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
const activeNames = ref(['1', '2']);

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
    activeNames.value = ['1', '2']
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
        activeNames.value = ['1', '2']
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
    background-color: #00000020;
    position: absolute;
    top: 15%;
    /*max-width: 49.5%;*/
    width: 462px;
    height: 320px;
    left: 2%;
    /* 内容居中 */
    display: flex;
    justify-content: center;
}

.el-table {
    --el-table-tr-bg-color: transparent;
    --el-table-header-bg-color: rgba(255, 255, 255, 0.1);

    --el-table-border-color: rgba(255, 255, 255, 0.2);

    --el-table-row-hover-bg-color: rgba(255, 255, 255, 0.1);
    --el-table-current-row-bg-color: rgba(255, 255, 255, 0.15);
}

:deep() .tableItem {
    background-color: rgba(131, 168, 236, 0.1) !important;
    color: rgba(255, 255, 255, 0.5);
}

.el-collapse {
    --el-collapse-header-bg-color: rgba(255, 255, 255, 0.1);
    /* 修改为你想要的颜色 */
    --el-collapse-border-color: rgba(255, 255, 255, 0.2);
    --el-collapse-content-bg-color: rgba(255, 255, 255, 0.15);
    --el-collapse-header-text-color: rgba(255, 255, 255, 0.5);
    --el-collapse-content-text-color: rgba(255, 255, 255, 0.5);
}

.el-pagination {
    --el-pagination-bg-color: transparent;
    --el-pagination-button-color: gray;
    --el-pagination-button-disabled-bg-color: transparent;
}

.pagination {
    position: absolute;
    top: min(93%, 420px);
    left: 20px;
    width: 462px;
    display: flex;
    justify-content: center;
}

.details {
    position: absolute;
    left: 500px;
    width: calc(98% - 500px);
    top: 15%;
    height: min(320px, 80%);
}
</style>