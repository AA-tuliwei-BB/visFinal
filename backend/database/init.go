package database

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3" // 导入SQLite驱动
)

// Init函数: 从preprocessed.csv 中读取数据并初始化sqlite数据库
func Init() {
	// 从 "./preprocessed.csv" 中读取数据
	file, err := os.Open("database/preprocessed.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	// 获取表头和其他数据，存放在header和records中
	reader.Read()
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// 初始化数据库
	db, err := sql.Open("sqlite3", "database/data.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	// 将file中的数据导入到数据库
	// 建表, 表的列：uid ,项目序号,编号,名称,类别,公布时间,类型,申报地区或单位,保护单位,省份,民族,关键词,描述
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS data
                    (uid INTEGER PRIMARY KEY, -- uid
                    project_number TEXT, -- 项目序号
                    number TEXT, -- 编号
                    name TEXT, -- 名称
                    category TEXT, -- 类别
                    publish_time TEXT, -- 公布时间
                    type TEXT, -- 类型
                    declare_area_or_unit TEXT, -- 申报地区或单位
                    protection_unit TEXT, -- 保护单位
                    province TEXT, -- 省份
                    nation TEXT, -- 民族
                    keyword TEXT, -- 关键词
                    description TEXT -- 描述
					)`)
	if err != nil {
		fmt.Println("Error creating table:", err)
		return
	}

	// 插入数据
	stmt, err := db.Prepare(`INSERT INTO data VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		fmt.Println("Error preparing statement:", err)
		return
	}
	defer stmt.Close()

	for _, record := range records {
		// 将record中的数据插入到数据库
		parsedRecord := make([]interface{}, len(record))
		for i, v := range record {
			if i == 0 {
				// 将第一列转换为int
				parsedRecord[i], _ = strconv.Atoi(v)
			} else {
				parsedRecord[i] = v
			}
		}

		_, err = stmt.Exec(parsedRecord...)
		if err != nil {
			fmt.Println("Error inserting record:", err)
			return
		}
	}
}
