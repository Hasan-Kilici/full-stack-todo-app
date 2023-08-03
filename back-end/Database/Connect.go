package Database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"todoList/Utils"
)

var connStr string = Utils.GetConfig("postgre", "connect")
var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err.Error())
	}
}

func closeDB() {
	if db != nil {
		db.Close()
	}
}

func Connect() {
	db, err := sql.Open("postgres", connStr)
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()
}