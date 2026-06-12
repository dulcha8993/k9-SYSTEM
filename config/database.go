package config

import (
	"log"
	"fmt"
	"os"

	"gorm.io/driver/postgres" //PostgreSQL driver for GORM (bridge between postgres and GORM)
	"gorm.io/gorm" // DB related functions (gorm)
)

var DB *gorm.DB //lobal database connection object (*gorm.DB - pointer to GORM database instance.)

func ConnectDatabase() {

	host := os.Getenv("DB_HOST");
	user := os.Getenv("DB_USER");
	password := os.Getenv("DB_PASSWORD");
	dbname := os.Getenv("DB_NAME");
	port := os.Getenv("DB_PORT");

	log.Println("DB_NAME:", dbname);

		//Database Connection String
	dsn := fmt.Sprintf(
	"host=%s user=%s dbname=%s port=%s sslmode=disable",
	host,
	user,
	dbname,
	port,
)

if password != "" {
	dsn = fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host,
		user,
		password,
		dbname,
		port,
	)
}
log.Println(dsn)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect database")
	}

	DB = database

	log.Println("Database connected successfully")


	sqlDB, _ := DB.DB()

var currentDB string

sqlDB.QueryRow("SELECT current_database()").Scan(&currentDB)

log.Println("Connected DB:", currentDB)
}