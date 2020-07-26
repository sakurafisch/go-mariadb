package main

import (
	"go-mariadb/repository"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	repository.Insert()
	repository.Select()
}
