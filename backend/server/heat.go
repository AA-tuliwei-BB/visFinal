package server

import (
	"backend/database"
	"encoding/json"
)

// Response
// {
//     "data":
//     [
//         {
//             "province": "福建",
//             "heat": 100
//         },
//         {
//             "province": "河南",
//             "heat": 200
//         },
//         ...
//     ]
// }

type HeatResponse struct {
	Data []struct {
		Province string `json:"province"`
		Heat     int    `json:"heat"`
	} `json:"data"`
}

func get_heat() (string, error) {
	sql := "SELECT province, COUNT(*) AS heat FROM data"
	predicate, args := get_predicate()
	sql += " WHERE " + predicate
	sql += " GROUP BY province"

	db := database.GetDB()
	rows, err := db.Query(sql, args...)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	// 查询数据库，获取response
	var response HeatResponse
	for rows.Next() {
		var province string
		var heat int
		err = rows.Scan(&province, &heat)
		if err != nil {
			return "", err
		}
		response.Data = append(response.Data, struct {
			Province string `json:"province"`
			Heat     int    `json:"heat"`
		}{Province: province, Heat: heat})
	}

	// 将response转换为json格式
	result, err := json.Marshal(response)
	if err != nil {
		return "", err
	}
	return string(result), nil
}
