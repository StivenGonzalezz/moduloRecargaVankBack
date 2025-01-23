package routes

import (
	"modulo_recarga/controllers"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func PaymentRoutes(router *mux.Router, db *gorm.DB) {
	// Crear una instancia de PaymentController
	paymentController := controllers.PaymentController{}

	// Crear la ruta y enlazarla con el método PostPayment
	router.HandleFunc("/api/payment/sendPayment", paymentController.PostPayment).Methods("POST")
}
