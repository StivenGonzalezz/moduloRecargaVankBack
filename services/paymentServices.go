package services

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/google/uuid"
	mercadopago "github.com/mercadopago/sdk-go"
)

func PostPayment(body map[string]interface{}) error {
	// Configuración del cliente de MercadoPago con el token de acceso
	client := mercadopago.NewClient(os.Getenv("ACCESS_TOKEN"))

	// Crear el objeto de pago con la estructura esperada
	transactionAmount, err := strconv.ParseFloat(fmt.Sprintf("%v", body["transaction_amount"]), 64)
	if err != nil {
		return fmt.Errorf("error en el monto de la transacción: %v", err)
	}

	paymentData := map[string]interface{}{
		"transaction_amount": transactionAmount,
		"description":        body["description"],
		"payment_method_id":  "pse",
		"payer": map[string]interface{}{
			"entity_type": "individual",
			"email":       body["email"],
			"identification": map[string]interface{}{
				"type":   body["identificationType"],
				"number": body["identificationNumber"],
			},
			"address": map[string]interface{}{
				"zip_code":      body["zipCode"],
				"street_name":   body["streetName"],
				"street_number": body["streetNumber"],
				"neighborhood":  body["neighborhood"],
				"city":          body["city"],
				"federal_unit":  body["federalUnit"],
			},
			"phone": map[string]interface{}{
				"area_code": body["phoneAreaCode"],
				"number":    body["phoneNumber"],
			},
		},
		"additional_info": map[string]interface{}{
			"ip_address": "127.0.0.1",
		},
		"transaction_details": map[string]interface{}{
			"financial_institution": body["financialInstitution"],
		},
		"external_reference": "MP123456789",
		"binary_mode":        true,
	}

	// Crear un identificador único para idempotencia
	idempotencyKey := uuid.New().String()

	// Realizar la solicitud de creación del pago
	response, err := client.CreatePayment(context.Background(), paymentData, map[string]interface{}{
		"idempotencyKey": idempotencyKey,
	})

	if err != nil {
		return fmt.Errorf("error al crear el pago: %v", err)
	}

	// Imprimir el resultado del pago
	fmt.Println("Resultado del pago:", response)
	return nil
}
