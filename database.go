package main

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "gopher"
	password = "gopherpw"
	dbname   = "pubdb"
)

func openConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func RunQuery(query string, args ...interface{}) *sql.Rows {
	conn := openConnection()
	defer conn.Close()

	rows, err := conn.Query(query, args...)

	if err != nil {
		log.Fatal(err)
	}

	return rows
}

func Insert(insert string, args ...interface{}) bool {
	conn := openConnection()
	defer conn.Close()

	_, err := conn.Exec(insert, args...)
	if err != nil {
		log.Fatal(err)
	}

	return err == nil
}
