package rest

import (
	"errors"
)

type DBLayer interface {
	GetAllProducts() ([]Product, error)
	GetPromos() ([]Product, error)
	GetCustomerByName(string, string) (Customer, error)
	GetCustomerByID(int) (Customer, error)
	GetProduct(int) (Product, error)
	AddUser(Customer) (Customer, error)
	SignInUser(username, password string) (Customer, error)
	SignOutUserById(int) error
	GetCustomerOrdersByID(int) ([]Order, error)
	AddOrder(Order) error
	GetCreditCardCID(int) (string, error)
	SaveCreditCardForCustomer(int, string) error
}

var ErrINVALIDPASSWORD = errors.New("Invalid password")
