package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Users struct {
	Iduser    *int      `json:"iduser"`
	Username  *string   `json:"username"`
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
	fmt.Println("Successfully connected!")
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
	fmt.Println("result: ", result)
	return result
}

func q_sql() []HeaderPostgres {
	fmt.Println("sql records")

	err := godotenv.Load("./env/env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	/*
		for _, e := range os.Environ() {
			pair := strings.SplitN(e, "=", 2)
			fmt.Printf("%s: %s\n", pair[0], pair[1])
		}
		fmt.Println("user", os.Getenv("USER"))
	*/

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

	fmt.Println("list of records")
	fmt.Println("---------------")
	var result = []HeaderPostgres{}
	for rows.Next() {
		var item HeaderPostgres
		rows.Scan(&item.Idheader, &item.Companyname, &item.Address, &item.NumberInvoice, &item.DateTime, &item.CreatedAt)
		result = append(result, item)
	}
	fmt.Println("result: ", result)
	return result
}

func q_sql_one(id int) []HeaderPostgres {
	fmt.Println("sql to psql")

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

	rows, err := db.Query("select * from header where id_header = $1", id)
	//rows, err := db.Query(`SELECT * FROM header WHERE id_header = $1`, id)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("\nRow selected successfully!")
	}
	defer rows.Close()

	var result = []HeaderPostgres{}
	for rows.Next() {
		var item HeaderPostgres
		rows.Scan(&item.Idheader, &item.Companyname, &item.Address, &item.NumberInvoice, &item.DateTime, &item.CreatedAt)
		result = append(result, item)
		fmt.Println("result: ", result)
	}
	return result
}

func create_token(dataPost []uint8) string {
	// Define the expiration time for the token
	expirationTime := time.Now().Add(24 * time.Hour)

	var InsertJson Users
	json.Unmarshal([]byte(dataPost), &InsertJson)

	user := User{
		Username: *InsertJson.Username,
		Email:    *InsertJson.Email,
	}

	// Create the claims containing the user information
	claims := jwt.MapClaims{
		//"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
		"exp":      expirationTime.Unix(),
	}

	// Create the JWT token with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Define the secret key used for signing the token
	// Note: Keep the key secure and do not hardcode it in your code
	secretKey := []byte("your-secret-key")

	// Sign the token with the secret key
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println("Error signing token:", err)
		return "error"
	}

	return signedToken

}

func insert_user_sql(dataPost []uint8) {
	fmt.Println("insert user sql with token")

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
	var InsertJson Users
	json.Unmarshal([]byte(dataPost), &InsertJson)

	// create token for user
	signedToken := create_token(dataPost)
	myTime := time.Now()

	//insert in postgres
	sqlStatement := `INSERT INTO usuario (username, email, token, created_at) VALUES ($1, $2, $3, $4)`
	_, err = db.Exec(sqlStatement, InsertJson.Username, InsertJson.Email, signedToken, myTime)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("\nRow inserted successfully!, token: ", signedToken)
	}
	return

}

func insert_header_sql(dataPost []uint8) {
	fmt.Println("insert header sql")

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
	fmt.Printf("Id Header: %v, Company Name %s \n", InsertJson.Idheader, InsertJson.Companyname)

	myTime := time.Now()
	//insert in postgres
	sqlStatement := `INSERT INTO header (companyname, address, numberinvoice, date_time) VALUES ($1, $2, $3, $4)`
	_, err = db.Exec(sqlStatement, InsertJson.Companyname, InsertJson.Address, InsertJson.NumberInvoice, myTime)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("\nRow inserted successfully!")
	}
	return
}

func update_sql(dataPost []uint8) {
	fmt.Println("update sql")

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
	var UpdateJson HeaderPostgres
	json.Unmarshal([]byte(dataPost), &UpdateJson)
	fmt.Println("type of InsertJson = ", reflect.TypeOf(UpdateJson))
	fmt.Printf("Id Header: %v, Company Name %s \n", UpdateJson.Idheader, UpdateJson.Companyname)

	//update in postgres
	rows, err := db.Query("UPDATE header SET companyname = $2 WHERE id_header = $1", UpdateJson.Idheader, UpdateJson.Companyname)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("\nRow updates successfully!")
	}
	defer rows.Close()
	return
}

func delete_sql_one(id int) {
	fmt.Println("delete sql")

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

	rows, err := db.Query("delete from header where id_header = $1", id)
	//rows, err := db.Query(`SELECT * FROM header WHERE id_header = $1`, id)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("\nRow deletes successfully!")
	}
	defer rows.Close()

	//var result = []HeaderPostgres{}
	//for rows.Next() {
	//	var item HeaderPostgres
	//	rows.Scan(&item.Idheader, &item.Companyname, &item.Address, &item.NumberInvoice, &item.DateTime, &item.CreatedAt)
	//	result = append(result, item)
	//	fmt.Println("result: ", result)
	//}
}

func insert_invoice_sql(dataPost []uint8) {
	fmt.Println("insert invoice sql")

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

	//get last id_ehader inserted
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
			fmt.Println("\nRow Detail inserted successfully!")
		}
	}
	return
}
