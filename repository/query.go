package repository

import (
	"database/sql"
	"fmt"
	"go-mariadb/config"
)

// Select the table of people
func Select() {
	db, err := sql.Open("mysql", "root:"+config.MariadbPassword+"@tcp(localhost:3306)/first")
	db.Ping()
	defer func() {
		if db != nil {
			db.Close()
		}
	}()
	if err != nil {
		fmt.Println("连接失败")
		return
	}

	stmt, err := db.Prepare("select * from people")
	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()
	if err != nil {
		fmt.Println("SQL语句预处理失败")
		return
	}

	rows, err := stmt.Query()
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	if err != nil {
		fmt.Println("获取SQL语句执行结果失败")
		return
	}

	for rows.Next() {
		var id int
		var name string
		var address string
		rows.Scan(&id, &name, &address)
		fmt.Println(id, name, address)
	}
}
