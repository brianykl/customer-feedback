package models

import (
	"log"
	// "testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := "host=localhost user=postgres password=polarBear$02 dbname=feedback port=5432 sslmode=disable TimeZone=America/New_York"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	log.Println("connected to database successfully")
	return db
}

func Insert(feedback *Feedback) error {
	db := Connect()
	result := db.Create(feedback)
	if result.Error != nil {
		log.Println("failed to insert feedback %v", result.Error)
		return result.Error
	}

	log.Println("feedback inserted successfully")
	return nil
}

func Execute(query string, args ...interface{}) error {
	db := Connect()

	// Execute the SQL statement
	result := db.Exec(query, args...)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
