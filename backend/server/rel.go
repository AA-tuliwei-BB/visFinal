package server

import (
	"backend/database"
	"encoding/json"
	"fmt"
)

func get_rel_from_province(name string, limit int) (string, error) {
	// SELECT keyword FROM data WHERE -- get predicate -- AND province = ?;
	pred_str, args := get_predicate()
	args = append(args, name)

	sql := "SELECT keyword FROM data WHERE " + pred_str + " AND province = ?;"
	db := database.GetDB()
	rows, err := db.Query(sql, args...)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	return get_keywords(rows, limit)
}

func get_rel_from_keyword(name string, limit int) (string, error) {
	// SELECT province FROM data WHERE -- get predicate -- AND keyword LIKE ?;
	pred_str, args := get_predicate()
	args = append(args, "%"+name+"%")
	args = append(args, limit)

	sql := "SELECT province, COUNT(*) as cnt FROM data " +
		"WHERE " + pred_str + " AND keyword LIKE ? " +
		"GROUP BY province ORDER BY cnt DESC LIMIT ?;"
	db := database.GetDB()
	rows, err := db.Query(sql, args...)
	if err != nil {
		fmt.Println(sql)
		return "", err
	}
	defer rows.Close()

	type RelResponse struct {
		Data []pair `json:"data"`
	}

	var response RelResponse
	for rows.Next() {
		var province string
		var count int
		err = rows.Scan(&province, &count)
		if err != nil {
			return "", err
		}
		response.Data = append(response.Data, pair{province, count})
	}

	result, err := json.Marshal(response)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

func get_rel(typename string, name string, limit int) (string, error) {
	if typename == "province" {
		return get_rel_from_province(name, limit)
	} else if typename == "keyword" {
		return get_rel_from_keyword(name, limit)
	} else {
		return "", fmt.Errorf("invalid typename")
	}
}
