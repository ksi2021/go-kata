package main

import (
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBConnect() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=pass dbname=kata port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, err
}
