package domian

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRespositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "0001", Name: "Peter", Zipcode: "101010", DateofBirth: "2000-01-01", Status: "1"},
		{Id: "0002", Name: "John", Zipcode: "111001", DateofBirth: "1999-01-01", Status: "1"},
	}
	return CustomerRepositoryStub{customers}
}
