package web

import (
	"net/http"

	"github.com/Angry-Potato/go-pay-me/implementation/payments"
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/jinzhu/gorm"
)

// AllPayments handler
func AllPayments(DB *gorm.DB) func(w rest.ResponseWriter, r *rest.Request) {
	return func(w rest.ResponseWriter, r *rest.Request) {
		allPayments, err := payments.All(DB)

		if err != nil {
			rest.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteJson(&allPayments)
	}
}

// CreatePayment handler
func CreatePayment(DB *gorm.DB) func(w rest.ResponseWriter, r *rest.Request) {
	return func(w rest.ResponseWriter, r *rest.Request) {
		payment := payments.Payment{}
		if err := r.DecodeJsonPayload(&payment); err != nil {
			rest.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		createdPayment, err := payments.Create(DB, &payment)

		if err != nil {
			if _, ok := err.(*payments.ValidationError); ok {
				rest.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			rest.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.WriteJson(&createdPayment)
	}
}

// DeleteAllPayments handler
func DeleteAllPayments(DB *gorm.DB) func(w rest.ResponseWriter, r *rest.Request) {
	return func(w rest.ResponseWriter, r *rest.Request) {
		err := payments.DeleteAll(DB)

		if err != nil {
			rest.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.WriteJson(`{}`)
	}
}

// SetAllPayments handler
func SetAllPayments(DB *gorm.DB) func(w rest.ResponseWriter, r *rest.Request) {
	return func(w rest.ResponseWriter, r *rest.Request) {
		newPayments := []payments.Payment{}
		if err := r.DecodeJsonPayload(&newPayments); err != nil {
			rest.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		createdPayments, err := payments.SetAll(DB, newPayments)

		if err != nil {
			if _, ok := err.(*payments.ValidationError); ok {
				rest.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			rest.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.WriteJson(createdPayments)
	}
}

// GetPayment handler
func GetPayment(DB *gorm.DB) func(w rest.ResponseWriter, r *rest.Request) {
	return func(w rest.ResponseWriter, r *rest.Request) {
		ID := r.PathParam("ID")
		foundPayment, err := payments.Get(DB, ID)

		if err != nil {
			if _, ok := err.(*payments.ValidationError); ok {
				rest.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			if gorm.IsRecordNotFoundError(err) {
				rest.NotFound(w, r)
				return
			}
			rest.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.WriteJson(&foundPayment)
	}
}
