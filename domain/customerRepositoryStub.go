package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer {
		{"1001", "Oyincode", "Abuja", "900001", "18/07/1997", "1"},
		{"2001", "Sandra", "Abuja", "900501", "21/06/1997", "1"},
	}

	return CustomerRepositoryStub{customers}
}