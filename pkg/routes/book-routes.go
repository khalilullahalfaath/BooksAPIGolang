package routes

// golang uses the package name as the root path for the project
import (
	"github.com/gorilla/mux"
	"github.com/khalilullahalfaath/BooksAPIGolang/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.UpdateBookById).Methods("PUT")
	router.HandleFunc("/book/{id}", controllers.DeleteBookById).Methods("DELETE")
}
