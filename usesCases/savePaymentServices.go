package usescases

import (
	"fmt"
	"log"
	"modulo_recarga/infrastructure/db"
	"modulo_recarga/infrastructure/db/models"

	"github.com/mercadopago/sdk-go/pkg/payment"
)

func SavePaymentResponse(response *payment.Response, body map[string]interface{}) error {

	db := db.DB

	if db == nil {
		fmt.Println("error al cargar db en save payment")
		return nil
	}
	payment := models.Payment{
		ExternalReference:         response.ExternalReference,
		TransactionAmount:         response.TransactionAmount,
		Description:               fmt.Sprintf("%v", body["description"]),
		PaymentMethodID:           "pse",
		PayerEmail:                fmt.Sprintf("%v", body["email"]),
		PayerIdentificationType:   fmt.Sprintf("%v", body["identificationType"]),
		PayerIdentificationNumber: fmt.Sprintf("%v", body["identificationNumber"]),
		ZipCode:                   fmt.Sprintf("%v", body["zipCode"]),
		StreetName:                fmt.Sprintf("%v", body["streetName"]),
		StreetNumber:              fmt.Sprintf("%v", body["streetNumber"]),
		Neighborhood:              fmt.Sprintf("%v", body["neighborhood"]),
		City:                      fmt.Sprintf("%v", body["city"]),
		FederalUnit:               fmt.Sprintf("%v", body["federalUnit"]),
		PhoneAreaCode:             fmt.Sprintf("%v", body["phoneAreaCode"]),
		PhoneNumber:               fmt.Sprintf("%v", body["phoneNumber"]),
		FinancialInstitution:      fmt.Sprintf("%v", body["financialInstitution"]),
		Status:                    response.Status,
		ID:                        uint(response.ID),
	}

	// inserto el pago en la db
	if err := db.Create(&payment).Error; err != nil {
		log.Printf("Error al guardar la respuesta del pago: %v", err)
		return err
	}

	return nil
}
