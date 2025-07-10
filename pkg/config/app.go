package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var (
	db *gorm.DB
)

func Connect() {
	// d, err := gorm.Open("mysql", "yash_death:Yash@1234/simplerest?charset=utf8&parseTime=True&loc=Local")
	d,err := gorm.Open("sqlite3", "./bookstore.db")
	if err != nil {
		panic(err)
	}

	db = d
	log.Println("Database connected successfully.")
}

func GetDB() *gorm.DB {
	return db
}
