package config

import (
	"github.com/fathurhidayat3/gorestapi/structs"
	"github.com/jinzhu/gorm"
)

// DBInit create connection to database
func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:Password@tcp(127.0.0.1:3306)/gobukudb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(structs.Buku{})
	return db
}
