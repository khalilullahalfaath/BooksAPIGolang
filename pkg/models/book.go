package models

import (
	"github.com/jinzhu/gorm"
	"github.com/khalilullahalfaath/BooksAPIGolang/pkg/config"
)

// create db object
var db *gorm.DB

// Book is a struct that represents the Book model in the database
// gorm.Model is a struct that contains the following fields:
type Book struct {
	gorm.Model
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
}

// init is a special function that is called before the main function
func init() {
	// connect to the database
	config.Connect()
	// get a handle to the DB object
	db = config.GetDB()
	// run the migrations: todo struct
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(id int64) (*Book, *gorm.DB) {
	var book Book
	db := db.Where("id = ?", id).Find(&book)
	return &book, db
}

func DeleteBookById(id int64) Book {
	var book Book
	db.Where("id = ?", id).Delete(&book)
	return book
}
