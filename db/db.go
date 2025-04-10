package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func InitDB(dbHost, dbPort, dbUser, dbPassword, dbName string) *sql.DB {
	conn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)

	DB, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}
	if err := DB.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Database connection established")
	return DB
}

func DeleteExpiredRecords(db *sql.DB) {

	curentTime := time.Now()

	query := `
	DELETE FROM urls
 WHERE ttl_seconds IS NOT NULL
   AND created_at + INTERVAL '1 second' * ttl_seconds < $1;

	`
	_, err := db.Exec(query, curentTime)
	if err != nil {
		log.Printf("Error deleting expired records: %v", err)
	} else {
		log.Printf("Expired records deleted successfully")
	}
}
