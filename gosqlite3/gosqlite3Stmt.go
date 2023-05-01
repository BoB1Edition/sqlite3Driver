package gosqlite3

import (
	"database/sql/driver"
)

type gosqlite3Stmt struct{}
type gosqlite3Rows struct{}

// Close implements driver.Rows
func (gosqlite3Rows) Close() error {
	panic("unimplemented Close")
}

// Columns implements driver.Rows
func (gosqlite3Rows) Columns() []string {
	panic("unimplemented Columns")
}

// Next implements driver.Rows
func (gosqlite3Rows) Next(dest []driver.Value) error {
	panic("unimplemented Next")
}

// Close implements driver.Stmt
func (stmt *gosqlite3Stmt) Close() error {
	panic("unimplemented Close")
}

// Exec implements driver.Stmt
func (stmt *gosqlite3Stmt) Exec(args []driver.Value) (driver.Result, error) {
	panic("unimplemented Exec")
}

// NumInput implements driver.Stmt
func (stmt *gosqlite3Stmt) NumInput() int {
	return -1
	//panic("unimplemented NumInput")
}

// Query implements driver.Stmt
func (stmt *gosqlite3Stmt) Query(args []driver.Value) (driver.Rows, error) {
	return gosqlite3Rows{}, nil
}
