package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type User struct {
	Iduser    int       `json:"iduser"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdat"`
	Role      string    `json:"role"`
}

type Header struct {
	Idheader      int       `json:"idheader"`
	Companyname   string    `json:"companyname"`
	Address       string    `json:"address"`
	NumberInvoice int       `json:"numberinvoice"`
	DateTime      time.Time `json:"datetime" format:"2006-01-02 15:04:05"`
	CreatedAt     time.Time `json:"createdat" format:"2006-01-02 15:04:05"`
}

type Detail struct {
	IdDetail    int       `json:"detailid"`
	IdHeader    int       `json:"idheader"`
	Description string    `json:"description"`
	Units       int       `json:"units"`
	Price       float32   `json:"price"`
	CreatedAt   time.Time `json:"createdat"`
}

type Invoice struct {
	Header
	Iddetail []Detail `json:"iddetail" `
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// middleware, validation
func (h *Header) IsEmptyCompanyName() bool {
	return h.Companyname == ""
}

func (h *Invoice) IsEmptyCompanyName() bool {
	return h.Companyname == ""
}

func (h *User) IsEmptyEmail() bool {
	return h.Email == ""
}

func main() {

	// Log messages
	/*
		Logger.WithFields(logrus.Fields{
			"key1": "value1",
			"key2": "value2",
		}).Info("This is an info log message.")

		Logger.Info("This is an info message.")
		Logger.Warn("This is a warning message.")
		Logger.Error("This is an error message.")

	*/

	router := mux.NewRouter()

	// swagger:route GET /getOne/{id} getOneRecord
	//
	// It returns one record.
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https, ws, wss
	//
	//     Parameters:
	//       + name: id
	//         in: path
	//         description: returns one record
	//         required: true
	//         type: integer
	//         format: int
	//
	//     Responses:
	//       200:
	//			description: OK
	//		 	  headers:
	//				Idheader:
	//				type: int
	//				description: "Data Header table"
	//       422:
	//				description: "Error 422"

	routes := []RouteDefinition{
		{Path: "/", Handler: home, Method: "GET"},
		{Path: "/get_all_header_invoices", Handler: getAllHeadersInvoices, Method: "GET"},
		{Path: "/get_all_header_invoices_total", Handler: getAllHeadersInvoicesTotal, Method: "GET"},
		{Path: "/get_one_header_invoice/{id}", Handler: getOneHeaderInvoice, Method: "GET"},
		{Path: "/createOneHeader", Handler: createOneRecordHeader, Method: "POST"},
		{Path: "/updateOne/{id}", Handler: updateOneRecord, Method: "PUT"},
		{Path: "/deleteOne/{id}", Handler: deleteOneRecord, Method: "DELETE"},
		{Path: "/createOneInvoice", Handler: createOneRecordInvoice, Method: "POST"},
		{Path: "/getUsers", Handler: getAllUsers, Method: "GET"},
		{Path: "/users/createOneUser", Handler: createOneUser, Method: "POST"},
		{Path: "/users/login", Handler: postUserLogin, Method: "POST"},
	}

	addRoutes(router, routes)

	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func home(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("<h1>api go & postgres & angular 15</h1>"))
}

func postUserLogin(w http.ResponseWriter, r *http.Request) {

	Logger.Info("login from Angular, return token and expiration token")

	// Parse request body
	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	//check in DB credentials
	credentials := get_credentials(creds.Email, creds.Password)
	if len(credentials) == 0 {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Create a response object
	response := credentials[0]
	responses := Response{
		Token:      response.Token,
		Expiracion: time.Now(),
		User:       response.User,
		Role:       response.Role,
	}

	// Convert the response object to JSON
	jsonResponse, err := json.Marshal(responses)
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		return
	}

	// set the content type
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// respond with JSON format
	w.Write(jsonResponse)
}

func getAllUsers(w http.ResponseWriter, _ *http.Request) {
	Logger.Info("get all users")

	var data = get_all_users()
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(data)
}

var users = []User{}

func createOneUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create one user with token")
	//var data = get_all_users()
	//w.Header().Set("content-type", "application/json")
	//json.NewEncoder(w).Encode(data)

	//check if data is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Pls send some data or validate data sent")
		return
	}
	println(r.Body)

	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	if user.IsEmptyEmail() {
		json.NewEncoder(w).Encode("Pls revise data, email is empty")
		return
	}

	//check if email exists already, search email in db
	email := find_email(user.Email)

	//get email from data
	if user.Email == email {
		// return the error to the front
		http.Error(w, "Email already exists", http.StatusNotFound)
		fmt.Println("Email already exists:", user.Email)
		return
	}
	// if response is ok and email dosn't exist
	w.WriteHeader(http.StatusOK)

	users = append(users, user)

	// Create JSON response
	finalJson, _ := json.MarshalIndent(user, "", "\t")
	jsonResponse := create_user_sql(finalJson)
	jsonResponseBytes, err := json.Marshal(jsonResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set content type and status code
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Write the JSON response to the client
	if _, err := w.Write(jsonResponseBytes); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getAllHeadersInvoices(w http.ResponseWriter, _ *http.Request) {
	Logger.Info("Get all records")

	// retrieve data from postgres
	var data = get_all_headers_invoices()
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func getAllHeadersInvoicesTotal(w http.ResponseWriter, _ *http.Request) {
	Logger.Info("Get all records with total")

	// retrieve data from postgres
	var data = get_all_headers_invoices_total()
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func getOneHeaderInvoice(w http.ResponseWriter, r *http.Request) {
	Logger.Info("Get one header invoice")

	w.Header().Set("content-type", "application/json")
	params_value := mux.Vars(r)
	fmt.Println("params in url", params_value)

	id_params := mux.Vars(r)["id"]
	Logger.Info("value of params id from url: ", id_params)
	Logger.Info("type of id params = ", reflect.TypeOf(id_params))
	//convert to int
	intVar, _ := strconv.Atoi(id_params)

	// get header in db
	var data = get_one_header_invoice(intVar)
	Logger.Info("type of id params = ", reflect.TypeOf(intVar))

	for _, header := range data {
		if header.Idheader == intVar {
			json.NewEncoder(w).Encode(header)
			return
		}
	}
	json.NewEncoder(w).Encode("No record found by this id")
}

var headers = []Header{}

func createOneRecordHeader(w http.ResponseWriter, r *http.Request) {
	Logger.Info("Create a new record header")
	w.Header().Set("content-type", "application/json")

	//check if data is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Pls send some data or validate data sent")
	}
	Logger.Info(r.Body)

	var header Header
	_ = json.NewDecoder(r.Body).Decode(&header)
	if header.IsEmptyCompanyName() {
		json.NewEncoder(w).Encode("Pls revise data company name is empty")
		return
	}

	headers = append(headers, header)

	// response (200)
	json.NewEncoder(w).Encode(header)

	// print json
	finalJson, _ := json.MarshalIndent(header, "", "\t")

	fmt.Println("final response json indented", string(finalJson))
	Logger.Info("type of json = ", reflect.TypeOf(finalJson))

	create_header_invoice(finalJson)
	return
}

func updateOneRecord(w http.ResponseWriter, r *http.Request) {
	Logger.Info("Update one record")
	w.Header().Set("content-type", "application/json")

	// check if required data is empty
	var headerRequire Header
	_ = json.NewDecoder(r.Body).Decode(&headerRequire)
	if headerRequire.IsEmptyCompanyName() {
		json.NewEncoder(w).Encode("Pls revise copany name is empty")
		return
	}

	// get id
	params_value := mux.Vars(r)
	Logger.Info("params in url", params_value)

	id_params := mux.Vars(r)["id"]
	fmt.Println("value of params id from url: ", id_params)

	intVar, _ := strconv.Atoi(id_params)
	var data = get_one_header_invoice(intVar)

	for _, header := range data {
		if header.Idheader == intVar {
			fmt.Println("record exist")

			//json of old data
			var UpdateJson HeaderPostgres
			finalJson, _ := json.MarshalIndent(header, "", "\t")
			json.Unmarshal([]byte(finalJson), &UpdateJson)
			fmt.Printf("Id Header: %v, Company Name %s \n", UpdateJson.Idheader, *UpdateJson.Companyname)

			//json to struct
			//header update
			//UpdateJson.Idheader = final.IdHeader
			//UpdateJson.Companyname

			//new struct to json

			// update to be done, required
			headers = append(headers, headerRequire)
			json.NewEncoder(w).Encode(headerRequire)
			final, _ := json.MarshalIndent(headerRequire, "", "\t")

			Logger.Info("record to be updated", string(final))

			update_one_header(final)
			return
		}
	}

	// response endpoint
	json.NewEncoder(w).Encode("No record found by this id")
	return
}

func deleteOneRecord(w http.ResponseWriter, r *http.Request) {
	Logger.Info("Delete one header")
	w.Header().Set("content-type", "application/json")

	//get id
	params_value := mux.Vars(r)
	fmt.Println("params in url", params_value)
	id_params := mux.Vars(r)["id"]
	fmt.Println("value of params id from url: ", id_params)
	intVar, _ := strconv.Atoi(id_params)

	var data = get_one_header_invoice(intVar)

	for _, header := range data {
		if header.Idheader == intVar {
			fmt.Println("record exist")
			delete_one_header_invoice(intVar)
			return
		}
	}
	json.NewEncoder(w).Encode("No record found by this id")
	return
}

var invoices = []Invoice{}

func createOneRecordInvoice(w http.ResponseWriter, r *http.Request) {
	Logger.Info("Create a new invoice")
	w.Header().Set("content-type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		fmt.Fprintf(w, "Received body: %s", string(body))
		return
	}
	defer r.Body.Close()

	if json.Valid(body) {
		println("created a valid json")
		Logger.Info(w, "Received body: %s", string(body))
		Logger.Info("response body:", string(body))

		create_one_invoice(body)
		return
	} else {
		http.Error(w, "Error, empty data", http.StatusBadRequest)
	}
}
