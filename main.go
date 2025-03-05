package main

import (
	"fmt"
	"log"
	"modulo_recarga/infrastructure/db"
	"modulo_recarga/infrastructure/db/models"
	"modulo_recarga/routes"
	"modulo_recarga/services"
	"os"
	"sync"
	"time"

	"net/http"

	"github.com/go-co-op/gocron/v2"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	// Verifica si el archivo .env existe antes de cargarlo
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error cargando el archivo .env")
		}
	} else {
		fmt.Println("No se encontró el archivo .env, usando variables de entorno del sistema.")
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
	_, err := db.InitGorm()
	if err != nil {
		log.Fatalf("Error conectando a la base de datos: %v\n", err)
	}
	fmt.Println("Conexión exitosa a la base de datos")

	//migrar modelos a la base de datos
	if err := db.DB.AutoMigrate(&models.Payment{}); err != nil {
		log.Fatalf("Error al migrar la base de datos: %v", err)
	}
	fmt.Println("Migración completada exitosamente")

	//inicializacion del scheduler para la creacion del job
	s, err := gocron.NewScheduler()
	if err != nil {
		fmt.Println("Error al crear scheduler")
	}
	var pendingPayments []models.Payment
	wg := &sync.WaitGroup{}

	//creacion del job para la actualizacion del estado de los pagos
	j, err := s.NewJob(
		gocron.DurationJob(
			15*time.Second,
		),
		//creacion de la tarea
		gocron.NewTask(
			func() {
				db.DB.Unscoped().Find(&pendingPayments, "status = ?", "pending")

				for _, payment := range pendingPayments {
					wg.Add(1)
					go services.GetUserPayments(int(payment.ID), wg)
				}
				wg.Wait()
				pendingPayments = nil
			},
		),
	)
	if err != nil {
		fmt.Println("Error")
	}
	fmt.Println("Jobs creado con el id:", j.ID())
	s.Start()

	//iniciar el server
	port := ":3000"
	log.Printf("Trabajando en el puerto %s", port)
	log.Fatal(http.ListenAndServe(port, handler))

	select {
	case <-time.After(30 * time.Second):
	}

	err = s.Shutdown()
	if err != nil {
		fmt.Println("error")
	}
}
