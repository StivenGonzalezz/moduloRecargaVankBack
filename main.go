package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Crear un nuevo enrutador
	router := mux.NewRouter()

	// Rutas para paymentMethods
	paymentMethodsRoutes(router)

	// Rutas para payment
	paymentRoutes(router)

	// Configurar CORS
	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5500"},
	}).Handler(router)

	// Iniciar el servidor
	port := ":3000"
	log.Printf("Trabajando en el puerto %s", port)
	log.Fatal(http.ListenAndServe(port, handler))
}

// paymentMethodsRoutes define las rutas para métodos de pago
func paymentMethodsRoutes(router *mux.Router) {
	router.HandleFunc("/payment-methods", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message":"Rutas de métodos de pago"}`))
	}).Methods("GET")
}

// paymentRoutes define las rutas para pagos
func paymentRoutes(router *mux.Router) {
	router.HandleFunc("/payments", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message":"Rutas de pagos"}`))
	}).Methods("GET")
}
