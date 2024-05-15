package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Filter struct {
	categories []string
	batch      []string
	ethnic     []string
	keywords   []string
	province   []string
}

type rowFilter struct {
	Category []string `json:"category"`
	Batch    []string `json:"batch"`
	Ethnic   string   `json:"ethnic"`
	Keyword  string   `json:"keyword"`
	Province []string `json:"province"`
}

var filter *Filter

func get_filter() *Filter {
	if filter == nil {
		filter = &Filter{
			categories: filter_all,
			batch:      filter_all,
			ethnic:     filter_all,
			keywords:   filter_all,
			province:   filter_all,
		}
	}
	return filter
}

func get_predicate() (string, []interface{}) {
	var query strings.Builder
	var args []interface{}
	filter := get_filter()

	strings_to_interface := func(strs []string) []interface{} {
		var result []interface{}
		for _, str := range strs {
			result = append(result, str)
		}
		return result
	}

	build_in_predicate := func(args_ []interface{}, name string) {
		if len(args_) > 0 {
			if (len(args_) == 1) && (args_[0] == "all") {
				query.WriteString("1=1 AND ")
				return
			}
			query.WriteString(name + " IN (")
			for i, arg := range args_ {
				if i > 0 {
					query.WriteString(", ")
				}
				query.WriteString("?")
				args = append(args, arg)
			}
			query.WriteString(") AND ")
		}
	}

	build_in_predicate(strings_to_interface(filter.categories), "category")
	build_in_predicate(strings_to_interface(filter.batch), "batch")
	build_in_predicate(strings_to_interface(filter.ethnic), "ethnic")
	build_in_predicate(strings_to_interface(filter.province), "province")

	// 构建keywords条件
	if len(filter.keywords) > 0 {
		if len(filter.keywords) == 1 && filter.keywords[0] == "all" {
			query.WriteString("1=1 AND ")
		} else {
			query.WriteString("(")
			for i, keyword := range filter.keywords {
				if i > 0 {
					query.WriteString(" OR ")
				}
				query.WriteString("name LIKE ?")
				args = append(args, "%"+keyword+"%")
			}
			query.WriteString(") AND ")
		}
	}

	// 移除最后一个 AND
	finalQuery := strings.TrimSuffix(query.String(), " AND ")
	fmt.Println("finalQuerySql: ", finalQuery)
	return finalQuery, args
}

func set_filter(r *http.Request) error {
	// 获取json数据
	decoder := json.NewDecoder(r.Body)
	var row rowFilter
	err := decoder.Decode(&row)
	if err != nil {
		return err
	}

	// 直接设置 categories, batch 和 province
	filter := get_filter()
	filter.categories = row.Category
	filter.batch = row.Batch
	filter.province = row.Province

	// ethnic 和 keyword 需要特殊处理：如果为空则设置为all，否则设置为以空格拆分的数组
	if row.Ethnic == "" {
		filter.ethnic = filter_all
	} else {
		filter.ethnic = strings.Split(row.Ethnic, " ")
	}
	if row.Keyword == "" {
		filter.keywords = filter_all
	} else {
		filter.keywords = strings.Split(row.Keyword, " ")
	}

	return nil
}
