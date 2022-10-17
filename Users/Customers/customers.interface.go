package Customers

type Customers interface {
	GetCustomers() []*Customer
	GetCustomerById(customerId uint) *Customer
	AddCustomer(customerId uint, customerFirstName string, customerLastName string, customerMail string, customerPhoneNumber string, customerActivity bool)
	DeleteCustomer(customerId uint)
}
