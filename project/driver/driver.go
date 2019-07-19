package driver

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1878"
	dbname   = "postgres"
)

var db *sql.DB

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectDB() *sql.DB {
	var err error

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)
	logFatal(err)

	err = db.Ping()
	logFatal(err)

	return db
}
