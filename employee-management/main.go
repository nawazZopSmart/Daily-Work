package main

import (
	"net/http"

	s "github.com/SN786/employee-management/internal/store/driver"
	"github.com/gorilla/mux"

	EmpHandler "github.com/SN786/employee-management/internal/http"
	ems "github.com/SN786/employee-management/internal/store/employee"
)

func main() {
	DBName := "emp"
	// tableName := "employee"
	// s.CreateTable(DBName, tableName)
	DbCon := s.DbConn(DBName)

	datastore := ems.New(DbCon)
	handler := EmpHandler.New(datastore)

	r := mux.NewRouter()

	r.HandleFunc("/emp", handler.Handler)

	http.ListenAndServe(":8090", r)

}
