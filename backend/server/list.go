package server

import (
	"backend/database"
	"encoding/json"
	"fmt"
)

// SELECT MIN(data.uid) AS uid,
//        data.name AS name,
//        data.category AS category,
//        data.batch AS batch,
//        GROUP_CONCAT(DISTINCT data.province, ', ') AS province,
//        GROUP_CONCAT(DISTINCT data.description, char(10) ||
// 				'And(Exception): ') AS description,
// FROM data
// INNER JOIN (
//     SELECT project_number AS pid, COUNT(project_number) AS count
//     FROM data
//     GROUP BY pid
// ) AS t1 ON data.project_number = t1.pid
// WHERE -- get filter --
// GROUP BY data.name, data.batch
// ORDER BY count DESC
// LIMIT ? OFFSET ?;

type ListResponse struct {
	Data []struct {
		Uid         int    `json:"id"`
		Name        string `json:"name"`
		Batch       string `json:"batch"`
		Category    string `json:"category"`
		Province    string `json:"province"`
		Description string `json:"description"`
	} `json:"data"`
}

func get_list(page int, size int) (string, error) {
	pred_str, args := get_predicate()
	sql := "SELECT MIN(data.uid) AS uid, " +
		"data.name AS name, " +
		"data.category AS category, " +
		"data.batch AS batch, " +
		"GROUP_CONCAT(DISTINCT data.province) AS province, " +
		// rtrim(replace(group_concat(DISTINCT column1||'@!'), '@!,', '|'),'@!')
		"replace(GROUP_CONCAT(DISTINCT " +
		"    rtrim(replace(data.description, '（分段）', char(10)), char(10)) )," +
		"    ',', '\nMore: \n') AS description " +
		"FROM data " +
		"INNER JOIN ( " +
		"    SELECT project_number AS pid, COUNT(project_number) AS count " +
		"    FROM data " +
		"    GROUP BY pid " +
		") AS t1 ON data.project_number = t1.pid " +
		"WHERE " + pred_str + " " +
		"GROUP BY data.name, data.batch " +
		"ORDER BY count DESC " +
		"LIMIT ? OFFSET ?;"

	args = append(args, size, (page-1)*size)

	db := database.GetDB()
	rows, err := db.Query(sql, args...)
	if err != nil {
		fmt.Println("test", err)
		return "", err
	}
	defer rows.Close()

	// 查询数据库，获取response
	var response ListResponse
	for rows.Next() {
		var uid int
		var name string
		var category string
		var batch string
		var province string
		var description string
		err = rows.Scan(&uid, &name, &category, &batch, &province, &description)
		if err != nil {
			return "", err
		}
		response.Data = append(response.Data, struct {
			Uid         int    `json:"id"`
			Name        string `json:"name"`
			Batch       string `json:"batch"`
			Category    string `json:"category"`
			Province    string `json:"province"`
			Description string `json:"description"`
		}{Uid: uid, Name: name, Category: category, Batch: batch, Province: province, Description: description})
	}

	// 将response转换为json格式
	result, err := json.Marshal(response)
	return string(result), err
}
