package crud_rest

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db, _ = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/emp")

type Emp struct {
	Id    int
	Name  string
	Email string
	Role  string
}

var X []Emp

func CreateEmp(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var emp Emp
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &emp)
	_, err := db.Exec("INSERT INTO employee (Name,Email,role) values(?,?,?)", emp.Name, emp.Email, emp.Role)
	if err != nil {
		log.Println("error:", err)
		w.WriteHeader(500)
		_, _ = w.Write([]byte("something unexpected happened"))
		return
	}

}

func GetByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	Id := vars["id"]
	intID, _ := strconv.Atoi(Id)

	var empl Emp
	err := db.QueryRow("select * from employee where id=?", intID).Scan(&empl.Id, &empl.Name, &empl.Email, &empl.Role)
	if err != nil {
		w.Write([]byte(`{"Error": "No_Data_Found"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(empl)

}

func GetAll(w http.ResponseWriter, r *http.Request) {

	res, err := db.Query("select * from employee")

	if err != nil {
		w.Write([]byte(`{"Error": "No_Data_Found"}`))
		return
	}
	var employees []Emp
	var id int
	var name, email, role string
	for res.Next() {
		e := res.Scan(&id, &name, &email, &role)
		if e != nil {
			w.Write([]byte(`{"Error": "While scanning"}`))
			return
		}
		emp := Emp{id, name, email, role}
		employees = append(employees, emp)

	}
	json.NewEncoder(w).Encode(employees)

	w.WriteHeader(http.StatusOK)
}

func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emp Emp
	_ = json.NewDecoder(r.Body).Decode(&emp)

	ID := emp.Id
	name := emp.Name
	email := emp.Email
	role := emp.Role

	res, err := db.Prepare("UPDATE employee SET name = ?, email = ?, role = ? WHERE id = ?")
	if err != nil {
		w.Write([]byte(`{"Error":"Someting Went Wrong in Prepare"}`))

	}
	defer res.Close()
	_, err2 := res.Exec(name, email, role, ID)

	if err2 != nil {
		w.Write([]byte(`{"Error":"Someting Went Wrong in Prepare"}`))
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteByID(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var emp Emp
	_ = json.NewDecoder(r.Body).Decode(&emp)

	ID := emp.Id

	_, err := db.Exec("delete from employee where id=?", ID)

	if err != nil {
		w.Write([]byte(`{"Error": "No_Data_Found"}`))
		return
	}
	w.WriteHeader(http.StatusOK)

}
