package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func connexion() *sql.DB {
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
	} else {
		fmt.Println("connected to psql")
	}
	//defer db.Close()
	return db
}
