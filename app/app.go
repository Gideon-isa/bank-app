package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	//mux := http.NewServeMux()
	//define routes
	router := mux.NewRouter()
	router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/customers", getAllCustomer).Methods("GET")
	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods("GET")
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)

	//starting server
	fmt.Println("Sarting the Server....")
	log.Fatal(http.ListenAndServe("localhost:8080", router))

}
