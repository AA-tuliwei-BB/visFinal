# 前端开发文档

## 项目结构

- 主框架：[vue3](https://cn.vuejs.org/)
- 组件库：[element-plus](element-plus.org)
- api访问：axios
- 全局变量使用：vuex

## vue3

主要使用组合式api，可以观察`.vue`文件的结构。`<script setup>`中写数据和方法，`<template>`中的`{{value}}`可字符串插入数据。

还有`v-bind, v-model, v-for, v-if, v-on(@)`等标签，可以看文档熟悉一下或者直接问llm。

## element-plus

与vue3组合使用的组件库。使用方法是看[组件文档](https://element-plus.org/zh-CN/component/overview.html)，查看示例，复制源码改改就完了。

## axios

axios库通过XML http request与后端交互。

在`src/utils/request.js`封装了一层interceptor处理错误，使用可以直接`request.get("/path", params)`。

获取数据时，获取数据的函数应该做成异步的，在获取数据处同步阻塞：

```js
async function func() {
    await response = getXXX();
}
```

与后端交互的函数我是放在`api.js`中。

## vuex

简单来说就是全局变量。全局变量的定义在`src/store.js`中。当前定义的全局变量有：

```js
state: {
    selectedCategories: [],
    selectedEthnicity: [],
    selectedBatches: "",
    selectedKeyword: "",
    selectedProvinces: [],
  }
```

都是字面意思，代表选中的类别、民族、申报批次、……。之后计划再加几个变量`updateXXX`来传递update消息。

这玩意只能在`<script setup>`中使用，使用方法：

```js
import { useStore } from 'vuex'
import { computed } from 'vue'

const store = useStore()

const selectedCategories = computed({
    get: () => store.state.selectedCategories,
    set: (value) => store.commit('setSelectedCategories', value)
})
```

之后`selectedCategories`可以正常地与template中的变量进行绑定，通过`selectedCategories.value`可以访问它的值。
