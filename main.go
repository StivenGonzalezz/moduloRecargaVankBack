package main

import (
	"fmt"
	"log"
	"modulo_recarga/infrastructure/db"
	"modulo_recarga/infrastructure/db/models"
	"modulo_recarga/routes"
	"modulo_recarga/services"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error cargando el archivo .env")
	}

	//crear un nuevo enrutador
	router := mux.NewRouter()

	//rutas para paymentMethods
	routes.PaymentMethodsRoutes(router)

	//rutas para payment
	routes.PaymentRoutes(router, db.DB)

	//configurar CORS
	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	}).Handler(router)

	//iniciar la conexión a la base de datos
	_, err = db.InitGorm()
	if err != nil {
		log.Fatalf("Error conectando a la base de datos: %v\n", err)
	}
	fmt.Println("Conexión exitosa a la base de datos")

	//migrar modelos a la base de datos
	if err := db.DB.AutoMigrate(&models.Payment{}); err != nil {
		log.Fatalf("Error al migrar la base de datos: %v", err)
	}
	fmt.Println("Migración completada exitosamente")

	// s, err := gocron.NewScheduler()
	// if err != nil {
	// 	fmt.Println("error a iniciar cron")
	// }

	var pendingPayments []models.Payment

	db.DB.Unscoped().Find(&pendingPayments, "status = ?", "pending")

	for _, payment := range pendingPayments {
		services.GetUserPayments(int(payment.ID))
	}

	//iniciar el server
	port := ":3000"
	log.Printf("Trabajando en el puerto %s", port)
	log.Fatal(http.ListenAndServe(port, handler))
}
