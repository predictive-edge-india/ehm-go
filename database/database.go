package database

import (
	"fmt"

	"github.com/iisc/demo-go/models"
	log "github.com/sirupsen/logrus"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB
var err error

func InitDatabase() *gorm.DB {
	Database, err = gorm.Open(postgres.Open(getDSN()), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	err = Database.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`).Error
	if err != nil {
		log.Fatal("Error creating extensions", err)
	}
	Migrate(Database)
	return Database
}

func Migrate(db *gorm.DB) {
	err = db.AutoMigrate(
		&models.EhmDevice{},
		&models.CurrentParameter{},
		&models.FuelPercentage{},
		&models.DeviceFault{},
		&models.PowerParam{},
	)
	if err != nil {
		log.Fatal("Error migrating DB", err)
	}
}

func getDSN() string {
	HOST := models.Environments.DBHost
	USER := models.Environments.DBUser
	DBPASSWORD := models.Environments.DBPassword
	DBNAME := models.Environments.DBName
	PORT := models.Environments.DBPort
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", HOST, USER, DBPASSWORD, DBNAME, PORT)
}
