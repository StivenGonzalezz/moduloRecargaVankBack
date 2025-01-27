package services

import (
	"context"
	"fmt"
	"log"
	usescases "modulo_recarga/usesCases"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/payment"
)

func GetUserPayments(id int, wg *sync.WaitGroup) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando archivo .env")
	}

	cfg, err := config.New(os.Getenv("ACCESS_TOKEN"))
	if err != nil {
		fmt.Printf("error al configurar el cliente: %v", err)
	}

	client := payment.NewClient(cfg)

	response, err := client.Get(context.Background(), id)

	if err != nil {
		fmt.Println("error")
	}

	if response.Status != "pending" {
		usescases.UpdateTransferStatus(response)
	}
	wg.Done()

}
