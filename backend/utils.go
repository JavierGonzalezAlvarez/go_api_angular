package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func find_email(email string) string {
	fmt.Println("looking for an email in db  :", email)
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

	rows, err := db.Query("select email from usuario where email = $1", email)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("\nRow (email) selected successfully!")
	}
	defer rows.Close()

	var result_email string
	for rows.Next() {
		var item Users
		rows.Scan(&item.Email)
		result_email = *item.Email
	}
	return result_email

}

func get_credentials(email string, password string) []Response {
	fmt.Println("looking for an email in db  :", email)
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

	rows, err := db.Query("select username, token from usuario where email = $1 and password = $2", email, password)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("\nRow (email - password) selected successfully!")
	}
	defer rows.Close()

	var result = []Response{}
	for rows.Next() {
		var item Response
		rows.Scan(&item.User, &item.Token)
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
	secretKey := []byte("ATTRSAOis98Aha87sNHY48725s45dLOQIJS")

	// Sign the token with the secret key
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println("Error signing token:", err)
		return "error"
	}

	return signedToken

}
