package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "test"
	password = "2525_ap"
	dbname   = "invoices"
)

type HeaderPostgres struct {
	Idheader      int       `json:"courseid"`
	Companyname   string    `json:"companyname"`
	Address       string    `json:"address"`
	NumberInvoice int       `json:"numberinvoice"`
	DateTime      time.Time `json:"datetime"`
	CreatedAt     time.Time `json:"createdat"`
}

func postgres() []HeaderPostgres {
	fmt.Println("connect to psql")

	err := godotenv.Load("./env/env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//connect db
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("LOCALHOST"), port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")

	//query
	rows, err := db.Query(`SELECT * FROM "header"`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var result = []HeaderPostgres{}
	for rows.Next() {
		var item HeaderPostgres

		//json.Marshal works with predefined datatypes. For JSON we need to create two functions -
		//Scan() - To parse JSON from database to Go struct.
		rows.Scan(&item.Idheader, &item.Companyname, &item.Address, &item.NumberInvoice, &item.DateTime, &item.CreatedAt)

		result = append(result, item)
		fmt.Println(result)
	}
	return result

}
