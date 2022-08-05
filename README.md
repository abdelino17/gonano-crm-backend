# Customers API

![go](https://img.shields.io/badge/go-1.18-informational)
![Twitter Follow](https://img.shields.io/twitter/follow/abdelFare?logoColor=lime&style=social)

the customers API for Udacity Golang Nanodegree Project 1.

## Setup

1. Install the last version of [golang](https://go.dev/doc/install) on your computer
2. Clone this Repo `git clone https://github.com/abdelino17/gonano-crm-backend`
3. Navigate to the new folder `cd gonano-crm-backend`
4. Run the apo with the command `go run .`
5. If you want to change the port or address of your API, you can define the two environment variables below:
   - `SERVER_PORT` (**8080** by default)
   - `SERVER_ADDRESS` (**localhost** by default)

## Usage

06 endpoints are available:

1. `GET /`: API Documentation
2. `GET /customers`: Get the list of all customers
3. `GET /customers/{id}`: Get the customer with the id corresponding
4. `POST /customers`: Add a customer
5. `PUT /customers/{id}`: Update a customer
6. `DELETE /customers/{id}`: Delete a customer
