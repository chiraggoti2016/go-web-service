// routes/routes.go
package routes

import (
    "github.com/chiraggoti2016/go-web-service/controllers"
    "github.com/gorilla/mux"
)

func RegisterBookRoutes(r *mux.Router) {
    r.HandleFunc("/api/books", controllers.GetBooks).Methods("GET")
    r.HandleFunc("/api/books/{id}", controllers.GetBook).Methods("GET")
    r.HandleFunc("/api/books", controllers.CreateBook).Methods("POST")
    r.HandleFunc("/api/books/{id}", controllers.UpdateBook).Methods("PUT")
    r.HandleFunc("/api/books/{id}", controllers.DeleteBook).Methods("DELETE")
}
