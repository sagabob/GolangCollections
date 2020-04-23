package rest

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Product struct {
	gorm.Model
	Image       string  `json:"img"`
	SmallImage  string  `gorm:"column:smallimg" json:"small_img"`
	ImagAlt     string  `json:"imgalt" gorm:"column:imgalt"`
	Price       float64 `json:"price"`
	Promotion   float64 `json:"promotion"` //sql.NullFloat64
	PoructName  string  `gorm:"column:productname" json:"productname"`
	Description string
}

func (Product) TableName() string {
	return "products"
}

type Customer struct {
	gorm.Model
	Name      string  `json:"name" sql:"-"`
	FirstName string  `gorm:"column:firstname" json:"firstname"`
	LastName  string  `gorm:"column:lastname" json:"lastname"`
	Email     string  `gorm:"column:email" json:"email"`
	Pass      string  `json:"password"`
	LoggedIn  bool    `gorm:"column:loggedin" json:"loggedin"`
	Orders    []Order `json:"orders"`
}

func (Customer) TableName() string {
	return "customers"
}

type Order struct {
	gorm.Model
	Product
	Customer
	CustomerID   int       `json:"customer_id" gorm:"column:customer_id"`
	ProductID    int       `json:"product_id" gorm:"column:product_id"`
	Price        float64   `gorm:"column:price" json:"sell_price"`
	PurchaseDate time.Time `gorm:"column:purchase_date" json:"purchase_date"`
}

func (Order) TableName() string {
	return "orders"
}

type DBORM struct {
	*gorm.DB
}

func NewORM(dbname, con string) (*DBORM, error) {
	db, err := gorm.Open(dbname, con)
	return &DBORM{
		DB: db,
	}, err
}

func (db *DBORM) GetAllProducts() ([]Product, error) {
	return nil, nil
}

func (db *DBORM) GetPromos() ([]Product, error) {
	return nil, nil
}

func (db *DBORM) GetCustomerByName(string, string) (Customer, error) {
	return Customer{}, nil
}

func (db *DBORM) GetCustomerByID(int) (Customer, error) {
	return Customer{}, nil
}

func (db *DBORM) GetProduct(int) (Product, error) {
	return Product{}, nil
}

func (db *DBORM) SignInUser(email, pass string) (customer Customer, err error) {

	//Obtain a *gorm.DB object representing our customer's row
	result := db.Table("Customers").Where(&Customer{Email: email})
	err = result.First(&customer).Error
	if err != nil {
		return customer, err
	}

	if !checkPassword(customer.Pass, pass) {
		return customer, ErrINVALIDPASSWORD
	}

	customer.Pass = ""
	//update the loggedin field
	err = result.Update("loggedin", 1).Error
	if err != nil {
		return customer, err
	}
	//return the new customer row
	return customer, result.Find(&customer).Error
}

func checkPassword(existingHash, incomingPass string) bool {
	//this method will return an error if the hash does not match the provided password string
	return bcrypt.CompareHashAndPassword([]byte(existingHash), []byte(incomingPass)) == nil
}

func (db *DBORM) SignOutUserById(int) error {
	return nil
}

func (db *DBORM) GetCustomerOrdersByID(int) ([]Order, error) {
	return nil, nil
}

func (db *DBORM) AddOrder(order Order) error {

	return db.Create(&order).Error
}
func (db *DBORM) AddUser(customer Customer) (Customer, error) {
	//pass received password by reference so that we can change it to it's hashed version
	hashPassword(&customer.Pass)
	customer.LoggedIn = true
	err := db.Create(&customer).Error
	customer.Pass = ""
	return customer, err
}

func hashPassword(s *string) error {
	if s == nil {
		return errors.New("Reference provided for hashing password is nil")
	}
	//converd password string to byte slice
	sBytes := []byte(*s)
	//Obtain hashed password
	hashedBytes, err := bcrypt.GenerateFromPassword(sBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	//update password string with the hashed version
	*s = string(hashedBytes[:])
	return nil
}

func (db *DBORM) GetCreditCardCID(id int) (string, error) {

	cusomterWithCCID := struct {
		Customer
		CCID string `gorm:"column:cc_customerid"`
	}{}

	return cusomterWithCCID.CCID, db.First(&cusomterWithCCID, id).Error
}
func (db *DBORM) SaveCreditCardForCustomer(id int, ccid string) error {
	result := db.Table("customers").Where("id=?", id)
	return result.Update("cc_customerid", ccid).Error
}
