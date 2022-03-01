package employee

import (
	"encoding/json"

	"net/http"
)

type Emp struct {
	Id    string
	Name  string
	Email string
}

var X []Emp

func CreateEmp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emp Emp
	_ = json.NewDecoder(r.Body).Decode(&emp)
	X = append(X, emp)
	json.NewEncoder(w).Encode(&emp)
	w.WriteHeader(http.StatusOK)

}

func GetByID(w http.ResponseWriter, r *http.Request) {

	Id := r.URL.Query().Get("id")
	var ok bool
	for _, e := range X {

		if e.Id == Id {
			ok = true
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(e)
		}
	}
	if !ok {
		w.Write([]byte(`{"Error": "No_Data_Found"}`))
	}
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	for _, e := range X {
		json.NewEncoder(w).Encode(e)

	}
	w.WriteHeader(http.StatusOK)
}

func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emp Emp
	_ = json.NewDecoder(r.Body).Decode(&emp)

	ID := emp.Id
	name := emp.Name
	email := emp.Email

	var empToUpdate Emp
	var indexToUpdate int
	for i, e := range X {
		if e.Id == ID {
			empToUpdate = e
			indexToUpdate = i
			break
		}
	}
	if empToUpdate.Id == "" {
		w.Write([]byte(`{"Error": "No_Data_Found"}`))
	}
	if len(name) != 0 {
		empToUpdate.Name = name
	}
	if len(email) != 0 {
		empToUpdate.Email = email
	}
	X[indexToUpdate] = empToUpdate

	w.WriteHeader(http.StatusOK)
}

func DeleteByID(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var emp Emp
	_ = json.NewDecoder(r.Body).Decode(&emp)

	ID := emp.Id
	var ok bool
	for i, e := range X {
		if e.Id == ID {
			ok = true
			X = append(X[:i], X[i+1:]...)
			break
		}
	}
	if !ok {
		w.Write([]byte(`{"Error": "No_Data_Found"}`))
	}
	w.WriteHeader(http.StatusOK)

}
