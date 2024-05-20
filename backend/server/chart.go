package server

import (
	"backend/database"
	"encoding/json"
	"fmt"
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

	return get_keywords(rows, 100)
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
