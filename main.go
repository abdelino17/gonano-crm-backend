// Customers API
//
// This is the customers API for Udacity Golang Nanodegree Project 1. You can find out more about the API at https://github.com/abdelino17/gonano-crm-backend
//
// Schemes: http
// Host: localhost:8080
// BasePath: /
// Version: 1.0.0
// Contact: Abdel FARE <bonjour@abdelfare.me> https://abdelfare.me
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
// swagger:meta
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var address string = "localhost"
var port string = "8080"

var customers []Customer

var (
	// Version is the git reference injected at build
	Version string
)

func init() {
	customers = make([]Customer, 0)
	file, _ := os.ReadFile("customers.json")
	_ = json.Unmarshal([]byte(file), &customers)

	if os.Getenv("SERVER_PORT") != "" {
		port = os.Getenv("SERVER_PORT")
	}

	if os.Getenv("SERVER_ADDRESS") != "" {
		address = os.Getenv("SERVER_ADDRESS")
	}
}

func main() {

	router := mux.NewRouter()
	ch := NewCustomerHandlers(customers)

	router.HandleFunc("/", showHomePage).Methods(http.MethodGet).Name("HomePage")
	router.HandleFunc("/customers", ch.GetCustomers).Methods(http.MethodGet).Name("GetCustomers")
	router.HandleFunc("/customers", ch.AddCustomer).Methods(http.MethodPost).Name("AddCustomer")
	router.HandleFunc("/customers/{id}", ch.GetCustomer).Methods(http.MethodGet).Name("GetCustomer")
	router.HandleFunc("/customers/{id}", ch.UpdateCustomer).Methods(http.MethodPut).Name("UpdateCustomer")
	router.HandleFunc("/customers/{id}", ch.DeleteCustomer).Methods(http.MethodDelete).Name("DeleteCustomer")

	log.Printf("Server is starting on port %s...", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}

func showHomePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}
