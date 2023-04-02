package main

import (
	"database/sql"
	_ "localhost/sqlite3/gosqlite3"
	"log"
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
}
