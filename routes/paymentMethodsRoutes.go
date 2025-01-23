package routes

import (
	"modulo_recarga/controllers"

	"github.com/gorilla/mux"
)

func PaymentMethodsRoutes(router *mux.Router) {
	// Crear una instancia de PaymentMethodsController
	paymentMethodsController := controllers.PaymentMethodsController{}

	// Crear la ruta y enlazarla con el método GetPaymentMethods
	router.HandleFunc("/api/payment/paymentMethods", paymentMethodsController.GetPaymentMethods).Methods("POST")
}
