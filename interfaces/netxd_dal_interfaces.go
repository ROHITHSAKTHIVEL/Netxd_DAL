package interfaces

import "github.com/ROHITHSAKTHIVEL/Netxd_DAL/models"

type ICustomer interface {
	CreateCustomer(customer *models.Customer) (*models.DBResponse, error)
}
