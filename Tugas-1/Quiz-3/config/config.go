package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username string = "root"
	password string = ""
	dbname   string = "quiz_3"
)

var dsn = fmt.Sprintf("%v:%v@/%v", username, password, dbname)

func ConnectToMySQL() (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
