package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Header struct {
	Idheader      int       `json:"headerid"`
	Companyname   string    `json:"companyname"`
	Address       string    `json:"address"`
	NumberInvoice int       `json:"numberinvoice"`
	DateTime      time.Time `json:"datetime"`
	CreatedAt     time.Time `json:"createdat"`
}

type Detail struct {
	IdDetail    int       `json:"detailid"`
	IdHeader    *Header   `json:"idheader"`
	Description string    `json:"description"`
	Units       int       `json:"units"`
	Price       float32   `json:"price"`
	CreatedAt   time.Time `json:"createdat"`
}

func main() {
	fmt.Println("api")

	router := mux.NewRouter()
	router.HandleFunc("/", home).Methods("GET")
	router.HandleFunc("/get", getAllRecords).Methods("GET")
	router.HandleFunc("/getOne/{id}", getOneRecord).Methods("GET")

	//http://localhost:4000/
	fmt.Println("running on http://localhost:4000/")
	fmt.Println("running on http://localhost:4000/get")
	fmt.Println("running on http://localhost:4000/getOne/1")

	log.Fatal(http.ListenAndServe(":4000", router))
}

// retrieve data from postgres
var data = postgres()

func home(w http.ResponseWriter, router *http.Request) {
	w.Write([]byte("<h1>api postgres</h1>"))
}

func getAllRecords(w http.ResponseWriter, router *http.Request) {
	fmt.Println("get all records")
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func getOneRecord(w http.ResponseWriter, router *http.Request) {
	fmt.Println("get one record")
	w.Header().Set("content-type", "application/json")

	params_value := mux.Vars(router)
	fmt.Println("params in url", params_value)

	id_params := mux.Vars(router)["id"]
	fmt.Println("value of params id from url: ", id_params)
	fmt.Println("type of id params = ", reflect.TypeOf(id_params))
	//convert to int
	intVar, _ := strconv.Atoi(id_params)
	//intVar, _ := strconv.ParseInt(params, 0, 8) // convert to int64
	fmt.Println("type of id params = ", reflect.TypeOf(intVar))

	for _, header := range data {
		if header.Idheader == intVar {
			json.NewEncoder(w).Encode(header)
			return
		}
	}

	json.NewEncoder(w).Encode("No record found by this id")
	return
}
