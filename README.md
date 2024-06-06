# API
[api](backend/api.md)

# 前端文档
[frontend](frontend/doc_frontend.md)

# 部署
## backend 部署方法
### 依赖
- go 1.21
- cgo (即需要有gcc环境)
### 运行
```shell
go run backend/main.go
# 注意要开放8890端口
```

## frontend 部署方法

### 处理依赖

```shell
npm install
```

### 动态运行

```shell
npm run dev
```

会将项目部署在`localhost:5173`。

# 开发工具
## 数据处理
- 语言: Python
- jieba: 提取关键词用
## 后端
- 语言: Golang
- SQLite3: 数据库
- net/http: http请求
- cors: 跨域请求
## 前端
- vue3: 前端框架
- vuex: 全局变量
- element-plus: 前端UI
- echarts: 绘制图表用
- axios: http请求

# 实现细节
## 数据预处理
对数据进行预处理为合适格式的csv。
使用python，核心代码为[preprocess.py](data/preprocess.py) 和
[util.py](data/util.py)。
### 简单处理
添加UID。  
从“申报地区或单位”中获取“省份”属性。  
从描述中找出出现次数最多的民族名称作为“民族”属性。  
提取关键词并添加为“关键词”属性。  
将描述里面的英文逗号全部改为中文逗号，以免与csv文件的逗号冲突。

### 关键词提取
使用jieba库的extract_tags, Textrank, 和分词功能提取关键词，
再经过词性筛选，禁用词列表筛选等方式得到最终获取的关键词。

## 后端

### 启动流程
先尝试开启数据库，再开启http服务器开始监听前端请求。
### 数据库
使用sqlite3作为数据库。  
在程序启动的时候会尝试开启数据库。如果已经存在数据库则打开，
否则使用数据预处理获得的preprocessed.csv文件建立新的数据库。

### 服务器主体
开启对前端请求的监听，当收到请求的时候，进行初步处理，
然后调用相关的模块得到结果返回给前端。

### filter模块
在内存中保存一个filter，当前端发送post请求的时候修改这个filter。
提供get_predicate()方法，当其他模块需要查询数据库时可通过get_predicate方法直接
得到当前filter对应的谓语，以拼接到sql中。

### heat模块
通过对数据库进行sql查询得到结果，序列化为json返回给前端。
sql:
``` sql
SELECT province, COUNT(*) AS heat FROM data
WHERE -- get_predicate
GROUP BY province
```
核心代码:
``` go
sql := "SELECT province, COUNT(*) AS heat FROM data"
predicate, args := get_predicate()
sql += " WHERE " + predicate
sql += " GROUP BY province"

db := database.GetDB()
rows, err := db.Query(sql, args...)
```

### list模块
通过对数据库进行sql查询得到结果，序列化为json返回给前端。  
sql:
``` sql
SELECT MIN(data.uid) AS uid
    data.name AS name,
    data.category AS category,
    data.batch AS batch,
    data.ethnic AS ethnic,
    data.keyword AS keyword,
    GROUP_CONCAT(DISTINCT data.province) AS province,
    replace(GROUP_CONCAT(DISTINCT rtrim(
        replace(data.description, '（分段）', char(10)), char(10) )),
    ',', '\nMore: \n') AS description
FROM data
INNER JOIN (
    SELECT project_number AS pid, COUNT(project_number) AS count 
    FROM data 
    WHERE -- get predicate
    GROUP BY pid 
) AS t1 ON data.project_number = t1.pid 
WHERE -- get predicate
GROUP BY data.name, data.batch 
ORDER BY count DESC 
LIMIT ? OFFSET ?;
```

### chart模块
keywork 全部读取后手动统计  
其他直接进行sql查询:
``` sql
SELECT ?, COUNT(*) AS cnt FROM data
GROUP BY ?  ORDER BY cnt DESC; -- ? 为字段名
```

### rel模块
1. 获取需要的省份：
    ``` sql
    SELECT province FROM data WHERE -- get predicate
    GROUP BY province ORDER BY COUNT(*) DESC LIMIT ?; -- ? 为省份数量上限
    ```
2. 对每个省份，通过和chart模块相似的方法获取一些关键词。  
3. 将第二步获取的关键词合并并排序，得到最终决定选取的关键词
4. 将省份，关键词，和省份与关键词直接的关联的权重这三组信息序列化为json返回给前端。