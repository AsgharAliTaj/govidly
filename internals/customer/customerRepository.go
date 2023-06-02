package customer

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Customerer interface {
	GetAllCustomer() ([]Customer, error)
	GetCustomer(uuid.UUID) Customer
	CreateCustomer(Customer) error
	UpdateCustomer(Customer) (Customer, error)
	DeleteCustomer(uuid.UUID) error
}

type customerRepository struct {
	database *sqlx.DB
}

func NewCustomerRepository(database *sqlx.DB) Customerer {
	return &customerRepository{database: database}
}

func (c *customerRepository) GetAllCustomer() (customers []Customer, err error) {
	return customers, nil
}

func (c *customerRepository) GetCustomer(uuid.UUID) (customer Customer) {
	return customer
}

func (c *customerRepository) CreateCustomer(cu Customer) (err error) {
	return nil
}

func (c *customerRepository) UpdateCustomer(cu Customer) (customer Customer, err error) {
	return customer, nil
}

func (c *customerRepository) DeleteCustomer(id uuid.UUID) (err error) {
	return nil
}
