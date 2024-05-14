# API
## 总览

|        Method        |  Path   |     Description      |
| :------------------: | :-----: | :------------------: |
| [POST](#post-filter) | /filter |     设置筛选条件     |
|   [GET](#get-heat)   |  /heat  | 获取地图的热力值数据 |
|   [GET](#get-list)   |  /list  |     获取列表数据     |
|  [GET](#get-chart)   | /chart  |     获取图表数据     |
|   [GET](#get-rel)    |  /rel   |    获取关系图数据    |

  
## POST /filter
### Description  
设置筛选条件，目前的想法是在服务器上维护一个session来保存筛选条件，
每次请求都会更新session中的筛选条件。
### Parameters  
None
### Data
```json  
{
    "category":
    [
        "category1",
        "category2",
        ...
    ],
    "batch":
    [
        "batch1",
        "batch2",
        ...
    ],
    "ethnic": "(ethnic filter string)", // 直接把输入拼进来
    "keyword": "(keyword filter string)", // 直接把输入拼进来
    "province":
    [
        "province1",
        "province2",
        ...
    ],
}
```
### example
`http://127.0.0.1/filter`
### Response  
None

## GET /heat
### Description
获取被选取的省份的地图热力值数据
### Parameters
None
### example
`http://127.0.0.1/heat`
### Response
```json
{
    "data":
    [
        {
            "province": "福建",
            "heat": 100
        },
        {
            "province": "河南",
            "heat": 200
        },
        ...
    ]
}
```

## GET /list
### Description
获取列表数据。  
注意，描述数据（用户点击的时候显示在右边的数据）会一起发过来，
其相关逻辑请在前端实现。
### Parameters
- `page`: 页数，如果超过最大页数了，则返回最后一页
- `size`: 每页大小
### example
`http://127.0.0.1/list?page=1&size=10`
### Response
```json
{
    "data":
    [
        {
            "id": 1,
            "name": "name1", // 名字
            "category": "category1", // 种类
            "batch": "batch1", // 批次
            "province": "province1", // 省份，没准会用到
            "description": "一大段描述... ",
        },
        ...
    ]
}
```

## GET /chart
### Description
获取图表数据。  
请注意，自由图表所用的数据都是kv结构
### Parameters
- `type`: 需要的数据种类，
    如`keyword`, `batch`, `category`, `ethnic`
### example
`http://127.0.0.1/chart?type=keyword`
### Response
不管是哪种type，返回都是是kv对的数组，如下
```json
{
    "data":
    [
        {
            "name": "keyword1",
            "value": 100
        },
        {
            "name": "keyword2",
            "value": 200
        },
        ...
    ]
}
```

## GET /rel
### Description
获取关系图数据。
### Parameters
- `type`: 地区还是关键词，`province`或`keyword`
- `name`: 地区或关键词的名字
- `limit`: 限制返回的数量
### example
`http://127.0.0.1/rel?type=province&name=福建&limit=10`
### Response
```json
{
    "data":
    [
        {
            "name": "keyword1",
            "value": 100
        },
        {
            "name": "keyword2",
            "value": 200
        },
        ...
    ]
}
```

# Todo List  

- [x] 确定API
- [x] 提取信息
    - [x] 提取关键词
    - [x] 提取省份
    - [x] 提取民族
- [ ] 建立数据库
- [ ] 完成各个API