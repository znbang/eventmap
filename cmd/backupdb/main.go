package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/znbang/eventmap/db"
)

var tables = []string{
	"users",
	"user_sessions",
	"books",
	"book_jobs",
	"chapters",
	"events",
}

func genInsertSql(table string, columns []string) string {
	var cols []string
	for i := 0; i < len(columns); i++ {
		cols = append(cols, "?")
	}
	return fmt.Sprintf("insert into %s (%s) values (%s)", table, strings.Join(columns, ","), strings.Join(cols, ","))
}

func sliceScan(rs *sql.Rows) ([]interface{}, error) {
	columns, err := rs.Columns()
	if err != nil {
		return []interface{}{}, err
	}

	values := make([]interface{}, len(columns))
	for i := range values {
		values[i] = new(interface{})
	}

	err = rs.Scan(values...)
	if err != nil {
		return values, err
	}

	for i := range columns {
		values[i] = *(values[i].(*interface{}))
	}

	return values, rs.Err()
}

func prepareQueryStmt(src *sql.DB, table string, id any) (*sql.Stmt, error) {
	if id == nil {
		return src.Prepare(fmt.Sprintf("select * from %s order by id limit 1000", table))
	} else {
		return src.Prepare(fmt.Sprintf("select * from %s where id>? order by id limit 1000", table))
	}
}

func query(src *sql.DB, table string, id any) ([][]interface{}, []string, error) {
	queryStmt, err := prepareQueryStmt(src, table, id)
	if err != nil {
		return nil, nil, err
	}
	defer queryStmt.Close()

	var rs *sql.Rows
	if id == nil {
		rs, err = queryStmt.Query()
	} else {
		rs, err = queryStmt.Query(id)
	}
	if err != nil {
		return nil, nil, err
	}
	defer rs.Close()

	columns, err := rs.Columns()
	if err != nil {
		return nil, nil, err
	}

	var rows [][]interface{}

	for rs.Next() {
		row, err := sliceScan(rs)
		if err != nil {
			return nil, nil, err
		}
		rows = append(rows, row)
	}

	return rows, columns, nil
}

func copyRows(table string, src *sql.DB, bak *sql.DB, id any) (any, error) {
	rows, columns, err := query(src, table, id)
	if err != nil {
		return nil, err
	}

	insertSql := genInsertSql(table, columns)

	tx, err := bak.Begin()
	insertStmt, err := tx.Prepare(insertSql)
	if err != nil {
		return nil, err
	}

	id = nil

	for _, row := range rows {
		_, err = insertStmt.Exec(row...)
		if err != nil {
			_ = tx.Rollback()
			return nil, err
		}

		id = row[0]
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return id, nil
}

func copyTable(table string, src *sql.DB, bak *sql.DB) error {
	id, err := copyRows(table, src, bak, nil)
	if err != nil {
		return err
	}
	for id != nil {
		id, err = copyRows(table, src, bak, id)
		if err != nil {
			return err
		}
	}
	return nil
}

func backup(srcDSN, bakDSN, schema string, tables []string) error {
	src, err := sql.Open("sqlite3", srcDSN)
	if err != nil {
		return fmt.Errorf("open database failed: %w", err)
	}
	defer src.Close()

	bak, err := sql.Open("sqlite3", bakDSN)
	if err != nil {
		return fmt.Errorf("open backup database failed: %w", err)
	}
	defer bak.Close()

	if _, err = bak.Exec(schema); err != nil {
		return fmt.Errorf("create schema failed: %w", err)
	}

	for _, table := range tables {
		log.Printf("copy table %s", table)
		if err = copyTable(table, src, bak); err != nil {
			return fmt.Errorf("copy table failed: %w", err)
		}
	}

	return nil
}

func main() {
	srcDSN := "eventmap.db"
	bakDSN := "eventmap." + time.Now().Format("20060102150405") + ".db"
	if err := backup(srcDSN, bakDSN, db.SqliteSchema, tables); err != nil {
		fmt.Println("backup failed:", err)
	}

	fmt.Println("Done")
}
