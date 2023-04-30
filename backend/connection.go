package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

/*
const (
	host     = "localhost"
	port     = 5432
	user     = "test"
	password = "2525_ap"
	dbname   = "invoices"
)
*/

func connexion() *sql.DB {
	fmt.Println("connect to psql")

	//connect db
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("LOCALHOST"),
		os.Getenv("PORT"),
		os.Getenv("USUARIO"),
		os.Getenv("PASSWORD"),
		os.Getenv("DBNAME"),
	)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	//defer db.Close()
	return db
}
