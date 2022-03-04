package driver

import (
	"database/sql"
	

	_ "github.com/go-sql-driver/mysql"
)

func DbConn(db_name string) (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "password"
	dbName := db_name
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}