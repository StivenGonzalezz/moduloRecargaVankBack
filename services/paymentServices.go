package services

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/payment"
	"modulo_recarga/usesCases"
)

func PostPayment(body map[string]interface{}) (*payment.Response, error) {

	// cargar las variables de entorno
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando archivo .env")
	}

	// configuración del cliente de MercadoPago con el accetkn
	cfg, err := config.New(os.Getenv("ACCESS_TOKEN"))
	if err != nil {
		return nil, fmt.Errorf("error al configurar el cliente: %v", err)
	}

	client := payment.NewClient(cfg)

	transactionAmount, err := strconv.ParseFloat(fmt.Sprintf("%v", body["transaction_amount"]), 64)
	if err != nil {
		return nil, fmt.Errorf("error en el monto de la transacción: %v", err)
	}

	paymentRequest := payment.Request{
		CallbackURL:       "https://thisisvank.com/Auth/SingIn",
		NotificationURL:   "https://thisisvank.com/Auth/SingIn",
		TransactionAmount: transactionAmount,
		Description:       fmt.Sprintf("%v", body["description"]),
		PaymentMethodID:   "pse",
		Payer: &payment.PayerRequest{
			EntityType: "individual",
			Email:      fmt.Sprintf("%v", body["email"]),
			Identification: &payment.IdentificationRequest{
				Type:   fmt.Sprintf("%v", body["identificationType"]),
				Number: fmt.Sprintf("%v", body["identificationNumber"]),
			},
			Address: &payment.AddressRequest{
				ZipCode:      fmt.Sprintf("%v", body["zipCode"]),
				StreetName:   fmt.Sprintf("%v", body["streetName"]),
				StreetNumber: fmt.Sprintf("%v", body["streetNumber"]),
				Neighborhood: fmt.Sprintf("%v", body["neighborhood"]),
				City:         fmt.Sprintf("%v", body["city"]),
				FederalUnit:  fmt.Sprintf("%v", body["federalUnit"]),
			},
			Phone: &payment.PhoneRequest{
				AreaCode: fmt.Sprintf("%v", body["phoneAreaCode"]),
				Number:   fmt.Sprintf("%v", body["phoneNumber"]),
			},
		},
		AdditionalInfo: &payment.AdditionalInfoRequest{
			IPAddress: "127.0.0.1",
		},
		TransactionDetails: &payment.TransactionDetailsRequest{
			FinancialInstitution: fmt.Sprintf("%v", body["financialInstitution"]),
		},
		ExternalReference: "MP123456789",
		BinaryMode:        true,
	}

	response, err := client.Create(context.Background(), paymentRequest)
	if err != nil {
		return nil, fmt.Errorf("error al crear el pago: %v", err)
	}

	if err := usescases.SavePaymentResponse(response, body); err != nil {
		return nil, fmt.Errorf("error al guardar la respuesta en la base de datos: %v", err)
	}

	return response, nil
}

