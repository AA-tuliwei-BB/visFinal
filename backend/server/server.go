package server

// 监听本地的http请求
// 要监听的请求如下：
// |        Method        |  Path   |     Description      |
// | :------------------: | :-----: | :------------------: |
// | [POST](#post-filter) | /filter |     设置筛选条件     |
// |   [GET](#get-heat)   |  /heat  | 获取地图的热力值数据 |
// |   [GET](#get-list)   |  /list  |     获取列表数据     |
// |  [GET](#get-chart)   | /chart  |     获取图表数据     |
// |   [GET](#get-rel)    |  /rel   |    获取关系图数据    |

import (
	"fmt"
	"net/http"
	"strconv"
)

// 启动服务器
func Start() {
	http.HandleFunc("/filter", filter_handler)
	http.HandleFunc("/heat", heat_handler)
	http.HandleFunc("/list", list_handler)
	http.HandleFunc("/chart", chart_handler)
	http.HandleFunc("/rel", rel_handler)
	fmt.Println("Server started at http://localhost:8890")
	err := http.ListenAndServe(":8890", nil)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

// 设置筛选条件
func filter_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handling request: filter...")
	// 确保是post请求
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := set_filter(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// 返回成功
	w.WriteHeader(http.StatusOK)
	fmt.Println("filter set successfully:")
	fmt.Println(*get_filter())
}

// 获取地图的热力值数据
func heat_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handling request: heat...")
	// 确保是get请求
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// 返回get_heat(r)的结果
	result, err := get_heat()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte(result))
	fmt.Println("heat data sent successfully.")
}

// 获取列表数据
func list_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handling request: list...")
	// 确保是get请求
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// 从r的参数中解析page和size
	r.ParseForm()
	page, err := strconv.Atoi(r.Form.Get("page"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	size, err := strconv.Atoi(r.Form.Get("size"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := get_list(page, size)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte(result))
	fmt.Println("list data sent successfully.")
}

// 获取图表数据
func chart_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handling request: chart...")
}

// 获取关系图数据
func rel_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handling request: rel...")
}
