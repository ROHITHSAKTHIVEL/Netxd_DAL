package interfaces

import (
	"Netxd_DAL/models"
)

type ICustomer interface {
	CreateCustomer(customer *models.Customer) error
}
