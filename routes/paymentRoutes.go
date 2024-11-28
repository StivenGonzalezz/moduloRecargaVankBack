package routes

import (
	"modulo_recarga/controllers" // Reemplaza con el nombre del paquete de tu controlador

	"github.com/gorilla/mux"
)

func PaymentRoutes(router *mux.Router) {
	// Crear una instancia de PaymentController
	paymentController := controllers.PaymentController{}

	// Crear la ruta y enlazarla con el método PostPayment
	router.HandleFunc("/api/payment/sendPayment", paymentController.PostPayment).Methods("POST")
}
