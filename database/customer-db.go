package database

import (
	"github.com/google/uuid"
	"github.com/predictive-edge-india/ehm-go/models"
)

func FindCustomerById(customerId uuid.UUID) models.Customer {
	var customer models.Customer
	Database.Where("id = ?", customerId).Find(&customer)
	return customer
}
