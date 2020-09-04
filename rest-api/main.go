package main

import (
	"log"
	"net/http"
	"os"

	routing "rest/routes"

	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func main() {

	loadEnv()

	PORT := os.Getenv("PORT")

	router := routing.NewRouter()

	log.Fatal(http.ListenAndServe(":"+PORT, router))
}
