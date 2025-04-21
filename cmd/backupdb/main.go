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
	for range columns {
		cols = append(cols, "?")
	}
	return fmt.Sprintf("insert into %s (%s) values (%s)", table, strings.Join(columns, ","), strings.Join(cols, ","))
}

func sliceScan(rs *sql.Rows) ([]any, error) {
	columns, err := rs.Columns()
	if err != nil {
		return []any{}, err
	}

	values := make([]any, len(columns))
	for i := range values {
		values[i] = new(any)
	}

	err = rs.Scan(values...)
	if err != nil {
		return values, err
	}

	for i := range columns {
		values[i] = *(values[i].(*any))
	}

	return values, rs.Err()
}

func query(src *sql.DB, table string, id any) ([][]any, []string, error) {
	queryStmt, err := func() (*sql.Stmt, error) {
		if id == nil {
			return src.Prepare(fmt.Sprintf("select * from %s order by id limit 1000", table))
		} else {
			return src.Prepare(fmt.Sprintf("select * from %s where id>? order by id limit 1000", table))
		}
	}()
	if err != nil {
		return nil, nil, err
	}
	defer queryStmt.Close()

	rs, err := func() (*sql.Rows, error) {
		if id == nil {
			return queryStmt.Query()
		} else {
			return queryStmt.Query(id)
		}
	}()
	if err != nil {
		return nil, nil, err
	}
	defer rs.Close()

	columns, err := rs.Columns()
	if err != nil {
		return nil, nil, err
	}

	var rows [][]any

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

	err = withTx(bak, func(tx *sql.Tx) error {
		if insertStmt, err := tx.Prepare(insertSql); err != nil {
			return err
		} else {
			for _, row := range rows {
				_, err = insertStmt.Exec(row...)
				if err != nil {
					return err
				}

				id = row[0]
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return id, nil
}

func withTx(db *sql.DB, fn func(tx *sql.Tx) error) error {
	if tx, err := db.Begin(); err != nil {
		return err
	} else if err = fn(tx); err != nil {
		_ = tx.Rollback()
		return err
	} else if err = tx.Commit(); err != nil {
		return err
	}
	return nil
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
