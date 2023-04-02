package gosqlite3

import (
	"database/sql"
	"database/sql/driver"
	"log"
	"net/url"
)

type gosqlite3Driver struct{}

// Open implements driver.Driver
func (d *gosqlite3Driver) Open(name string) (driver.Conn, error) {
	conn := new(gosqlite3Connection)
	URL, err := url.Parse(name)
	if err != nil {
		return conn, err
	}
	log.Print(URL.Scheme)
	return conn, err
}

func init() {
	sql.Register("mysql", &gosqlite3Driver{})
}
