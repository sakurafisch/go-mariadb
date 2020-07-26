package repository

import (
	"database/sql"
	"fmt"
	"go-mariadb/config"
)

// Delete from the table of people
func Delete() {
	db, err := sql.Open("mysql", "root:"+config.MariadbPassword+"@tcp(localhost:3306)/first")
	defer func() {
		if db != nil {
			db.Close()
		}
	}()
	if err != nil {
		fmt.Println("数据库连接失败")
		return
	}

	stmt, err := db.Prepare("DELETE FROM people WHERE id = ?")
	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()
	if err != nil {
		fmt.Println("SQL语句预处理失败")
		return
	}

	result, err := stmt.Exec(4)
	if err != nil {
		fmt.Println("SQL语句执行失败")
	}

	count, err := result.RowsAffected()
	if err != nil {
		fmt.Println("获取SQL语句执行结果失败")
	}
	if count > 0 {
		fmt.Println("删除数据项成功")
	} else {
		fmt.Println("删除数据项失败")
	}

}
