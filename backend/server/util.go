package server

import (
	"database/sql"
	"encoding/json"
	"sort"
	"strings"
)

func get_keywords(rows *sql.Rows, limit int) (string, error) {
	// 将得到的keyword通过空格分割后统计，取出现次数最多的前20个
	keyword_map := make(map[string]int)
	for rows.Next() {
		var keyword string
		err := rows.Scan(&keyword)
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
	if len(response.Data) > limit {
		response.Data = response.Data[:limit]
	}

	// 将数据转换为json格式
	result, err := json.Marshal(response)
	if err != nil {
		return "", err
	}
	return string(result), nil
}
