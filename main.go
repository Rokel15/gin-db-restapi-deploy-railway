package main

import (
	"database/sql"
	"fmt"
	"golang-gin-db-restapi-lokal/database"
	"golang-gin-db-restapi-lokal/routers"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "passwrongwkwk"
// 	dbName   = "for_testing1"
// )

var (
	DB  *sql.DB
	err error
)

func main() {
	var PORT = ":" + os.Getenv("PORT")

	// psqlInfo := fmt.Sprintf(`host=%s port=%d user=%s password=%s dbname=%s sslmode=disable`,
	// 	host,
	// 	port,
	// 	user,
	// 	password,
	// 	dbName,
	// )

	// psqlInfo := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
	// 	os.Getenv("DB_HOST"),
	// 	os.Getenv("DB_PORT"),
	// 	os.Getenv("DB_USER"),
	// 	os.Getenv("DB_PASSWORD"),
	// 	os.Getenv("DB_NAME"),
	// )

	psqlInfo := os.Getenv("DATABASE_URL")

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
