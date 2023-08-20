package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/khalilullahalfaath/BooksAPIGolang/pkg/routes"
)

func main() {
	// create a new router
	router := mux.NewRouter()
	// register routes
	routes.RegisterBookStoreRoutes(router)
	http.Handle("/", router)
	// start the server
	log.Fatal(http.ListenAndServe(":8080", router))
}

