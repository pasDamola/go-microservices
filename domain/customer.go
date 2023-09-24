package domain

import "github.com/pasDamola/go-microservices/errs"

type Customer struct {
	Id string
	Name string
	City string
	Zipcode string
	DateofBirth string
	Status string
}

type CustomerRepository interface {
	// status == 0, status == 1, status = "" 
	FindAll(status string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}

