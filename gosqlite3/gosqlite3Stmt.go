package gosqlite3

import (
	"database/sql/driver"
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	selectrowregex string = `(?ism)select ("?.+"?[^,\s]).*from.*`
)

type gosqlite3Stmt struct {
	query  string
	inputs int
}

type gosqlite3Rows struct {
	columns []string
	mu      sync.Mutex
}

// Close implements driver.Rows
func (gosqlite3Rows) Close() error {
	//panic("unimplemented Close")
	return nil
}

func (driversRows *gosqlite3Rows) runQuery(query string) {
	driversRows.mu.Lock()
	defer driversRows.mu.Unlock()
	re := regexp.MustCompile(selectrowregex)
	if !re.MatchString(query) {
		driversRows.columns = nil
		return
	}
	rowsstring := re.FindStringSubmatch(query)
	rowslist := strings.Split(rowsstring[1], ",")
	driversRows.columns = make([]string, 0)
	for i := range rowslist {
		rowslist[i] = strings.TrimSpace(rowslist[i])
		rowslist[i] = strings.Trim(rowslist[i], "\n")
		rowslist[i] = strings.Trim(rowslist[i], ",")
	}
}

// Columns implements driver.Rows
func (driversRows *gosqlite3Rows) Columns() []string {
	driversRows.mu.Lock()
	defer driversRows.mu.Unlock()
	//log.Print(driversRows)
	return driversRows.columns
}

// Next implements driver.Rows
func (gosqlite3Rows) Next(dest []driver.Value) error {
	//panic("unimplemented Next")
	return nil
}

// Close implements driver.Stmt
func (stmt *gosqlite3Stmt) Close() error {
	log.Print("unimplemented Close")
	return nil
}

// Exec implements driver.Stmt
func (stmt *gosqlite3Stmt) Exec(args []driver.Value) (driver.Result, error) {
	panic("unimplemented Exec")
}

// NumInput implements driver.Stmt
func (stmt *gosqlite3Stmt) NumInput() int {
	return stmt.inputs
	//panic("unimplemented NumInput")
}

// Query implements driver.Stmt
func (stmt *gosqlite3Stmt) Query(args []driver.Value) (driver.Rows, error) {
	query := stmt.query
	for i := 0; i < stmt.inputs; i++ {
		re, err := regexp.Compile(fmt.Sprintf(`(?s)\$%d([^\d])`, i+1))
		if err != nil {
			return nil, err
		}
		switch args[i].(type) {
		case int64:
			query = re.ReplaceAllString(query, fmt.Sprintf("%d", args[i])+"$1")
		case float64:
			query = re.ReplaceAllString(query, fmt.Sprintf("%f", args[i])+"$1")
		case bool:
			query = re.ReplaceAllString(query, fmt.Sprintf("%t", args[i])+"$1")
		case string:
			query = re.ReplaceAllString(query, fmt.Sprintf("'%s'", args[i])+"$1")
		case []byte:
			query = re.ReplaceAllString(query, fmt.Sprintf("'%s'", string(args[i].([]byte)))+"$1")
		case time.Time:
			t := args[i].(time.Time)
			query = re.ReplaceAllString(query, fmt.Sprintf("'%s'", t.Format("2006-01-02 15:04:05"))+"$1")
		}
	}
	ret := new(gosqlite3Rows)
	ret.runQuery(query)
	return ret, nil
}

func (stmt *gosqlite3Stmt) parseQuery(query string) error {
	stmt.query = query
	re, err := regexp.Compile(`(?s).*?(\$\d+)\s?.*?`)
	if err != nil {
		return err
	}
	if !re.MatchString(query) {
		stmt.inputs = 0
	} else {
		all := re.FindAllStringSubmatch(query, -1)
		r := make([]string, 0)
		for _, sm := range all {
			if len(sm) > 1 {
				r = append(r, sm[1][1:])
			}
		}
		sort.Strings(r)
		last := r[len(r)-1]
		stmt.inputs, err = strconv.Atoi(last)
		if err != nil {
			return err
		}
	}
	return nil
}
