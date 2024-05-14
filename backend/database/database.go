package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3" // 导入SQLite驱动
)

var db *sql.DB

func Open() {
	// 读取data.db中的数据，若没有，调用Init函数初始化数据库
	var err error
	db, err = sql.Open("sqlite3", "database/data.db")
	if err != nil {
		if err.Error() == "no such table: data" {
			fmt.Println("Database not found, initializing...")
			Init()
		} else {
			fmt.Println("Error opening database:", err)
			return
		}
	}
}

func Close() {
	if db != nil {
		db.Close()
	}
}

func PrintRecord(number int) {
	// 打印前number条记录
	rows, err := db.Query("SELECT * FROM data LIMIT ?", number)
	if err != nil {
		fmt.Println("Error querying database:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var uid int
		var colomns = make([]interface{}, 12)
		for i := range colomns {
			colomns[i] = new(string)
		}
		colomns = append([]interface{}{&uid}, colomns...)
		err = rows.Scan(colomns...)
		if err != nil {
			fmt.Println("Error scanning rows:", err)
			return
		}
		fmt.Print(int(uid), "    ")
		for i := 1; i < len(colomns); i++ {
			fmt.Print(*colomns[i].(*string), "    ")
		}
		fmt.Println()
	}
}
