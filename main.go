package main

import (
	"log"
	"net/http"
	"netLog/db"
	"netLog/routes"
	"os"

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

	port := os.Getenv("PORT")
	log.Println("Server started on " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
