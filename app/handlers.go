package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zip"`
}

// greet
func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func getAllCustomer(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "Mark", City: "Nigeria", Zipcode: "101123"},
		{Name: "Gideon", City: "Lagos", Zipcode: "111101"},
	}
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)

}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintln(w, vars["customer_id"])
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post request received")
}
