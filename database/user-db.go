package database

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/predictive-edge-india/ehm-go/helpers"
	"github.com/predictive-edge-india/ehm-go/models"
	"github.com/rs/zerolog/log"
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
	Database.Where("email=?", email).First(&user)

	if user.IsIdNull() {
		return user, errors.New("no user found")
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

func FindUserRoleForCustomerUser(ctx *fiber.Ctx, user models.User, customer models.Customer) (models.UserRole, error) {
	var userRole models.UserRole
	Database.First(&userRole, "customer_id = ? AND user_id = ?", customer.Id.String(), user.Id.String())

	if userRole.Id.String() == uuid.Nil.String() {
		Database.First(&userRole, "access_type = ? AND user_id = ?", models.UserRoleEnum.SuperAdministrator.Number, user.Id.String())
	}

	if userRole.Id.String() == uuid.Nil.String() {
		return userRole, helpers.BadRequestError(ctx, "Invalid Customer UUID!")
	}

	return userRole, nil
}

func FindCurrentUserCustomer(ctx *fiber.Ctx, user models.User) (models.Customer, error) {
	var currentCustomer models.Customer
	customerId, err := uuid.Parse(ctx.Params("customerId"))
	if err != nil {
		log.Error().AnErr("FindCurrentUserCustomer: UUID parsing", err).Send()
		return currentCustomer, helpers.BadRequestError(ctx, "Invalid Customer UUID!")
	}

	currentCustomer = FindCustomerById(customerId)
	if currentCustomer.IsIdNull() {
		return currentCustomer, helpers.ResourceNotFoundError(ctx, "Customer")
	}

	return currentCustomer, nil
}
