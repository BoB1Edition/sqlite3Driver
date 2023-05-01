package main

import (
	"database/sql"
	"log"

	_ "github.com/BoB1Edition/sqlite3Driver/gosqlite3"
)

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	db, err := sql.Open("gosqlite3", "sqlite://test.db")
	checkError(err)
	defer db.Close()
	rows, err := db.Query("select * from testtable;")
	rows.Next()
	checkError(err)
	defer rows.Close()
}
