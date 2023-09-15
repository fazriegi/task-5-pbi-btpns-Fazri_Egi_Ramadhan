package database

import (
	"fmt"
	"log"
	"os"
	"task-5-pbi-btpns-Fazri_Egi_Ramadhan/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Open() error {
	dbUsername := os.Getenv("DBUSERNAME")
	dbPassword := os.Getenv("DBPASSWORD")
	dbHost := os.Getenv("DBHOST")
	dbPort := os.Getenv("DBPORT")
	dbName := os.Getenv("DBNAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUsername, dbPassword, dbHost, dbPort, dbName)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return fmt.Errorf("Failed to connect to database: %w", err)
	}

	log.Println("Database connected...")

	InitMigrate()

	return nil
}

func InitMigrate() {
	DB.AutoMigrate(&models.User{}, &models.Photo{})
}
