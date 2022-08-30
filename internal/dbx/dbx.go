package dbx

import (
	"log"
	
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func Open(dsn string) (*DB, error) {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &DB{DB: db}, nil
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
