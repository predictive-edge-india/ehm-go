package database

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/predictive-edge-india/ehm-go/helpers"
	"github.com/predictive-edge-india/ehm-go/models"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func FindUserById(id uuid.UUID) models.User {
	var user models.User
	Database.First(&user, "id = ?", id)
	return user
}

func FindUserByIdWithCustomers(id uuid.UUID) models.User {
	var user models.User
	Database.Preload("UserRoles").First(&user, "id = ?", id)
	return user
}

func FindUserAuth(ctx *fiber.Ctx) models.User {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	idString := claims["id"].(string)
	id, _ := uuid.Parse(idString)
	return FindUserByIdWithCustomers(id)
}

func FindUserByEmail(email string) (models.User, error) {
	var user models.User
	if err := Database.Where("email=?", email).First(&user).Error; err != nil {
		return user, err
	}

	if user.IsIdNull() {
		return user, errors.New("record not found")
	}
	return user, nil
}

func FindUserByPhone(phone string) (models.User, error) {
	var user models.User
	Database.Where("phone=?", phone).First(&user)
	if len(user.Id.String()) > 0 {
		return user, nil
	}
	return user, errors.New("no user found")
}

func DeleteUser(id uuid.UUID) error {
	var user models.User
	Database.First(&user, "id = ?", id)
	return Database.Delete(&user).Error
}

func FindUserRoleForUser(ctx *fiber.Ctx, user models.User) (models.UserRole, error) {
	var userRole models.UserRole
	Database.
		Order("access_type ASC, created_at DESC").
		First(&userRole, "user_id = ?", user.Id.String())
	if userRole.Id.String() == uuid.Nil.String() {
		return userRole, helpers.BadRequestError(ctx, "Invalid User UUID!")
	}
	return userRole, nil
}

func FindUserRoleForCustomerUser(ctx *fiber.Ctx, user models.User, customer models.Customer) (models.UserRole, error) {
	var userRole models.UserRole
	err := Database.First(&userRole, "customer_id = ? AND user_id = ?", customer.Id.String(), user.Id.String()).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return userRole, err
	}

	if userRole.Id.String() == uuid.Nil.String() {
		err := Database.First(&userRole, "access_type = ? AND user_id = ?", models.UserRoleEnum.SuperAdministrator.Number, user.Id.String()).Error
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return userRole, err
		}
	}

	if userRole.Id.String() == uuid.Nil.String() {
		return userRole, helpers.BadRequestError(ctx, "Invalid Customer UUID!")
	}

	return userRole, nil
}

func FindCurrentUserCustomer(ctx *fiber.Ctx, user models.User) (models.Customer, bool, error) {
	var currentCustomer models.Customer

	customerIdQuery := ctx.Query("customer", "")
	if customerIdQuery == "" {
		return currentCustomer, false, errors.New("no customer id provided")
	}

	customerId, err := uuid.Parse(customerIdQuery)
	if err != nil {
		log.Error().AnErr("FindCurrentUserCustomer: UUID parsing", err).Send()
		return currentCustomer, true, helpers.BadRequestError(ctx, "Invalid Customer UUID!")
	}

	currentCustomer = FindCustomerById(customerId)
	if currentCustomer.IsIdNull() {
		return currentCustomer, true, helpers.ResourceNotFoundError(ctx, "Customer")
	}

	return currentCustomer, true, nil
}
