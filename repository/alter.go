package repository

import (
	"database/sql"
	"fmt"
	"go-mariadb/config"
)

// Update table of people
func Update() {
	db, err := sql.Open("mysql", "root:"+config.MariadbPassword+"@tcp(localhost:3306)/first")
	db.Ping()
	defer func() {
		if db != nil {
			db.Close()
		}
	}()
	if err != nil {
		fmt.Println("数据库连接失败！")
		return
	}

	stmt, err := db.Prepare("UPDATE people SET name = ? , address = ? WHERE id = ?") // 其中 ? 是占位符
	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()
	if err != nil {
		fmt.Println("SQL语句预处理失败")
		return
	}

	result, err := stmt.Exec("李四", "朝阳", 3)
	if err != nil {
		fmt.Println("SQL语句执行失败")
		return
	}

	count, err := result.RowsAffected()
	if err != nil {
		fmt.Println("获取SQL执行结果失败")
		return
	}

	if count > 0 {
		fmt.Println("数据项修改成功")
	} else {
		fmt.Println("数据项修改失败")
	}
}
