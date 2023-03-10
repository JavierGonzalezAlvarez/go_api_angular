package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
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

func q_sql() []HeaderPostgres {
	fmt.Println("sql to psql")

	err := godotenv.Load("./env/env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	/*
		for _, e := range os.Environ() {
			pair := strings.SplitN(e, "=", 2)
			fmt.Printf("%s: %s\n", pair[0], pair[1])
		}
	*/
	//fmt.Println("user", os.Getenv("USER"))

	var db = connexion()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	defer db.Close()

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

func insert_sql(dataPost []uint8) {
	fmt.Println("insert to psql")

	err := godotenv.Load("./env/env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var db = connexion()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	defer db.Close()
	fmt.Println("json from api post", string(dataPost))

	// decode structure data: from json to struct
	var InsertJson HeaderPostgres
	json.Unmarshal([]byte(dataPost), &InsertJson)
	fmt.Println("type of InsertJson = ", reflect.TypeOf(InsertJson))
	fmt.Printf("Id Header: %v, Company Name %s", InsertJson.Idheader, InsertJson.Companyname)

	//insert in postgres
	sqlStatement := `INSERT INTO header (companyname) VALUES ($1)`
	_, err = db.Exec(sqlStatement, InsertJson.Companyname)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("\nRow inserted successfully!")
	}

	/*
		rows, err := db.Query(`SELECT * FROM "header"`)
		if err != nil {
			log.Fatal(err)
		}
	*/
	//defer rows.Close()

	return

}
