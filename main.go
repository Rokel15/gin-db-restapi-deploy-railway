package main

import (
	"database/sql"
	"fmt"
	"golang-gin-db-restapi-lokal/database"
	"golang-gin-db-restapi-lokal/routers"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "liatdibawahlaptop"
	dbName   = "for_testing1"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	var PORT = ":8080"

	psqlInfo := fmt.Sprintf(`host=%s port=%d user=%s password=%s dbname=%s sslmode=disable`, host, port, user, password, dbName)

	DB, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatalf("Error Opening Database: %v\n", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error pinging Database: %v\n", err)
	}

	database.DBMigrate(DB)

	defer DB.Close()

	routers.StartServer().Run(PORT)
	fmt.Println("Successfully Connected")
}
