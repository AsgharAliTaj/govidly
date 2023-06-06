package customer

import (
	"github.com/jmoiron/sqlx"
)

type Customerer interface {
	GetAllCustomer() ([]Customer, error)
	GetCustomer(string) (Customer, error)
	CreateCustomer(Customer) error
	UpdateCustomer(Customer) (Customer, error)
	DeleteCustomer(string) error
}

type CustomerRepository struct {
	database *sqlx.DB
}

func (c *CustomerRepository) GetAllCustomer() (customers []Customer, err error) {
	err = c.database.Select(&customers, "SELECT * FROM customers ORDER BY name ASC")
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (c *CustomerRepository) GetCustomer(id string) (customer Customer, err error) {
	err = c.database.Get(&customer, "SELECT * from customers where id = $1", id)
	if err != nil {
		return customer, err
	}
	return customer, nil
}

func (c *CustomerRepository) CreateCustomer(customer Customer) (err error) {
	err = c.database.Get(
		&customer,
		"INSERT INTO customers (id, name, phone, is_gold) VALUES ($1, $2, $3, $4) RETURNING *",
		&customer.ID,
		&customer.Name,
		&customer.Phone,
		&customer.IsGold,
	)
	if err != nil {
		return err
	}
	return nil
}

func (c *CustomerRepository) UpdateCustomer(cu Customer) (customer Customer, err error) {
	return customer, nil
}

func (c *CustomerRepository) DeleteCustomer(id string) (err error) {
	return nil
}
