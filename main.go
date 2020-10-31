package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"bookstore/routes"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	router := routes.Routes()
	fmt.Println("Running server on port " + os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
