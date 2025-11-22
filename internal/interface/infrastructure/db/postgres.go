package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	if err := godotenv.Load(".env.development"); err != nil {
		log.Println("Warning: .env.development file not found, using environment variables")
	}
}

func ConnectDB() *gorm.DB {
	dns := GetDNS()

	log.Printf("Connecting to db at - %s", dns)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	return db
}

func GetDNS() string {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DBNAME")
	sslMode := os.Getenv("POSTGRES_SSLMODE")

	return "host=" + host + " port=" + port + " user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=" + sslMode
}
