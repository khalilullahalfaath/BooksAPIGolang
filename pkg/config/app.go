package config

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var (
	db *gorm.DB
)

// connect to the database
func Connect() {
	// open a db connection
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// open a db connection
	d, err := gorm.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME")+"?charset=utf8&parseTime=True&loc=Local")

	// d, err := gorm.Open("mysql", "root:root@tcp(localhost:3306)/booksapp?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

// return a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
