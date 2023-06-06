package customer

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type CustomerHandler struct {
	CustomerRepository Customerer
}

func (c *CustomerHandler) CustomerGetAll(w http.ResponseWriter, r *http.Request) {
	customers, err := c.CustomerRepository.GetAllCustomer()
	b, err := json.Marshal(customers)
	// e := json.NewEncoder(w)
	// e.Encode()
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
	fmt.Println("customer post ")
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
	fmt.Println(customerFromRequest)
}

func (c *CustomerHandler) CustomerPut(w http.ResponseWriter, r *http.Request) {
}

func (c *CustomerHandler) CustomerDelete(w http.ResponseWriter, r *http.Request) {
}
