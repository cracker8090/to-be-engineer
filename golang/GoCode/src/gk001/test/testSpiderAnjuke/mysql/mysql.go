package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var db *sql.DB

func init() {
	db, _ := sql.Open("mysql", "root:104104aa@tcp(47.116.8.77:3306)/game_news?charset=utf8")
	//db, _ := sql.Open("mysql", "root:104104aa@tcp(127.0.0.1:3306)/game?charset=utf8")
	db.SetMaxIdleConns(1000)
	err := db.Ping()
	if err != nil {
		fmt.Println("Failed to connect to mysql,err:" + err.Error())
		os.Exit(1)
	}

}

func DBCon() *sql.DB {
	return db
}
