package dbx

import (
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type DB struct {
	*gorm.DB
}

func (db *DB) Close() {
	if d, err := db.DB.DB(); err != nil {
		log.Fatal(err)
	} else {
		if err = d.Close(); err != nil {
			log.Fatal(err)
		}
	}
}

func Open(dsn string) (*DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &DB{DB: db}, nil
}

func Sqlite(dsn string) (*DB, error) {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &DB{DB: db}, nil
}

func CopyTable(rows any, db *DB, bak *DB) error {
	return db.FindInBatches(rows, 100, func(tx *gorm.DB, batch int) error {
		return bak.Create(rows).Error
	}).Error
}
