package router

import (
	"example/go/mongodb/controller"
	"fmt"

	"github.com/gorilla/mux"
)

func MyRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/students", controller.GetAllStudent).Methods("GET")
	r.HandleFunc("/api/student/{id}", controller.GetOneStudent).Methods("GET")
	r.HandleFunc("/api/student", controller.CreateStudent).Methods("POST")
	r.HandleFunc("/api/student/{id}", controller.DeleteStudent).Methods("DELETE")
	r.HandleFunc("/api/student/{id}", controller.UpdateStudent).Methods("PATCH")
	r.HandleFunc("/api/students", controller.ClearStudent).Methods("DELETE")
	fmt.Println("Run Server")
	return r
}
