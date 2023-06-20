package database

import (
	"GolangApiFiber/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb() {
	dsn := fmt.Sprintf(
		"host=localhost " +
			"user=test " +
			"password=test " +
			"dbname=fiber " +
			"port=5432 " +
			"sslmode=disable " +
			"TimeZone=Asia/Shanghai",
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migrations")
	db.AutoMigrate(&models.Fact{})
	db.AutoMigrate(&models.People{})

	DB = Dbinstance{
		Db: db,
	}
}
