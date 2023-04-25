# api in go and angular 15
doc: https://go.dev/doc/database/querying#multiple_rows

## create/compile file
* $ go mod init psql
* $ go build connection.go api.go // $ go build *.go
* $ go run connection.go api.go // $ go run *.go

* CREATE DATABASE invoices WITH OWNER test;

## package
* $ go get -u github.com/lib/pq
* $ go get -u github.com/gorilla/mux
* $ go get github.com/joho/godotenv
* $ go get github.com/rs/cors
* $ go get -u github.com/go-swagger/go-swagger/cmd/swagger

## docs
https://goswagger.io/
https://swagger.io/docs/specification/describing-parameters/

## create .yaml
$ swagger generate spec --scan-models --output=./swagger.yaml

## run swagger
$ swagger serve -F=swagger swagger.yaml
http://localhost:38757/docs#/

## endpoints
* method: get all
http://localhost:4000/get

* method: get (one)
http://localhost:4000/getOne/1

* method: create (one)
http://localhost:4000/createOne

* method: put
http://localhost:4000/updateOne/1

* method: delete
http://localhost:4000/deleteOne/6

