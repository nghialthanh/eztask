package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	Dbdriver := os.Getenv("DB_DRIVER")
	// DBURL := os.Getenv("POSTGRES_SOURCE")
	DbHost := os.Getenv("DB_HOST")
	DbPort := os.Getenv("DB_PORT")
	DbUser := os.Getenv("DB_USER")
	DbName := os.Getenv("DB_NAME")
	DbPassword := os.Getenv("DB_PASSWORD")
	// if Dbdriver == "postgres" {
	_ = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	POSTGRES_SOURCE := os.Getenv("POSTGRES_SOURCE")
	db, err := gorm.Open(postgres.Open(POSTGRES_SOURCE), &gorm.Config{})
	if err != nil {
		fmt.Printf("Cannot connect to %s database", Dbdriver)
		log.Fatal("This is the error:", err)
	}

	fmt.Println("Connected!")

	return db
}
