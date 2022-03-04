package store

import (
	e "github.com/SN786/employee-management/entities"
)

type EmployeManager interface {
	InsertData(string, string, string) error
	GetDetailsById(int) (*e.Employee, error)
	UpdateById(int, string, string, string) error
	DeleteById(int) error
}
