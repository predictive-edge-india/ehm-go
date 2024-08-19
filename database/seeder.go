package database

import (
	"database/sql"
	"errors"

	"github.com/predictive-edge-india/ehm-go/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedUser(db *gorm.DB) (*models.User, error) {
	newUser := new(models.User)
	if err := db.First(&models.User{}).Scan(&newUser).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		newUser.Name = "Raunak Kamat"
		newUser.Email = "raunak@predictiveedge.co.in"
		newUser.Phone = sql.NullString{
			Valid:  true,
			String: "+919011014293",
		}
		newUser.ProfilePic = sql.NullString{
			Valid:  true,
			String: "https://i.pravatar.cc/150?img=13",
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("zzpow1614"), 8)
		if err != nil {
			return newUser, err
		}
		newUser.PasswordHash = string(hashedPassword)

		err = db.Create(&newUser).Error
		if err != nil {
			return newUser, err
		}
	}
	return newUser, nil
}

func SeedCustomer(db *gorm.DB) (*models.Customer, error) {
	newCustomer := new(models.Customer)
	if err := db.Select("id, name").First(&models.Customer{}).Scan(&newCustomer).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		newCustomer.Name = "Pai Kane Group"
		newCustomer.Phone = "+919637078081"
		newCustomer.Email = "paikane@gmail.com"
		newCustomer.LogoUrl = sql.NullString{
			Valid:  true,
			String: "https://i.pravatar.cc/150?img=10",
		}
		newCustomer.Position = models.GeoJson{
			Type:        "Point",
			Coordinates: []float64{72.8318, 19.1254},
		}
		newCustomer.Address1 = "Pai Lane"
		newCustomer.Address2 = sql.NullString{
			Valid:  true,
			String: "Pai Lane",
		}
		newCustomer.City = "Mumbai"
		newCustomer.State = "Maharashtra"
		newCustomer.Country = "India"
		newCustomer.PostalCode = 400001

		err = db.Create(&newCustomer).Error
		if err != nil {
			return newCustomer, err
		}
	}
	return newCustomer, nil
}

func SeedUserRole(db *gorm.DB, user *models.User, organisation *models.Customer) (*models.UserRole, error) {
	newUserRole := new(models.UserRole)
	if err := db.First(&models.UserRole{}).Scan(&newUserRole).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		newUserRole.UserId = user.Id
		newUserRole.CustomerId = nil
		newUserRole.AccessType = models.UserRoleEnum.SuperAdministrator.Number
		err = db.Create(&newUserRole).Error
		if err != nil {
			return newUserRole, err
		}
	}
	return newUserRole, nil
}
