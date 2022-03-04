package employee

import (
	"database/sql"
	"errors"

	model "github.com/SN786/employee-management/entities"
	_ "github.com/go-sql-driver/mysql"
)

type EmployeeStorer struct {
	db *sql.DB
}

func New(db *sql.DB) *EmployeeStorer {
	return &EmployeeStorer{db: db}
}

//Creating Table
// func CreateTable(DBName string, tableName string) {
// 	db := DbConn(DBName)
// 	defer db.Close()
// 	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %v(id int PRIMARY KEY AUTO_INCREMENT, Name varchar(30) NOT NULL, Email varchar(30), role varchar(30));", tableName)
// 	_, err := db.Exec(query)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

//Create
func (e EmployeeStorer) InsertData(name string, email string, role string) error {
	_, err := e.db.Exec("INSERT INTO employee (Name,Email,role) values(?,?,?)", name, email, role)
	if err != nil {
		return errors.New("not inserted")
	}
	return nil
}

// Read
func (e EmployeeStorer) GetDetailsById(id int) (*model.Employee, error) {
	var empl model.Employee
	err := e.db.QueryRow("select * from employee where id=?", id).Scan(&empl.ID, &empl.Name, &empl.Email, &empl.Role)
	if err != nil {
		return nil, err
	}
	return &empl, nil
}

//Update
func (e EmployeeStorer) UpdateById(id int, Name string, Email string, role string) error {

	res, err := e.db.Prepare("UPDATE employee SET name = ?, email = ?, role = ? WHERE id = ?")
	if err != nil {
		return errors.New("query doesn't prepare")
	}
	defer res.Close()
	_, err2 := res.Exec(Name, Email, role, id)

	if err2 != nil {
		return err2
	}
	return nil
}

//Delete
func (e EmployeeStorer) DeleteById(id int) error {

	_, err := e.db.Exec("delete from employee where id=?", id)

	if err != nil {
		return err
	}

	return nil

}
