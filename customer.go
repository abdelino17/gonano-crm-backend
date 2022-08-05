package main

// swagger:model Customer
type Customer struct {
	Id        string `json:"customer_id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Contacted bool   `json:"contacted"`
}
