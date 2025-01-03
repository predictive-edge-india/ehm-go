package database

import (
	"fmt"

	"github.com/predictive-edge-india/ehm-go/models"
	"github.com/rs/zerolog/log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Database *gorm.DB
var err error

func InitDatabase() *gorm.DB {
	Database, err = gorm.Open(postgres.Open(getDSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		log.Fatal().AnErr("DB connection", err).Send()
	}

	if err = Database.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`).Error; err != nil {
		log.Fatal().AnErr("Creating UUID extension", err).Send()
	}

	if err = Database.Exec("CREATE EXTENSION IF NOT EXISTS postgis;").Error; err != nil {
		log.Error().AnErr("Creating postgis extension", err).Send()
	}
	Migrate(Database)
	return Database
}

func Migrate(db *gorm.DB) {
	err = db.AutoMigrate(
		&models.Customer{},
		&models.User{},
		&models.UserRole{},
		&models.AssetClass{},
		&models.Asset{},
		&models.DeviceType{},
		&models.Device{},
		&models.AssetDevice{},
		&models.AssetParameter{},
		&models.DeviceLastLocation{},
		&models.AlarmStatusFlag{},
		&models.DGStatus{},
		&models.PowerData{},
		&models.E483CanData{},
		&models.PowerLimit{},
	)
	if err != nil {
		log.Fatal().AnErr("Migrating DB", err).Send()
	}

	if db.Migrator().HasTable(&models.User{}) {
		seedData(db)
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

func seedData(db *gorm.DB) {
	// user seed
	newUser, err := SeedUser(db)
	if err != nil {
		log.Error().AnErr("SeedUser", err).Send()
	}

	// customer seed
	customer, err := SeedCustomer(db)
	if err != nil {
		log.Error().AnErr("SeedCustomer", err).Send()
	}

	// user role seed
	_, err = SeedUserRole(db, newUser, customer)
	if err != nil {
		log.Error().AnErr("SeedUserRole", err).Send()
	}
}
