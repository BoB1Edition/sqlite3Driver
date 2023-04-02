package gosqlite3

import "database/sql/driver"

type gosqlite3Connection struct{}

// Begin implements driver.Conn
func (*gosqlite3Connection) Begin() (driver.Tx, error) {
	panic("unimplemented")
}

// Close implements driver.Conn
func (*gosqlite3Connection) Close() error {
	panic("unimplemented")
}

// Prepare implements driver.Conn
func (*gosqlite3Connection) Prepare(query string) (driver.Stmt, error) {
	panic("unimplemented")
}
