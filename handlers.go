package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type CustomerHandlers struct {
	customers map[string]Customer
}

func NewCustomerHandlers(myCustomers []Customer) CustomerHandlers {
	customers := make(map[string]Customer)

	for _, customer := range myCustomers {
		customers[customer.Id] = customer
	}

	return CustomerHandlers{
		customers,
	}
}

// swagger:operation GET /customers GetCustomers
// Returns list of customers
// ---
// produces:
// - application/json
// responses:
//
//	'200':
//	    description: Successful operation
func (ch *CustomerHandlers) GetCustomers(w http.ResponseWriter, r *http.Request) {
	writeResponse(w, http.StatusOK, ch.customers)
}

// swagger:operation POST /customers AddCustomer
// Add one customer
// ---
// produces:
// - application/json
// responses:
//
//	'201':
//	    description: Successful operation
//	'422':
//	    description: Invalid customer data
func (ch *CustomerHandlers) AddCustomer(w http.ResponseWriter, r *http.Request) {
	var customer Customer

	log.Println("Add customer operation")
	reqBody, _ := io.ReadAll(r.Body)
	if err := json.Unmarshal(reqBody, &customer); err != nil {
		writeResponse(w, http.StatusUnprocessableEntity, NewValidationError("Invalid customer data"))
		return
	}

	customer.Id = uuid.NewString()
	ch.customers[customer.Id] = customer
	writeResponse(w, http.StatusCreated, customer)
}

// swagger:operation GET /customers/{id} GetCustomer
// Get one customer
// ---
// produces:
// - application/json
// parameters:
//   - name: id
//     in: path
//     description: customer ID
//     required: true
//     type: string
//
// responses:
//
//	'200':
//	    description: Successful operation
//	'404':
//	    description: Not found
func (ch *CustomerHandlers) GetCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	log.Println("Get customer operation")
	customer, ok := ch.customers[id]
	if !ok {
		writeResponse(w, http.StatusNotFound, NewNotFoundError("This customer does not exist"))
		return
	}

	writeResponse(w, http.StatusOK, customer)
}

// swagger:operation PUT /customers/{id} UpdateCustomer
// Update one customer
// ---
// produces:
// - application/json
// parameters:
//   - name: id
//     in: path
//     description: customer ID
//     required: true
//     type: string
//
// responses:
//
//	'200':
//	    description: Successful operation
//	'404':
//	    description: Not found
//	'422':
//	    description: Invalid customer data
func (ch *CustomerHandlers) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	log.Println("Update customer operation")
	_, ok := ch.customers[id]
	if !ok {
		writeResponse(w, http.StatusNotFound, NewNotFoundError("This customer does not exist"))
		return
	}

	var updatedCustomer Customer

	reqBody, _ := io.ReadAll(r.Body)
	if err := json.Unmarshal(reqBody, &updatedCustomer); err != nil {
		writeResponse(w, http.StatusUnprocessableEntity, NewValidationError("Invalid customer data"))
		return
	}
	updatedCustomer.Id = id

	ch.customers[id] = updatedCustomer
	writeResponse(w, http.StatusOK, updatedCustomer)
}

// swagger:operation DELETE /customers/{id} DeleteCustomer
// Delete one customer
// ---
// produces:
// - application/json
// parameters:
//   - name: id
//     in: path
//     description: customer ID
//     required: true
//     type: string
//
// responses:
//
//	'200':
//	    description: Successful operation
//	'404':
//	    description: Not found
func (ch *CustomerHandlers) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	log.Println("Delete customer operation")
	_, ok := ch.customers[id]
	if !ok {
		writeResponse(w, http.StatusNotFound, NewNotFoundError("This customer does not exist"))
		return
	}

	delete(ch.customers, id)
	writeResponse(w, http.StatusOK, "Customer has been deleted")
}
