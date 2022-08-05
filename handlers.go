package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type CustomerHandlers struct {
	customers []Customer
}

// swagger:operation GET /customers GetCustomers
// Returns list of customers
// ---
// produces:
// - application/json
// responses:
//	  '200':
//	      description: Successful operation
func (ch *CustomerHandlers) GetCustomers(w http.ResponseWriter, r *http.Request) {
	writeResponse(w, http.StatusOK, ch.customers)
}

// swagger:operation POST /customers AddCustomer
// Add one customer
// ---
// produces:
// - application/json
// responses:
//     '201':
//         description: Successful operation
//     '422':
//         description: Invalid customer data
func (ch *CustomerHandlers) AddCustomer(w http.ResponseWriter, r *http.Request) {
	var customer Customer

	log.Println("Add customer operation")
	reqBody, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(reqBody, &customer); err != nil {
		writeResponse(w, http.StatusUnprocessableEntity, NewValidationError("Invalid customer data"))
		return
	}

	customer.Id = uuid.NewString()
	ch.customers = append(ch.customers, customer)
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
// responses:
//     '200':
//         description: Successful operation
//     '404':
//         description: Not found
func (ch *CustomerHandlers) GetCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	log.Println("Get customer operation")
	index := ch.findCustomer(id)
	if index == -1 {
		writeResponse(w, http.StatusNotFound, NewNotFoundError("This customer does not exist"))
		return
	}

	writeResponse(w, http.StatusOK, ch.customers[index])
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
// responses:
//     '200':
//         description: Successful operation
//     '404':
//         description: Not found
//     '422':
//         description: Invalid customer data
func (ch *CustomerHandlers) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	log.Println("Update customer operation")
	index := ch.findCustomer(id)
	if index == -1 {
		writeResponse(w, http.StatusNotFound, NewNotFoundError("This customer does not exist"))
		return
	}

	var customer Customer

	reqBody, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(reqBody, &customer); err != nil {
		writeResponse(w, http.StatusUnprocessableEntity, NewValidationError("Invalid customer data"))
		return
	}
	customer.Id = id
	ch.customers[index] = customer

	writeResponse(w, http.StatusOK, ch.customers[index])
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
// responses:
//     '200':
//         description: Successful operation
//     '404':
//         description: Not found
func (ch *CustomerHandlers) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	log.Println("Delete customer operation")
	index := ch.findCustomer(id)
	if index == -1 {
		writeResponse(w, http.StatusNotFound, NewNotFoundError("This customer does not exist"))
		return
	}

	ch.customers = append(ch.customers[:index], ch.customers[index+1:]...)

	writeResponse(w, http.StatusOK, "Customer has been deleted")
}

func (ch *CustomerHandlers) findCustomer(id string) int {
	var index int = -1

	for i := 0; i < len(ch.customers); i++ {
		if ch.customers[i].Id == id {
			index = i
		}
	}

	return index
}
