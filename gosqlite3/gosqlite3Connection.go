package gosqlite3

import (
	"database/sql/driver"
	"log"
	"os"
)

type gosqlite3Connection struct {
	hFile *os.File
}

// Begin implements driver.Conn
func (conn *gosqlite3Connection) Begin() (driver.Tx, error) {
	panic("unimplemented Begin")
}

// Close implements driver.Conn
func (conn *gosqlite3Connection) Close() error {
	log.Print("connection driver.Conn")
	return conn.hFile.Close()
}

// Prepare implements driver.Conn
func (conn *gosqlite3Connection) Prepare(query string) (driver.Stmt, error) {
	stmt := new(gosqlite3Stmt)
	stmt.parseQuery(query)
	return stmt, nil
}
