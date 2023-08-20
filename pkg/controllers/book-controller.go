package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/khalilullahalfaath/BooksAPIGolang/pkg/models"
	"github.com/khalilullahalfaath/BooksAPIGolang/pkg/utils"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId, err := strconv.ParseInt(vars["id"], 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	book, _ := models.GetBookById(bookId)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId, err := strconv.ParseInt(vars["id"], 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	deletedBook := models.DeleteBookById(bookId)
	res, _ := json.Marshal(deletedBook)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBookById(w http.ResponseWriter, r *http.Request) {
	var updatedBook = &models.Book{}
	utils.ParseBody(r, updatedBook)
	vars := mux.Vars(r)
	bookId, err := strconv.ParseInt(vars["id"], 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	bookDetails, db := models.GetBookById(bookId)
	if updatedBook.Title != "" {
		bookDetails.Title = updatedBook.Title
	}
	if updatedBook.Author != "" {
		bookDetails.Author = updatedBook.Author
	}
	if updatedBook.Publisher != "" {
		bookDetails.Publisher = updatedBook.Publisher
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
