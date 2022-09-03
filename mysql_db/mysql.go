package mysql_db

import "fmt"
import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func InitDB() (err error) {
	dsn := "root:@tcp(127.0.0.1:3306)/go_web?charset=utf8mb4&parseTime=True"
	// 也可以使用MustConnect连接不成功就panic
	Db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	Db.SetMaxOpenConns(20)
	Db.SetMaxIdleConns(10)
	return
}