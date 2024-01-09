package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//postgresql://postgres:postgres@db:5432/nest_js_api?schema=public
func Init() *gorm.DB {
	dbURL := "postgres://postgres:postgres@db:5432/nest_js_api"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
