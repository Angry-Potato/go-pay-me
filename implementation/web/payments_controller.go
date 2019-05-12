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
