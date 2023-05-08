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

const testquery0 string = `select testrow as t1, t2 "fa sd", t3, 
t4
, t5 as "ewrwerwe"
,t6,t7,t1.*
from testtable t1
left join testtable t2 on t1.id=t2.id 
full join testtable t2 on t2.id=t3.id 
where id = $1 or id = $1;
`

func main() {
	db, err := sql.Open("gosqlite3", "sqlite://test1.db")
	checkError(err)
	defer db.Close()
	//rows, err := db.Query("select testrow from testtable where id = $1 or id = $1;", 0)
	rows, err := db.Query(testquery0, 0)
	checkError(err)
	defer rows.Close()
	rows.Next()
	var testrow int
	rows.Scan(&testrow)
	log.Print(testrow)
	res, err := db.Exec("insert into testtable values(3,2);")
	checkError(err)
	ra, err := res.RowsAffected()
	checkError(err)
	log.Print(ra)
}
