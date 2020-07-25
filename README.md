# Go 连接 MariaDB

## 如何运行

1. 先建表
2. 重命名./config/config.go.example，并在该文件中写上正确的配置。
3. 执行 go run main.go

## 开发记录

### 依赖管理

```bash
go mod init go-mariadb
go get -u -v https://github.com/go-sql-driver/mysql
```

### 建表 SQL

```sql
CREATE TABLE people (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name varchar(20),
    address varchar(100)
)
```
