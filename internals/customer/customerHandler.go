package customer

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type CustomerHandler struct {
	CustomerRepository Customerer
}

func (c *CustomerHandler) CustomerGetAll(w http.ResponseWriter, r *http.Request) {
	customers, err := c.CustomerRepository.GetAllCustomer()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, err := json.Marshal(customers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (c *CustomerHandler) CustomerGet(w http.ResponseWriter, r *http.Request) {
	customerId := chi.URLParam(r, "customerId")

	customerFromDb, err := c.CustomerRepository.GetCustomer(customerId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, err := json.Marshal(customerFromDb)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (c *CustomerHandler) CustomerPost(w http.ResponseWriter, r *http.Request) {
	var customerFromRequest struct {
		Name   string `json:"name"`
		Phone  string `json:"phone"`
		IsGold bool   `json:"isGold"`
	}
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	err := json.Unmarshal(body, &customerFromRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var customer Customer
	customerUUID := uuid.New()
	customer.ID = customerUUID
	customer.Name = customerFromRequest.Name
	customer.Phone = customerFromRequest.Phone
	customer.IsGold = customerFromRequest.IsGold
	err = c.CustomerRepository.CreateCustomer(customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Customer Created"))
}

func (c *CustomerHandler) CustomerPut(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var customer Customer
	err := json.Unmarshal(body, &customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	customerId := (customer.ID).String()
	customerFromDb, err := c.CustomerRepository.GetCustomer(customerId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	customerFromDb.Name = customer.Name
	customerFromDb.Phone = customer.Phone
	customerFromDb.IsGold = customer.IsGold
	err = c.CustomerRepository.UpdateCustomer(customerFromDb)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("customer updated"))
}

func (c *CustomerHandler) CustomerDelete(w http.ResponseWriter, r *http.Request) {
	customerId := chi.URLParam(r, "customerId")
	err := c.CustomerRepository.DeleteCustomer(customerId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("customer successfully deleted"))
}
