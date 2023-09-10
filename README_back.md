# api in go and angular 15
doc: https://go.dev/doc/database/querying#multiple_rows

## create folder and modules
* $ go mod init packages
* $ go mod tidy

[...]

# create and add modules to a workspace
path: backend
- go clean -modcache

- go work init api
- go work use ./go_api_postgres/backend


# compile compile file
* $ go work build -> just for my workspace
* $ go build *.go
* $ go run *.go

* CREATE DATABASE invoices WITH OWNER test;

## package
* $ go get -u github.com/lib/pq
* $ go get -u github.com/gorilla/mux
* $ go get github.com/joho/godotenv
* $ go get github.com/rs/cors
* $ go get -u github.com/go-swagger/go-swagger/cmd/swagger
* $ go get -u github.com/golang-jwt/jwt/v5
* $ go get github.com/sirupsen/logrus

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

* method: post
http://localhost:8080/createOneInvoice

-Post header:
{
    "companyname": "test2",
    "address": "stqw 12ole",
    "numberinvoice": 234,
    "datetime": "2023-04-30T15:14:06Z",
    "createdat": "2023-04-30T15:14:06Z"
}

- POST http://localhost:8080/createOneHeader

{
    "idheader": "14",
    "companyname": "jga",
    "price": 32,
    "author": null
}

-Post Invoice:
{
    "companyname": "nananan",
    "address": "stqw 1212 ole",
    "numberinvoice": 342,
    "datetime": "2023-04-30T15:14:06Z",
    "createdat": "2023-01-02T15:04:05Z",
    "iddetail": [{
            "detailid": 0,
            "idheader": 0,
            "description": "description1_detail1",
            "units": 2,
            "price": 12.4,
            "createdat": "2023-12-02T15:04:05Z"
        },
        {
            "detailid": 0,
            "idheader": 0,
            "description": "description2_detail1",
            "units": 21,
            "price": 212.4,
            "createdat": "2023-08-22T15:04:05Z"
        }
    ]
}

* method post
http://localhost:8080/createOneUser

{
    #"username": "jjg",
    "password": "1234",
    "email": "email@gmail.com"
}

* method post
http://localhost:8080/users/login

{
    "email": "email@gmail.com",
    "password": "1234"
}
