package main

import (
	"net/http"

	s "github.com/SN786/rest_api/employee"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/create", s.CreateEmp).Methods("POST")
	r.HandleFunc("/get", s.GetByID).Methods("GET")
	r.HandleFunc("/getall", s.GetAll).Methods("GET")
	r.HandleFunc("/update", s.Update).Methods("PATCH")
	r.HandleFunc("/delete", s.DeleteByID).Methods("DELETE")

	http.ListenAndServe(":8080", r)

}
