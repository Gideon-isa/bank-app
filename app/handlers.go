package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/Gideon-isa/banking/service"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zip"`
}

type CustomerHandelers struct {
	service service.CustomerService
}

func (ch *CustomerHandelers) getAllCustomer(w http.ResponseWriter, r *http.Request) {
	// customers := []Customer{
	// 	{Name: "Mark", City: "Nigeria", Zipcode: "101123"},
	// 	{Name: "Gideon", City: "Lagos", Zipcode: "111101"},
	// }

	customers, _ := ch.service.GetAllCustomer()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)

}
