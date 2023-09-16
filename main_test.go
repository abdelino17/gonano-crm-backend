package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var ch = NewCustomerHandlers(customers)

// Tests happy path of submitting a well-formed GET /customers request
func TestGetCustomersHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/customers", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ch.GetCustomers)
	handler.ServeHTTP(rr, req)

	// Checks for 200 status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("getCustomers returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Checks for JSON response
	if ctype := rr.Header().Get("Content-Type"); ctype != "application/json" {
		t.Errorf("Content-Type does not match: got %v want %v",
			ctype, "application/json")
	}
}

// Tests happy path of submitting a well-formed POST /customers request
func TestAddCustomerHandler(t *testing.T) {
	requestBody := strings.NewReader(`
		{
			"customer_id": "test-01",
    	"name": "Alice",
    	"role": "CEO",
    	"email": "alice@udacity.com",
    	"phone": "+123400000",
    	"contacted": true
		}
	`)

	req, err := http.NewRequest("POST", "/customers", requestBody)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ch.AddCustomer)
	handler.ServeHTTP(rr, req)

	// Checks for 201 status code
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("addCustomer returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	// Checks for JSON response
	if ctype := rr.Header().Get("Content-Type"); ctype != "application/json" {
		t.Errorf("Content-Type does not match: got %v want %v",
			ctype, "application/json")
	}
}

// Tests unhappy path of deleting a user that doesn't exist
func TestDeleteCustomerHandler(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/customers/3295cbed-cdd5-4c35-a3a8-343081a0fa0d", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ch.DeleteCustomer)
	handler.ServeHTTP(rr, req)

	// Checks for 404 status code
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("deleteCustomer returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}

// Tests unhappy path of getting a user that doesn't exist
func TestGetCustomerHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/customers/3295cbed-cdd5-4c35-a3a8-343081a0fa0d", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ch.GetCustomer)
	handler.ServeHTTP(rr, req)

	// Checks for 404 status code
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("getCustomer returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}
