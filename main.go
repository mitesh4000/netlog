package main

import (
	"log"
	"net/http"
	"netLog/db"
	"netLog/routes"
)



func main() {
	
	db.InitDb("./netlog.db")

	routes.SetupRoutes()

	defer db.CloseDB()

	log.Println("Server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))

}
