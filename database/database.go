package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error

	const POSTGRES = "host=localhost user=postgres password=insan dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	dsn := POSTGRES
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Cannot connect to database")
	}

	fmt.Println("Connect to database")
}
