package Customers

import "fmt"

type Customer struct {
	id          uint
	FirstName   string
	LastName    string
	Mail        string
	PhoneNumber string
	Activity    bool
}

func (c *Customer) GetCustomers() []*Customer {
	return nil
}
func (c *Customer) GetCustomerById(customerId uint) *Customer {
	return &Customer{
		id:          0,
		FirstName:   "",
		LastName:    "",
		Mail:        "",
		PhoneNumber: "",
		Activity:    false,
	}
}
func (c *Customer) AddCustomer(customerId uint, customerFirstName string, customerLastName string, customerMail string, customerPhoneNumber string, customerActivity bool) {
	newCustomer := Customer{
		id:          customerId,
		FirstName:   customerFirstName,
		LastName:    customerLastName,
		Mail:        customerMail,
		PhoneNumber: customerPhoneNumber,
		Activity:    customerActivity,
	}
	fmt.Println(newCustomer)

}
func (c *Customer) DeleteCustomer(customerId uint) {

}
