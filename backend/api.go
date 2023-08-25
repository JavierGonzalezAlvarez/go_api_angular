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
	Iduser   int    `json:"iduser"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	//Token     string    `json:"token"`
	CreatedAt time.Time `json:"createdat"`
}

type Header struct {
	Idheader      int       `json:"idheader"`
	Companyname   string    `json:"companyname"`
	Address       string    `json:"address"`
	NumberInvoice int       `json:"numberinvoice"`
	DateTime      time.Time `json:"datetime" format:"2006-01-02 15:04:05"`
	CreatedAt     time.Time `json:"createdat"`
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

type Response struct {
	Token      string    `json:"token"`
	Expiracion time.Time `json:"expiracion"`
	User       string    `json:"username"`
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
	fmt.Println("back api")
	fmt.Println("running on http://localhost:8080/")
	fmt.Println("running on http://localhost:8080/get")
	fmt.Println("running on http://localhost:8080/getOne/1")
	fmt.Println("running on http://localhost:8080/createOneHeader")
	fmt.Println("running on http://localhost:8080/updateOne/1")
	fmt.Println("running on http://localhost:8080/deleteOne/1")
	fmt.Println("running on http://localhost:8080/createOneInvoice")

	fmt.Println("running on http://localhost:8080/getUsers")
	fmt.Println("running on http://localhost:8080/createOneUser")

	fmt.Println("running on http://localhost:8080/users/login")

	router := mux.NewRouter()

	router.HandleFunc("/", home).Methods("GET")
	router.HandleFunc("/get", getAllRecords).Methods("GET")
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
	router.HandleFunc("/getOne/{id}", getOneRecord).Methods("GET")
	router.HandleFunc("/createOneHeader", createOneRecordHeader).Methods("POST")
	router.HandleFunc("/updateOne/{id}", updateOneRecord).Methods("PUT")
	router.HandleFunc("/deleteOne/{id}", deleteOneRecord).Methods("DELETE")
	router.HandleFunc("/createOneInvoice", createOneRecordInvoice).Methods("POST")

	router.HandleFunc("/getUsers", getAllUsers).Methods("GET")
	router.HandleFunc("/createOneUser", createOneUser).Methods("POST")

	router.HandleFunc("/users/login", postUserLogin).Methods("POST")

	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func home(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("<h1>api go & postgres & angular 15</h1>"))
}

func postUserLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getting login from Angular, return token and expiration token")
	// Parse request body
	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	//TODO: get credentials from DB

	// Create a response object
	response := Response{
		Token:      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImVtYWlsQGdtYWlsLmNvbSIsImV4cCI6MTY5MDYzOTE1MCwidXNlcm5hbWUiOiJqamcifQ.tK-WL-EPsC2L-NQR9L_-TX29liERv4l2h5M2r4HUpCk",
		Expiracion: time.Now(),
		User:       "jjg",
	}

	// Convert the response object to JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		return
	}

	// set the content type
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// respond with JSON
	w.Write(jsonResponse)
}

func getAllUsers(w http.ResponseWriter, _ *http.Request) {
	fmt.Println("get all users")
	var data = get_all_users()
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(data)
}

var users = []User{}

func createOneUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create one user with token")
	var data = get_all_users()
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(data)

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
		json.NewEncoder(w).Encode("Email already exists")
		fmt.Println("Email already exists:", user.Email)
		return
	}

	users = append(users, user)
	//response (200)
	json.NewEncoder(w).Encode(user)
	//print json
	finalJson, _ := json.MarshalIndent(user, "", "\t")
	fmt.Println("final response json indented", string(finalJson))
	fmt.Println("type of json = ", reflect.TypeOf(finalJson))

	insert_user_sql(finalJson)
	return
}

func getAllRecords(w http.ResponseWriter, _ *http.Request) {
	fmt.Println("get all records")
	// retrieve data from postgres
	var data = q_sql()
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func getOneRecord(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get one record")
	w.Header().Set("content-type", "application/json")
	params_value := mux.Vars(r)
	fmt.Println("params in url", params_value)

	id_params := mux.Vars(r)["id"]
	fmt.Println("value of params id from url: ", id_params)
	fmt.Println("type of id params = ", reflect.TypeOf(id_params))
	//convert to int
	intVar, _ := strconv.Atoi(id_params)
	var data = q_sql_one(intVar)
	fmt.Println("type of id params = ", reflect.TypeOf(intVar))

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
	fmt.Println("create a new record header")
	w.Header().Set("content-type", "application/json")

	//check if data is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Pls send some data or validate data sent")
	}
	println(r.Body)
	var header Header
	_ = json.NewDecoder(r.Body).Decode(&header)
	if header.IsEmptyCompanyName() {
		json.NewEncoder(w).Encode("Pls revise data company name is empty")
		return
	}

	headers = append(headers, header)
	//response (200)
	json.NewEncoder(w).Encode(header)
	//print json
	finalJson, _ := json.MarshalIndent(header, "", "\t")
	fmt.Println("final response json indented", string(finalJson))
	fmt.Println("type of json = ", reflect.TypeOf(finalJson))

	insert_header_sql(finalJson)
	return
}

func updateOneRecord(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update one record")
	w.Header().Set("content-type", "application/json")

	// check if required data is empty
	var headerRequire Header
	_ = json.NewDecoder(r.Body).Decode(&headerRequire)
	if headerRequire.IsEmptyCompanyName() {
		json.NewEncoder(w).Encode("Pls revise copany name is empty")
		return
	}

	//get id
	params_value := mux.Vars(r)
	fmt.Println("params in url", params_value)
	id_params := mux.Vars(r)["id"]
	fmt.Println("value of params id from url: ", id_params)
	intVar, _ := strconv.Atoi(id_params)
	var data = q_sql_one(intVar)

	for _, header := range data {
		if header.Idheader == intVar {
			fmt.Println("record exist")

			//json of old data
			var UpdateJson HeaderPostgres
			finalJson, _ := json.MarshalIndent(header, "", "\t")
			json.Unmarshal([]byte(finalJson), &UpdateJson)
			fmt.Printf("Id Header: %v, Company Name %s \n", UpdateJson.Idheader, UpdateJson.Companyname)

			//json to struct
			//header update
			//UpdateJson.Idheader = final.IdHeader
			//UpdateJson.Companyname

			//new struct to json

			// update to be done, required
			headers = append(headers, headerRequire)
			json.NewEncoder(w).Encode(headerRequire)
			final, _ := json.MarshalIndent(headerRequire, "", "\t")
			fmt.Println("record to be updated", string(final))

			update_sql(final)
			return
		}
	}
	json.NewEncoder(w).Encode("No record found by this id")
	return
}

func deleteOneRecord(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete one record")
	w.Header().Set("content-type", "application/json")

	//get id
	params_value := mux.Vars(r)
	fmt.Println("params in url", params_value)
	id_params := mux.Vars(r)["id"]
	fmt.Println("value of params id from url: ", id_params)
	intVar, _ := strconv.Atoi(id_params)
	var data = q_sql_one(intVar)

	for _, header := range data {
		if header.Idheader == intVar {
			fmt.Println("record exist")
			delete_sql_one(intVar)
			return
		}
	}
	json.NewEncoder(w).Encode("No record found by this id")
	return
}

var invoices = []Invoice{}

func createOneRecordInvoice(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create a new invoice")
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
		fmt.Fprintf(w, "Received body: %s", string(body))
		fmt.Println("response body:", string(body))
		insert_invoice_sql(body)
		return
	} else {
		http.Error(w, "Error, empty data", http.StatusBadRequest)
	}
}
