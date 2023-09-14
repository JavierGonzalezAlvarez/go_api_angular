package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Users struct {
	Iduser    *int      `json:"iduser"`
	Username  *string   `json:"username"`
	Password  *string   `json:"password"`
	Email     *string   `json:"email"`
	Token     *string   `json:"token"`
	CreatedAt time.Time `json:"createdat"`
}

type HeaderPostgres struct {
	Idheader      int       `json:"idheader"`
	Companyname   *string   `json:"companyname"`
	Address       *string   `json:"address"`
	NumberInvoice *int      `json:"numberinvoice"`
	DateTime      time.Time `json:"datetime" format:"2006-01-02 15:04:05"`
	CreatedAt     time.Time `json:"createdat"`
}

type DetailPostgres struct {
	IdDetail    int       `json:"detailid"`
	IdHeader    int       `json:"idheader"`
	Description *string   `json:"description"`
	Units       *int      `json:"units"`
	Price       *float32  `json:"price"`
	CreatedAt   time.Time `json:"createdat"`
}

type InvoicePostgres struct {
	Header
	Iddetail []Detail `json:"iddetail"`
}

type Response struct {
	Token      string    `json:"token"`
	Expiracion time.Time `json:"expiracion"`
	User       string    `json:"username"`
}

type ResponseLogin struct {
	Email string `json:"email"`
}

// Create a response structure
type ResponseTotalHeaders struct {
	TotalCount int              `json:"totalCount"`
	Results    []HeaderPostgres `json:"results"`
}

func get_all_users() []Users {
	fmt.Println("sql users")

	err := godotenv.Load("./env/env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var db = connexion()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	Logger.Info("Successfully connected!")
	defer db.Close()

	//query
	rows, err := db.Query(`SELECT id, username, email FROM usuario`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("list of users")
	fmt.Println("---------------")
	var result = []Users{}
	for rows.Next() {
		var item Users
		rows.Scan(&item.Iduser, &item.Username, &item.Email)
		result = append(result, item)
	}

	return result
}

func get_all_headers_invoices_total() []ResponseTotalHeaders {

	Logger.Info("sql all headers of invoices")

	err := godotenv.Load("./env/env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var db = connexion()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	Logger.Info("Successfully connected!")
	defer db.Close()

	// query
	rows, err := db.Query(`SELECT * FROM "header"`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	Logger.Info("list of records")

	var result = []HeaderPostgres{}

	for rows.Next() {
		var item HeaderPostgres
		rows.Scan(&item.Idheader, &item.Companyname, &item.Address, &item.NumberInvoice, &item.DateTime, &item.CreatedAt)
		result = append(result, item)
	}

	for _, item := range result {
		fmt.Printf("Idheader: %d\n", item.Idheader)
		fmt.Printf("Companyname: %s\n", *item.Companyname)
		fmt.Printf("Address: %s\n", *item.Address)
		fmt.Printf("NumberInvoice: %d\n", item.NumberInvoice)
		fmt.Printf("DateTime: %s\n", item.DateTime)
		fmt.Printf("CreatedAt: %s\n", item.CreatedAt)
		fmt.Println("---------------")
	}

	totalCount := len(result)
	// Create a response instance
	response := ResponseTotalHeaders{
		TotalCount: totalCount,
		Results:    result,
	}

	return []ResponseTotalHeaders{response}

}

func get_all_headers_invoices() []HeaderPostgres {

	Logger.Info("sql all headers of invoices")

	err := godotenv.Load("./env/env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var db = connexion()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	Logger.Info("Successfully connected!")
	defer db.Close()

	// query
	rows, err := db.Query(`SELECT * FROM "header"`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	Logger.Info("list of records")

	var result = []HeaderPostgres{}

	for rows.Next() {
		var item HeaderPostgres
		rows.Scan(&item.Idheader, &item.Companyname, &item.Address, &item.NumberInvoice, &item.DateTime, &item.CreatedAt)
		result = append(result, item)
	}

	for _, item := range result {
		fmt.Printf("Idheader: %d\n", item.Idheader)
		fmt.Printf("Companyname: %s\n", *item.Companyname)
		fmt.Printf("Address: %s\n", *item.Address)
		fmt.Printf("NumberInvoice: %d\n", item.NumberInvoice)
		fmt.Printf("DateTime: %s\n", item.DateTime)
		fmt.Printf("CreatedAt: %s\n", item.CreatedAt)
		fmt.Println("---------------")
	}

	return result

}

func get_one_header_invoice(id int) []HeaderPostgres {
	Logger.Info("Sql get one header invoice")

	err := godotenv.Load("./env/env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var db = connexion()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	Logger.Info("Successfully connected!")
	defer db.Close()

	rows, err := db.Query("select * from header where id_header = $1", id)
	if err != nil {
		log.Fatal(err)
	} else {
		Logger.Info("Row selected successfully!")
	}
	defer rows.Close()

	var result = []HeaderPostgres{}
	for rows.Next() {
		var item HeaderPostgres
		rows.Scan(&item.Idheader, &item.Companyname, &item.Address, &item.NumberInvoice, &item.DateTime, &item.CreatedAt)
		result = append(result, item)
	}

	return result
}

func create_user_sql(dataPost []uint8) []ResponseLogin {
	Logger.Info("insert user sql with token")

	err := godotenv.Load("./env/env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var db = connexion()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	Logger.Info("Successfully connected!")

	defer db.Close()
	fmt.Println("json from api post", string(dataPost))

	// decode structure data: from json to struct
	var InsertJson Users
	json.Unmarshal([]byte(dataPost), &InsertJson)

	// create token for user
	signedToken := create_token(dataPost)
	myTime := time.Now()
	fmt.Println("datetime", myTime)

	//insert in postgres
	sqlStatement := `INSERT INTO usuario (username, password, email, token, created_at) VALUES ($1, $2, $3, $4, $5)`
	result, err := db.Exec(sqlStatement, InsertJson.Username, InsertJson.Password, InsertJson.Email, signedToken, myTime)
	if err != nil {
		log.Fatal(err)

	} else {
		// If there is no error, you can retrieve additional information about the execution
		rowCount, err := result.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}

		// Check the number of rows affected
		fmt.Printf("Rows inserted successfully! Rows affected: %d, token: %s\n", rowCount, signedToken)

		//fmt.Println("\nRow inserted successfully!, token: ", signedToken)
		response := ResponseLogin{Email: *InsertJson.Email}
		responses := []ResponseLogin{response}

		return responses
	}

	return nil
}

func create_header_invoice(dataPost []uint8) {
	Logger.Info("Create a header of an invoice")

	err := godotenv.Load("./env/env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var db = connexion()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	Logger.Info("Successfully connected!")
	defer db.Close()

	Logger.Info("json from api post", string(dataPost))

	// decode structure data: from json to struct
	var InsertJson HeaderPostgres
	json.Unmarshal([]byte(dataPost), &InsertJson)

	Logger.Info("type of InsertJson = ", reflect.TypeOf(InsertJson))
	fmt.Printf("Id Header: %v, Company Name %s \n", InsertJson.Idheader, *InsertJson.Companyname)

	myTime := time.Now()

	//insert in postgres
	sqlStatement := `INSERT INTO header (companyname, address, numberinvoice, date_time) VALUES ($1, $2, $3, $4)`
	_, err = db.Exec(sqlStatement, InsertJson.Companyname, InsertJson.Address, InsertJson.NumberInvoice, myTime)
	if err != nil {
		log.Fatal(err)
	} else {
		Logger.Info("Row inserted successfully!")
	}

	return
}

func update_one_header(dataPost []uint8) {
	Logger.Info("update sql")

	err := godotenv.Load("./env/env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var db = connexion()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	Logger.Info("Successfully connected!")
	defer db.Close()

	Logger.Info("json from api post", string(dataPost))

	// decode structure data: from json to struct
	var UpdateJson HeaderPostgres
	json.Unmarshal([]byte(dataPost), &UpdateJson)
	Logger.Info("type of InsertJson = ", reflect.TypeOf(UpdateJson))

	fmt.Printf("Id Header: %v, Company Name %s \n", UpdateJson.Idheader, *UpdateJson.Companyname)

	//update in postgres
	rows, err := db.Query("UPDATE header SET companyname = $2 WHERE id_header = $1", UpdateJson.Idheader, UpdateJson.Companyname)
	if err != nil {
		log.Fatal(err)
	} else {
		Logger.Info("Row updates successfully!")
	}
	defer rows.Close()

	return
}

func delete_one_header_invoice(id int) {

	Logger.Info(("Delete one header invoice"))

	err := godotenv.Load("./env/env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var db = connexion()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	Logger.Info("Successfully connected!")
	defer db.Close()

	rows, err := db.Query("delete from header where id_header = $1", id)
	if err != nil {
		Logger.Error(err)
	} else {
		fmt.Println("\nRow deletes successfully!")
	}
	defer rows.Close()

}

func create_one_invoice(dataPost []uint8) {
	Logger.Info("Create invoice 8header & detail)")

	err := godotenv.Load("./env/env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var db = connexion()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	Logger.Info("Successfully connected!")
	defer db.Close()

	// decode structure data: from json to struct, datetime format must be correct
	//var InsertJson InvoicePostgres // or
	InvoiceStruct := new(InvoicePostgres)
	fmt.Println("--------------------------------------------------")
	fmt.Println(reflect.TypeOf(InvoiceStruct))         //*main.InvoicePostgress
	fmt.Println(reflect.ValueOf(InvoiceStruct).Kind()) //ptr

	_ = json.Unmarshal([]byte(dataPost), &InvoiceStruct)
	if err != nil {
		fmt.Println("error format", dataPost)
		return
	} else {
		fmt.Println("json", string(dataPost))
	}

	fmt.Println("--------------------------------------------------")
	fmt.Println("values of the json: ", *InvoiceStruct)
	fmt.Println("type of InsertJson = ", reflect.TypeOf(InvoiceStruct))
	fmt.Println("type of dataPost = ", reflect.TypeOf(dataPost))
	fmt.Printf("Id Header: %v, Company Name: %s, Number Invoice: %d, Date Time: %s \n", InvoiceStruct.Idheader, InvoiceStruct.Companyname, InvoiceStruct.NumberInvoice, InvoiceStruct.DateTime)
	fmt.Printf("Detail: %v \n", InvoiceStruct.Iddetail)

	for i, detail := range InvoiceStruct.Iddetail {
		fmt.Printf("Detail %d:\n", i+1)
		fmt.Println("Detail ID:", detail.IdDetail)
		fmt.Println("Header ID:", detail.IdHeader)
		fmt.Println("Description:", detail.Description)
		fmt.Println("Units:", detail.Units)
		fmt.Println("Price:", detail.Price)
		fmt.Println("Created at:", detail.CreatedAt)
	}
	fmt.Println("--------------------------------------------------")

	myTime := time.Now()
	sqlStatementHeader := `INSERT INTO header (companyname, address, numberinvoice, date_time, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id_header`
	_, err = db.Exec(sqlStatementHeader, InvoiceStruct.Companyname, InvoiceStruct.Address, InvoiceStruct.NumberInvoice, InvoiceStruct.DateTime, myTime)
	if err != nil {
		log.Fatal(err)
		//db.Rollback()
		return
	} else {
		fmt.Println("\nRow Header inserted successfully!")
	}

	//get last id_header inserted
	var id_header int
	err = db.QueryRow("SELECT id_header FROM header ORDER BY id_header DESC LIMIT 1").Scan(&id_header)
	if err != nil {
		return
	}

	for _, detail := range InvoiceStruct.Iddetail {
		sqlStatementDetail := `INSERT INTO detail (id_header, description, units, price, created_at) VALUES ($1, $2, $3, $4, $5)`
		_, err := db.Exec(sqlStatementDetail, id_header, detail.Description, detail.Units, detail.Price, detail.CreatedAt)

		if err != nil {
			//db.Rollback()
			return
		} else {
			Logger.Info("Row Detail inserted successfully!")
		}
	}
	return
}
