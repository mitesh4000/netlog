package db

import (
	"database/sql"
	"fmt"
	"log"
	"netLog/models"

	_ "modernc.org/sqlite"
)


var DB *sql.DB

type IPData struct {
    ID       int
    IP       string
    City     string
    Region   string
    Country  string
    Loc      string
    Org      string
    Postal   string
    Timezone string
}
func InitDb(DataSource string){
	var err error
	DB,err = sql.Open("sqlite",DataSource)
	if err != nil{
		log.Fatal(err)
	}


	if err = DB.Ping(); err != nil{
		log.Fatal(err)
	}
}

func CloseDB(){
	if err := DB.Close(); err !=nil{
		log.Fatal(err)
	}
}

func InsertVisitorInfo(data *models.Visitor ) (string,error){

    insertQuery := `
    INSERT INTO visitors (ip, city, region, country, loc, org, postal, timezone)
    VALUES (?, ?, ?, ?, ?, ?, ?, ?);
    `
    _, err := DB.Exec(insertQuery, data.IP, data.City, data.Region,data.Country, data.Loc, data.Org, data.Postal,data.Timezone)
    if err != nil {
        return err.Error(), fmt.Errorf("failed to insert visitor info: %v", err)
    }
    return "success",nil
}


func QueryVisitors() []IPData{

    rows, err := DB.Query("SELECT id, ip, city, region, country, loc, org, postal, timezone FROM visitors")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    var allIpData []IPData
    for rows.Next() {
        var  ipData IPData
        
        err := rows.Scan(&ipData.ID, &ipData.IP, &ipData.City, &ipData.Region, &ipData.Country, &ipData.Loc, &ipData.Org, &ipData.Postal, &ipData.Timezone)
        if err != nil {
            log.Fatal(err)
        }
        allIpData = append(allIpData, ipData)
    }
        return allIpData;
}

func GetTotalVisitors() (int, error) {
    var totalVisitors int

    err := DB.QueryRow("SELECT COUNT(*) FROM visitors").Scan(&totalVisitors)
    log.Printf("Total %d visitors", totalVisitors)
    if err != nil {
        return 0, err
    }

    return totalVisitors, nil
}
