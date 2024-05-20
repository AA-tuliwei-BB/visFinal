<template>
    <el-table ref="tableRef" class="table" :data="exhibitData" highlight-current-row border stripe
        @current-change="updateDetails">
        <el-table-column v-for="item in columns" :prop="item.property" :label="item.label" :width="item.width"
            class="tableItem" />
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
            width: "140"
        }, {
            property: "ethnic",
            label: "民族",
            width: "60"
        }, {
            property: "category",
            label: "种类",
            width: "90"
        }, {
            property: "batch",
            label: "批次",
            width: "110"
        }, {
            property: "province",
            label: "省份",
            width: "60"
        }
    ];


const updateList = computed({
    get: () => store.state.updateList,
    set: (value) => store.commit('setUpdateList', value)
})

const tableRef = ref(null);
const pageSize = 8;
const totalSize = ref(0);
const currentPage = ref(1);
const activeNames = ref(null);

const fullData = ref(null);
const exhibitData = ref(null);
const nullData = {
    keyword: "",
    description: ""
}
const displayData = ref(nullData);

const update = async (val) => {
    fullData.value = await getList(currentPage.value, pageSize);
    totalSize.value = fullData.value.num;
    exhibitData.value = fullData.value.data;
    updateList.value = false;
}

const updateDetails = (val) => {
    displayData.value = val;
    if (!displayData.keyword) {
        displayData.keyword = "";
    }
    if (!displayData.description) {
        displayData.description = "";
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
    top: 15.5%;
    width: 49.5%;
    height: 75%;

}

.tableItem {
    font-size: 2vh;
}

.pagination {
    position: absolute;
    top: 90.5%;
    width: 50%;
    display: flex;
    justify-content: center;
}

.details {
    position: absolute;
    left: 50.5%;
    top: 15.5%;
    width: 48.5%;
    height: 83%;
}
</style>