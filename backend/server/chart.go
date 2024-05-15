package server

import (
	"backend/database"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

// keywork 全部读取后手动统计

// 其他：
// SELECT ?, COUNT(*) as count FROM data GROUP BY ?; -- ? 为字段名

type pair struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type ChartResponse struct {
	Data []pair `json:"data"`
}

func get_chart_keyword() (string, error) {
	// SELECT keyword FROM data WHERE -- get predicate --;
	pred_str, args := get_predicate()
	sql := "SELECT keyword FROM data WHERE " + pred_str + ";"
	db := database.GetDB()
	rows, err := db.Query(sql, args...)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	// 将得到的keyword通过空格分割后统计，取出现次数最多的前20个
	keyword_map := make(map[string]int)
	for rows.Next() {
		var keyword string
		err = rows.Scan(&keyword)
		if err != nil {
			return "", err
		}
		words := strings.Fields(keyword)
		for _, word := range words {
			keyword_map[word]++
		}
	}

	// 将map转换为slice
	var response ChartResponse
	for k, v := range keyword_map {
		response.Data = append(response.Data, pair{k, v})
	}
	// 排序
	sort.Slice(response.Data, func(i, j int) bool {
		return response.Data[i].Value > response.Data[j].Value
	})
	// 取前20个
	if len(response.Data) > 20 {
		response.Data = response.Data[:20]
	}

	// 将数据转换为json格式
	result, err := json.Marshal(response)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

func get_chart(typename string) (string, error) {
	// keyword 较为特殊，需要单独处理
	if typename == "keyword" {
		return get_chart_keyword()
	} else if typename != "category" &&
		typename != "batch" && typename != "ethnic" {
		return "", fmt.Errorf("invalid typename")
	}

	pred_str, args := get_predicate()

	// 从数据库中获取图表数据
	sql_stmt := "SELECT " + typename + ", COUNT(*) as cnt FROM data " +
		"WHERE " + pred_str + " GROUP BY " + typename + " " +
		"ORDER BY cnt DESC;"
	db := database.GetDB()
	rows, err := db.Query(sql_stmt, args...)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	// 将数据存入response
	var response ChartResponse
	for rows.Next() {
		var name string
		var value int
		err = rows.Scan(&name, &value)
		if err != nil {
			return "", err
		}
		response.Data = append(response.Data, pair{name, value})
	}

	// 将数据转换为json格式
	result, err := json.Marshal(response)
	if err != nil {
		return "", err
	}
	return string(result), nil
}
