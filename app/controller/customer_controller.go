package controller

import "net/http"

type CustomerController interface {
	GetCustomerByID(w http.ResponseWriter, r *http.Request)
}
