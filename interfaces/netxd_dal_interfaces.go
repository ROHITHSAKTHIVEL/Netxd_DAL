package interfaces



type ICustomer interface {
	CreateCustomer(customer *models.Customer) (*models.DBResponse, error)
}
