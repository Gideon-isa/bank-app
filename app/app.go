package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Gideon-isa/banking/domian"
	"github.com/Gideon-isa/banking/service"
	"github.com/gorilla/mux"
)

func Start() {

	//
	router := mux.NewRouter()

	//wiring
	ch := CustomerHandelers{service: service.NewCustomerService(domian.NewCustomerRespositoryStub())}

	//define routes
	router.HandleFunc("/customers", ch.getAllCustomer).Methods("GET")

	//starting server
	fmt.Println("Sarting the Server....")
	log.Fatal(http.ListenAndServe("localhost:8080", router))

}
