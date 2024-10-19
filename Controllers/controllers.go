package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"netLog/db"
	"netLog/models"
	"os"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	visitors := db.QueryVisitors()
	json.NewEncoder(w).Encode(visitors)
}

func checkIpAddress(ip string) bool {
	if net.ParseIP(ip) == nil {
		log.Fatalf("IP Address: %s - Invalid\n", ip)
		return false
	} else {
		fmt.Printf("IP Address: %s - Valid\n", ip)
		return true
	}
}

// https://ipinfo.io/${ip}?token=${process.env.IPINFO_API_KEY}
func AddNewVisitor(w http.ResponseWriter, r *http.Request) {

	remoteIp := "152.59.2.94"
	if !checkIpAddress(remoteIp) {
		log.Fatal("the ip is not valid")
	}

	ipInfoKey := os.Getenv("IP_INFO_KEY")

	url := fmt.Sprintf("https://ipinfo.io/%s?token="+ipInfoKey, remoteIp)
	fmt.Println(url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error making GET request: %v", err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		log.Fatalf("Error: calling ipinfo api", response)
	}

	w.Header().Set("Content-Type", "application/json")
	var newVisitor models.Visitor
	errr := json.NewDecoder(response.Body).Decode(&newVisitor)
	fmt.Println(response.Body)
	db.InsertVisitorInfo(&newVisitor)
	if errr != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(newVisitor)
}
