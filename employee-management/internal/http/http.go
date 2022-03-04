package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	model "github.com/SN786/employee-management/entities"
	"github.com/SN786/employee-management/internal/store"
)

type EmpHandler struct {
	datastore store.EmployeManager
}

func New(employee store.EmployeManager) *EmpHandler {
	return &EmpHandler{datastore: employee}
}

//This Function handles Routing
func (e EmpHandler) Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		e.GetByID(w, r)
	case http.MethodPost:
		e.CreateEmp(w, r)
	case http.MethodPatch:
		e.Update(w, r)
	case http.MethodDelete:
		e.DeleteByID(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

//This Function Call GetDetailsById Function after getting data from request body
func (e EmpHandler) GetByID(w http.ResponseWriter, r *http.Request) {

	Id := r.URL.Query().Get("id")
	intID, _ := strconv.Atoi(Id)
	fmt.Println(intID)
	var empl *model.Employee
	empl, err := e.datastore.GetDetailsById(intID)

	if err != nil {
		w.Write([]byte(`{"Error": "No_Data_Found"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(empl)

}

//This Function Call InsertData Function after getting data from request body
func (e EmpHandler) CreateEmp(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var emp model.Employee
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &emp)

	err := e.datastore.InsertData(emp.Name, emp.Email, emp.Role)

	if err != nil {
		w.WriteHeader(500)
		_, _ = w.Write([]byte("something unexpected happened"))
		return
	}
	w.WriteHeader(http.StatusOK)

}

//This Function Call UpdateById Function after getting data from request body
func (e EmpHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emp model.Employee
	_ = json.NewDecoder(r.Body).Decode(&emp)

	ID := emp.ID
	name := emp.Name
	email := emp.Email
	role := emp.Role

	err := e.datastore.UpdateById(ID, name, email, role)
	if err != nil {
		w.Write([]byte(`{"Error":"Someting Went Wrong in Prepare"}`))
	}

	w.WriteHeader(http.StatusOK)
}

//This Function Call DeleteById Function after getting data from request body
func (e EmpHandler) DeleteByID(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var emp model.Employee
	_ = json.NewDecoder(r.Body).Decode(&emp)

	ID := emp.ID

	err := e.datastore.DeleteById(ID)

	if err != nil {
		w.Write([]byte(`{"Error": "No_Data_Found"}`))
		return
	}
	w.WriteHeader(http.StatusOK)

}
