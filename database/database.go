package database

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"washington/go_api/config"
	"washington/go_api/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Database instance struct (type)
type Dbinstance struct {
	Db *gorm.DB
}

// Database instance
var DB Dbinstance

// Function to connect to the database
func Connect() {
	p := config.Config("DB_PORT")
	// Because our config function returns a string,
	// we are parsing our str to int here
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		fmt.Println("Error parsing str to int")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Australia/Sydney", config.Config("DB_HOST"), config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"), port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("Connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")
	// Run auto gorm migration
	db.AutoMigrate(&model.User{}, &model.Book{}, &model.Author{})
	if err != nil {
		log.Fatal("Failed to migrate. \n", err)
	}

	DB = Dbinstance{
		Db: db,
	}
}
