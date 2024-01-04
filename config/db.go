package config

import (
	"fmt"
	"mentalartsgo/models"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@database:5432/myapp?sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	start := time.Now()
	for sqlDB.Ping() != nil {
		if start.After(start.Add(10 * time.Second)) {
			fmt.Println("Failed to connect db after  10 seconds ")
			break
		}
	}
	fmt.Println("Connected to DB:", sqlDB.Ping() == nil)

	// DB değişkenine atama yap
	DB = db

	// AutoMigrate işlemi
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Monster{})
	db.AutoMigrate(&models.User{}, &models.Beyblade{}, &models.Monster{})

}
