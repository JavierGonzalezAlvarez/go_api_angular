package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Header struct {
	Idheader      int       `json:"courseid"`
	Companyname   string    `json:"companyname"`
	Address       string    `json:"address"`
	NumberInvoice int       `json:"numberinvoice"`
	DateTime      time.Time `json:"datetime"`
	CreatedAt     time.Time `json:"createdat"`
}

type Detail struct {
	IdDetail    int       `json:"courseid"`
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

	//http://localhost:4000/
	fmt.Println("running on http://localhost:4000/")
	fmt.Println("running on http://localhost:4000/get")

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
