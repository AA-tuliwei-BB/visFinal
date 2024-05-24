package server

import (
	"backend/database"
	"encoding/json"
	"sort"
)

func get_rel_from_province(name string, limit int) ([]pair, error) {
	// SELECT keyword FROM data WHERE -- get predicate -- AND province = ?;
	pred_str, args := get_predicate()
	args = append(args, name)

	sql := "SELECT keyword FROM data WHERE " + pred_str + " AND province = ?;"
	db := database.GetDB()
	rows, err := db.Query(sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	json_str, err := get_keywords(rows, limit)
	if err != nil {
		return nil, err
	}
	// json_str to []pair
	var response ChartResponse
	err = json.Unmarshal([]byte(json_str), &response)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

func get_provinces(limit int) ([]string, error) {
	// SELECT province
	// FROM data WHERE -- get predicate --;
	// GROUP BY province ORDER BY COUNT(*) DESC LIMIT ?;

	pred_str, args := get_predicate()
	sql := "SELECT province FROM data WHERE " + pred_str +
		" GROUP BY province ORDER BY COUNT(*) DESC LIMIT ?;"
	args = append(args, limit)

	db := database.GetDB()
	rows, err := db.Query(sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var provinces []string
	for rows.Next() {
		var province string
		err = rows.Scan(&province)
		if err != nil {
			return nil, err
		}
		provinces = append(provinces, province)
	}

	return provinces, nil
}

func get_keywords_and_links(provinces []string, limit int) (
	[]NodeObject, []LinkObject, error) {

	// 通过get_rel_from_province获取每个省份的关键词
	var keywords map[string]int = make(map[string]int)
	// link_map: a map of "(province, keyword) -> count"
	link_map := make(map[[2]string]int)

	for _, province := range provinces {
		keywords_, err := get_rel_from_province(province, limit)
		if err != nil {
			return nil, nil, err
		}
		for _, keyword := range keywords_ {
			keywords[keyword.Name] += keyword.Value
			link_map[[2]string{province, keyword.Name}] += keyword.Value
		}
	}

	// 从keywords转为slice并排序
	var keyword_list []pair
	for k, v := range keywords {
		keyword_list = append(keyword_list, pair{k, v})
	}
	sort.Slice(keyword_list, func(i, j int) bool {
		return keyword_list[i].Value > keyword_list[j].Value
	})
	if len(keyword_list) > limit {
		keyword_list = keyword_list[:limit]
	}

	// 将keywords转换为NodeObject
	var keyword_nodes []NodeObject
	for _, keyword := range keyword_list {
		keyword_nodes = append(keyword_nodes, NodeObject{keyword.Name})
	}

	// 通过provinces和keywords中的元素和link_map构造links
	var links []LinkObject
	for _, province := range provinces {
		for _, keyword := range keyword_list {
			if link_map[[2]string{province, keyword.Name}] > 0 {
				links = append(links, LinkObject{
					province, keyword.Name,
					link_map[[2]string{province, keyword.Name}],
				})
			}
		}
	}

	return keyword_nodes, links, nil
}

type LinkObject struct {
	Source string `json:"source"`
	Target string `json:"target"`
	Value  int    `json:"value"`
}

type NodeObject struct {
	Name string `json:"name"`
}

type RelResponse struct {
	Provinces []NodeObject `json:"provinces"`
	Keywords  []NodeObject `json:"keywords"`
	Links     []LinkObject `json:"links"`
}

func get_rel(numProvince int, numKeyword int) (string, error) {
	provinces, err := get_provinces(numProvince)
	if err != nil {
		return "", err
	}
	// get province nodes
	var province_nodes []NodeObject
	for _, province := range provinces {
		province_nodes = append(province_nodes, NodeObject{province})
	}

	keyword_nodes, links, err := get_keywords_and_links(provinces, numKeyword)
	if err != nil {
		return "", err
	}

	response := RelResponse{province_nodes, keyword_nodes, links}
	result, err := json.Marshal(response)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

// func get_rel_from_keyword(name string, limit int) (string, error) {
// 	// SELECT province FROM data WHERE -- get predicate -- AND keyword LIKE ?;
// 	pred_str, args := get_predicate()
// 	args = append(args, "%"+name+"%")
// 	args = append(args, limit)

// 	sql := "SELECT province, COUNT(*) as cnt FROM data " +
// 		"WHERE " + pred_str + " AND keyword LIKE ? " +
// 		"GROUP BY province ORDER BY cnt DESC LIMIT ?;"
// 	db := database.GetDB()
// 	rows, err := db.Query(sql, args...)
// 	if err != nil {
// 		fmt.Println(sql)
// 		return "", err
// 	}
// 	defer rows.Close()

// 	type RelResponse struct {
// 		Data []pair `json:"data"`
// 	}

// 	var response RelResponse
// 	for rows.Next() {
// 		var province string
// 		var count int
// 		err = rows.Scan(&province, &count)
// 		if err != nil {
// 			return "", err
// 		}
// 		response.Data = append(response.Data, pair{province, count})
// 	}

// 	result, err := json.Marshal(response)
// 	if err != nil {
// 		return "", err
// 	}
// 	return string(result), nil
// }
