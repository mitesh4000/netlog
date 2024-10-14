package main

import (
	"log"
	"net/http"
	"netLog/db"
	"netLog/routes"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db.InitDb("./netlog.db")

	routes.SetupRoutes()

	defer db.CloseDB()

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
