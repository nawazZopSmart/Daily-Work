package main

import (
	"net/http"

	s "github.com/SN786/mysql_rest_api/crud_rest"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/emp", s.CreateEmp).Methods("POST")
	r.HandleFunc("/emp/{id}", s.GetByID).Methods("GET")
	r.HandleFunc("/emp", s.GetAll).Methods("GET")
	r.HandleFunc("/emp", s.Update).Methods("PATCH")
	r.HandleFunc("/emp", s.DeleteByID).Methods("DELETE")

	http.ListenAndServe(":8090", r)

}
