package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
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

type HeaderPostgres struct {
	Idheader      int       `json:"idheader"`
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

	for _, e := range os.Environ() {

		pair := strings.SplitN(e, "=", 2)
		fmt.Printf("%s: %s\n", pair[0], pair[1])
	}

	fmt.Println("user", os.Getenv("USER"))
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
		rows.Scan(&item.Idheader, &item.Companyname, &item.Address, &item.NumberInvoice, &item.DateTime, &item.CreatedAt)
		result = append(result, item)
		fmt.Println(result)
	}
	return result

}
