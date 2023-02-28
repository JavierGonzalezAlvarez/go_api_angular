# api in go and angular 15

## crear file
* $ go mod init psql
* $ go build connection.go api.go // $ go build *.go
* $ go run connection.go api.go // $ go run *.go

* CREATE DATABASE invoices WITH OWNER test;

* $ go get -u github.com/lib/pq
* $ go get -u github.com/gorilla/mux
* $ go get github.com/joho/godotenv